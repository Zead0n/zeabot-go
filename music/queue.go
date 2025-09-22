package music

type LoopMode int8

const (
	LoopOff LoopMode = iota
	LoopTrack
	LoopQueue
)

type (
	Queue struct {
		list    []Track
		current *Track
		mode    LoopMode
	}
)
