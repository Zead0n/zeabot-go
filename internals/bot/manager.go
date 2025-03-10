package bot

type QueueManager struct {
	queues map[uint]*Queue
}

func (qm *QueueManager) Get(guildId string)
