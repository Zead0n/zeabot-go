package music

type LoopMode int8

const (
	Off LoopMode = iota
	TrackLoop
	QueueLoop
)

type (
	Queue struct {
		list    []Track
		current *Track
		mode    LoopMode
	}
)
