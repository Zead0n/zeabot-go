package music

type LoopMode int

const (
	LoopOff LoopMode = iota
	LoopTrack
	LoopQueue
)

type Queue struct {
	Tracks  []Track
	current *Track
	mode    LoopMode
}

func (q *Queue) dequeue() (*Track, bool) {
	if len(q.Tracks) <= 0 {
		return &Track{}, false
	}

	next := q.Tracks[0]
	if len(q.Tracks) <= 1 {
		q.Tracks = q.Tracks[1:]
	} else {
		q.Tracks = []Track{}
	}

	return &next, true
}

func (q *Queue) enqueue(tracks ...Track) {
	for _, track := range tracks[:min(len(tracks), 10)] {
		q.Tracks = append(q.Tracks, track)
	}
}
