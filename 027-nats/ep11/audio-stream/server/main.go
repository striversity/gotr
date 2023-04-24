package main

import (
	"awesome/audio-stream/model"
	"encoding/json"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"time"

	"github.com/charmbracelet/log"
	"github.com/nats-io/nats.go"
)

var (
	username   string
	password   string
	hostname   = "localhost"
	port       = 4222
	stream     = "songs"
	plDataFile string
	rg         = rand.New(rand.NewSource(time.Now().Unix()))
)

func init() {
	flag.StringVar(&username, "u", username, "username for NATS Server")
	flag.StringVar(&password, "p", password, "password for NATS Server")
	flag.StringVar(&hostname, "host", hostname, "NATS Server hostname")
	flag.IntVar(&port, "port", port, "NATS Server port")
	flag.StringVar(&stream, "stream", stream, "name of media stream (eg: radio, music/songs, videos, talk, etc)")
	flag.StringVar(&plDataFile, "pl", plDataFile, "path to the JSON playlist data file")
	flag.Parse()
}

func main() {
	if err := run(); err != nil {
		log.Fatal("fatal error", "error", err)
	}
}

func run() error {
	Playlists, err := loadPlaylists(plDataFile)
	if err != nil {
		return err
	}

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
		return fmt.Errorf("error getting JetStream context: %v", err)
	}

	// create JetStream kv store bucket call songs
	songBucket, err := js.CreateKeyValue(&nats.KeyValueConfig{
		Bucket: stream,
	})
	if err == nats.ErrBadBucket {
		songBucket, err = js.KeyValue(stream)
	}
	if err != nil {
		return fmt.Errorf("error creating or accessing '%v' stream bucket: %v", stream, err)
	}

	for playlist, tracks := range Playlists {
		log.Info("creating playlist", "stream-bucket", stream, "genre", playlist)
		go streamPlaylist(songBucket, playlist, tracks)
	}

	// cleanly exit application if signal is caught
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	log.Info("exiting on signal")
	return nil
}

func loadPlaylists(plDataFile string) (model.Playlists, error) {
	pls := model.Playlists{}
	f, err := os.Open(plDataFile)
	if err != nil {
		return nil, fmt.Errorf("error opening playlist data file: %v", err)
	}
	defer f.Close()

	err = json.NewDecoder(f).Decode(&pls)
	if err != nil {
		return nil, fmt.Errorf("error decoding JSON: %v", err)
	}

	return pls, nil
}

// streamPlaylist continuously loops over the list of tracks for this playlist, and selects a random track to store in the 'playlist' key.
func streamPlaylist(jsKeyVal nats.KeyValue, playlist string, tracks []model.Track) {
	numTracks := len(tracks)

	for {
		// select a random track
		trackIndex := rg.Intn(numTracks)
		track := tracks[trackIndex]

		log.Info("playing track", "playlist", playlist, "artist", track.Artist, "title", track.Title)
		// store track to the 'playlist' key
		jsKeyVal.Put(playlist, track.ToJSONBytes())
		trackDuration := 5 + time.Duration(rg.Intn(5))
		time.Sleep(trackDuration * time.Second)
	}
}
