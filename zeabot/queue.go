package zeabot

import (
	"sync"

	"github.com/disgoorg/disgolink/v3/lavalink"
	"github.com/disgoorg/snowflake/v2"
)

type LoopState string

const (
	LoopOff   LoopState = "Off"
	LoopTrack LoopState = "Track"
	LoopQueue LoopState = "Queue"
)

func (state LoopState) String() string {
	switch state {
	case LoopOff:
		return "Off"
	case LoopTrack:
		return "Track"
	case LoopQueue:
		return "Queue"
	default:
		return "Unknown"
	}
}

type QueueManager struct {
	mu     sync.RWMutex
	queues map[snowflake.ID]*Queue
}

func (qm *QueueManager) Get(guildID snowflake.ID) *Queue {
	qm.mu.RLock()
	defer qm.mu.RUnlock()

	queue, ok := qm.queues[guildID]
	if !ok {
		queue = &Queue{
			Tracks: make([]lavalink.Track, 0),
			Mode:   LoopOff,
		}

		qm.mu.Lock()
		defer qm.mu.Unlock()

		qm.queues[guildID] = queue
	}

	return queue
}

func (qm *QueueManager) Delete(guildID snowflake.ID) {
	qm.mu.Lock()
	defer qm.mu.Unlock()

	delete(qm.queues, guildID)
}

type Queue struct {
	mu        sync.RWMutex
	Tracks    []lavalink.Track
	Mode      LoopState
	ChannelID snowflake.ID
}

func (q *Queue) Add(tracks ...lavalink.Track) {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.Tracks = append(q.Tracks, tracks...)
}

func (q *Queue) Next() (*lavalink.Track, bool) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if len(q.Tracks) <= 0 {
		return &lavalink.Track{}, false
	}

	track := q.Tracks[0]
	q.Tracks = q.Tracks[1:]

	return &track, true
}

func (q *Queue) GetTracks() []lavalink.Track {
	q.mu.RLock()
	defer q.mu.RUnlock()

	tracks := make([]lavalink.Track, 0, len(q.Tracks))
	for _, track := range q.Tracks {
		tracks = append(tracks, track)
	}

	return tracks
}

func (q *Queue) Clear() {
	q.Tracks = make([]lavalink.Track, 0)
}
