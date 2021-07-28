package main

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/striversity/gotr/misc-ep006/ex02/au"
)

func main() {
	player := au.NewMP3Player()
	player.Enqueue(&au.MP3Song{}).Enqueue(&au.MP3Song{}).Enqueue(&au.MP3Song{}).Play().Stop()
	for _, song := range player.Queue() {
		_ = song // use song
	}
}
