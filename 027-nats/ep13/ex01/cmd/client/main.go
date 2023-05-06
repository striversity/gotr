package main

import (
	"awesome/model"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/signal"

	"github.com/charmbracelet/log"
	"github.com/nats-io/nats.go"
)

var (
	username string
	password string
	hostname = "localhost"
	port     = 4222
	stream   string
	genre    string
)

func init() {
	flag.StringVar(&username, "u", username, "username for NATS Server")
	flag.StringVar(&password, "p", password, "password for NATS Server")
	flag.StringVar(&hostname, "host", hostname, "NATS Server hostname")
	flag.IntVar(&port, "port", port, "NATS Server port")
	flag.StringVar(&stream, "stream", stream, "name of media stream (eg: radio, music/songs, videos, talk, etc)")
	flag.StringVar(&genre, "genre", genre, "specific genre within the stream (eg: pop, rock, jazz, etc)")
	flag.Parse()
}

func main() {
	if err := run(); err != nil {
		log.Fatal("fatal error", "error", err)
	}
}

func run() error {
	url := fmt.Sprintf("nats://%v:%v", hostname, port)
	if username != "" {
		url = fmt.Sprintf("nats://%v:%v@%v:%v", username, password, hostname, port)
	}

	nc, err := nats.Connect(url)
	if err != nil {
		return fmt.Errorf("error connecting to NATS: %v", err)
	}
	defer nc.Close()

	js, err := nc.JetStream()
	if err != nil {
		return fmt.Errorf("error connecting to JetStream: %v", err)
	}

	// list all streams available if none was specified
	if stream == "" {
		streams := getAllStreams(js)

		for idx, s := range streams {
			fmt.Printf("%v - %v\n", idx+1, s)
		}

		return nil // exit application without error
	}

	// list all genres in the specified stream
	if genre == "" {
		genres := getAllGenres(js, stream)

		for idx, s := range genres {
			fmt.Printf("%v - %v\n", idx+1, s)
		}

		return nil // exit application without error
	}

	bucket, err := js.KeyValue(stream)
	if err != nil {
		return fmt.Errorf("error connecting to stream: %v", err)
	}

	genreKey, err := bucket.Watch(genre)
	if err != nil {
		return fmt.Errorf("error connecting to genre: %v", err)
	}

	go playSongs(genre, genreKey.Updates())

	// cleanly exit application if signal is caught
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	log.Info("exiting on signal")
	return nil
}

func getAllStreams(js nats.JetStreamContext) []string {
	var streams []string

	for storeName := range js.KeyValueStoreNames() {
		// strip out the prefix 'KV_'
		storeName = storeName[3:]
		streams = append(streams, storeName)
	}

	return streams
}

func getAllGenres(js nats.JetStreamContext, stream string) []string {
	bucket, err := js.KeyValue(stream)
	if err != nil {
		return nil
	}

	genres, _ := bucket.Keys()
	return genres
}

func playSongs(genre string, ch <-chan nats.KeyValueEntry) {
	log.Info("listening for songs", "genre", genre)

	var track model.Track

	for entry := range ch {
		if entry == nil { // on first connect, entry is nil
			continue
		}

		if err := json.Unmarshal(entry.Value(), &track); err != nil {
			continue
		}

		fmt.Printf("Playing song '%s' by %v\n", track.Title, track.Artist)
	}
}
