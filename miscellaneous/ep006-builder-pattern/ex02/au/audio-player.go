package au

type (
	audioPlayer interface {
		Play()  audioPlayer
		Stop() audioPlayer
		Pause() audioPlayer
		Enqueue(song Song) audioPlayer
		Dequeue(song Song) audioPlayer
		Queue() []Song
	}

	Song interface {
		//..
	}
)
