package main

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/striversity/gotr/misc-ep006/ex01/au"
)

func main() {
	player := au.NewMP3Player()
	player.Enqueue(&au.MP3Song{})
	player.Enqueue(&au.MP3Song{})
	player.Enqueue(&au.MP3Song{})
	player.Play()
	player.Stop()
}
