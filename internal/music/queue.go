package music

type LoopMode int8

const (
	LoopOff LoopMode = iota
	LoopTrack
	LoopQueue
)

type QueueManager struct {
	queues map[string]*Queue
}

func (qm *QueueManager) GetGuildQueue(guildId string) *Queue {
	queue, ok := qm.queues[guildId]
	if !ok {
		queue = &Queue{
			tracks:  make([]YtdlpEntry, 0),
			current: nil,
			mode:    LoopOff,
		}
		qm.queues[guildId] = queue
	}

	return queue
}

type Queue struct {
	tracks  []YtdlpEntry
	current *YtdlpEntry
	mode    LoopMode
}

func (q *Queue) Next() (*YtdlpEntry, bool) {
	next := q.tracks[0]
	q.tracks = q.tracks[1:]
	return &next, true
}

func (q *Queue) Play(track *YtdlpEntry) bool {
	q.current = track
	return true
}
