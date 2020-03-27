/*
# Word Count - concurrent file read

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

const defaultWorkers = 4

var (
	cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
	memprofile = flag.String("memprofile", "", "write memory profile to `file`")
	maxWorkers int
)

func main() {
	flag.IntVar(&maxWorkers, "w", defaultWorkers, "number of workers for processing input")
	flag.Parse()
	// CPU profiling
	// ----
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

	// -----------------
	if len(flag.Args()) == 0 {
		log.Error("No files to process")
		return
	}

	workersWG := new(sync.WaitGroup)
	partialResults := make(chan map[string]int, maxWorkers)
	workQueue := make(chan string, maxWorkers)
	reducerWG := new(sync.WaitGroup)
	finalResult := make(map[string]int)

	start := time.Now()
	reducer(reducerWG, finalResult, partialResults)
	for i := 0; i < maxWorkers; i++ { // start workers
		processFile(workersWG, partialResults, workQueue)
	}
	for _, fn := range flag.Args() {
		workQueue <- fn // send work
	}
	close(workQueue)      // no more work to hand out, worker goroutines cleanup
	workersWG.Wait()      // wait for all workers to finish, only applies to processFile()
	close(partialResults) // signal aggregator to exit
	reducerWG.Wait()

	defer fmt.Printf("Processing took: %v\n", time.Since(start))
	printResult(finalResult)

	// Mem profiling
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
		// -------------
	}
}

// processFile wait for work on the workQueue
func processFile(wg *sync.WaitGroup, result chan<- map[string]int, workQueue <-chan string) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		var w string
		for fn := range workQueue { // get work
			res := make(map[string]int)
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
				res[w] = res[w] + 1
			}
			result <- res // send result
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

// reducer aggregates the intermediate-result from each worker. it exits when the 'in' queue closes
func reducer(wg *sync.WaitGroup, finResult map[string]int,
	in <-chan map[string]int) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for res := range in {
			for k, v := range res {
				finResult[k] += v
			}
		}
	}()
}
