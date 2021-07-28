package au

import "github.com/sirupsen/logrus"

type (
	mp3Player struct {
		songs []Song
	}

	MP3Song struct {
		data []byte
	}
)

func NewMP3Player() *mp3Player {
	return &mp3Player{}
}

func (player *mp3Player) Play() {
	logrus.Info("mp3 play()")
}
func (player *mp3Player) Stop() {
	logrus.Info("mp3 stop()")
}
func (player *mp3Player) Pause() {
	logrus.Info("mp3 pause()")
}
func (player *mp3Player) Enqueue(song Song) {
	logrus.Info("mp3 add to playlist()")
}
func (player *mp3Player) Dequeue(song Song) {
	logrus.Info("mp3 remove from playlist()")
}
func (player *mp3Player) Queue() []Song {
	logrus.Info("mp3 get playlist()")
	return player.songs
}
