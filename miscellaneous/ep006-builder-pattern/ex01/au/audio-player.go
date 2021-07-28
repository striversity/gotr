package au

type (
	audioPlayer interface {
		Play()
		Stop()
		Pause()
		Enqueue(song Song)
		Dequeue(song Song)
		Queue() []Song
	}

	Song interface {
		//..
	}
)
