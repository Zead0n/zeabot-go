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
			list:    make([]Track, 0),
			current: nil,
			mode:    LoopOff,
		}
		qm.queues[guildId] = queue
	}

	return queue
}

type Queue struct {
	list    []Track
	current *Track
	mode    LoopMode
}
