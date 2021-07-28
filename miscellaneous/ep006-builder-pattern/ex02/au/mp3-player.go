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

func (player *mp3Player) Play() audioPlayer {
	logrus.Info("mp3 play()")
	return player
}
func (player *mp3Player) Stop() audioPlayer {
	logrus.Info("mp3 stop()")
	return player
}
func (player *mp3Player) Pause() audioPlayer {
	logrus.Info("mp3 pause()")
	return player
}
func (player *mp3Player) Enqueue(song Song) audioPlayer {
	logrus.Info("mp3 add to playlist()")
	return player
}
func (player *mp3Player) Dequeue(song Song) audioPlayer {
	logrus.Info("mp3 remove from playlist()")
	return player
}
func (player *mp3Player) Queue() []Song {
	logrus.Info("mp3 get playlist()")
	return player.songs
}
