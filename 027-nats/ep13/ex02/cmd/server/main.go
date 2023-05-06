package main

import (
	"awesome/model"
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
	plDataFile string
	rg         = rand.New(rand.NewSource(time.Now().Unix()))
)

func init() {
	flag.StringVar(&username, "u", username, "username for NATS Server")
	flag.StringVar(&password, "p", password, "password for NATS Server")
	flag.StringVar(&hostname, "host", hostname, "NATS Server hostname")
	flag.IntVar(&port, "port", port, "NATS Server port")
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

	for genre, tracks := range Playlists {
		log.Info("creating playlist", "stream-bucket", genre, "genre", genre)
		go streamPlaylist(js, genre, tracks)
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

// streamPlaylist continuously loops over the list of tracks for this playlist/genre,
// and selects a random track to add to the bucket. NOTE: If the key already exists,
// select another track.
func streamPlaylist(js nats.JetStreamContext, genre string, tracks []model.Track) {
	genreBucket, err := js.CreateKeyValue(&nats.KeyValueConfig{
		Bucket: genre,
		TTL:    1 * time.Minute,
	})
	if err !=nil {
		genreBucket, err = js.KeyValue(genre)
	}
	if err != nil {
		log.Error("error creating or accessing stream bucket", "genre", genre, "error", err)
		return
	}

	// post the first track on startup, then delay 5 seconds adding new tracks every 10 seconds
	track := getNextTrack(tracks)
	genreBucket.Put(genre+"-"+track.Id, track.ToJSONBytes())
	time.Sleep(5 * time.Second)

	for {
		track := getNextTrack(tracks)
		if _, err := genreBucket.Create(genre+"-"+track.Id, track.ToJSONBytes()); err != nil {
			continue
		}

		log.Info("queuing track", "genre", genre, "artist", track.Artist, "title", track.Title)
		time.Sleep(10 * time.Second)
	}
}

// getNextTrack returns the next track to play from the playlist.
func getNextTrack(tracks []model.Track) model.Track {
	numTracks := len(tracks)
	trackIndex := rg.Intn(numTracks)
	track := tracks[trackIndex]
	return track
}
