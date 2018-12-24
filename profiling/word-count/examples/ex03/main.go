/*
# Word Count - iterative

## TODO 1 - Count the occurrence of each _word_ in text files

## Requirements

1. Filenames are passed as arguments to the program
2. Use _bufio.Scanner_ to read words from a file.
3. Use _strings.ToLower()_ when comparing words.
*/
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
	"strings"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

var (
	cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
	memprofile = flag.String("memprofile", "", "write memory profile to `file`")
)

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	// ----
	if len(flag.Args()) == 0 {
		log.Error("No files to process")
		return
	}

	result := make(map[string]int)
	resLock := new(sync.Mutex)
	wg := new(sync.WaitGroup)

	start := time.Now()
	for _, fn := range flag.Args() {
		processFile(wg, fn, result, resLock)
	}
	wg.Wait()

	defer fmt.Printf("Processing took: %v\n", time.Since(start))
	printResult(result)
	// ----
	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		runtime.GC() // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
		f.Close()
	}
}

func processFile(wg *sync.WaitGroup, fn string, result map[string]int, resLock *sync.Mutex) {
	wg.Add(1)
	go func() {
		defer wg.Done()

		var w string
		r, err := os.Open(fn)
		if nil != err {
			log.Warn(err)
			return
		}
		defer r.Close()

		sc := bufio.NewScanner(r)
		sc.Split(bufio.ScanWords)

		for sc.Scan() {
			w = strings.ToLower(sc.Text())
			resLock.Lock()
			result[w] = result[w] + 1
			resLock.Unlock()
		}
	}()
}

func printResult(result map[string]int) {
	fmt.Printf("%-10s%s\n", "Count", "Word")
	fmt.Printf("%-10s%s\n", "-----", "----")

	for w, c := range result {
		fmt.Printf("%-10v%s\n", c, w)
	}
}
