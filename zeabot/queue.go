package zeabot

import (
	"sync"

	"github.com/disgoorg/disgolink/v3/lavalink"
	"github.com/disgoorg/snowflake/v2"
)

type LoopState int8

const (
	LoopOff LoopState = iota
	LoopTrack
	LoopQueue
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
	queues map[snowflake.ID]*Queue
	mu     sync.Mutex
}

func (qm *QueueManager) Get(guildID snowflake.ID) *Queue {
	queue, ok := qm.queues[guildID]
	if !ok {
		queue = &Queue{
			Tracks: make([]lavalink.Track, 0),
			Mode:   LoopOff,
		}
		qm.queues[guildID] = queue
	}

	return queue
}

func (qm *QueueManager) Delete(guildID snowflake.ID) {
	delete(qm.queues, guildID)
}

type Queue struct {
	Tracks    []lavalink.Track
	Mode      LoopState
	ChannelID snowflake.ID
}

func (q *Queue) Add(tracks ...lavalink.Track) {
	q.Tracks = append(q.Tracks, tracks...)
}

func (q *Queue) Next() (lavalink.Track, bool) {
	if len(q.Tracks) <= 0 {
		return lavalink.Track{}, false
	}

	track := q.Tracks[0]
	q.Tracks = q.Tracks[1:]

	return track, true
}

func (q *Queue) Clear() {
	q.Tracks = make([]lavalink.Track, 0)
}

func (q *Queue) SetLoop(stateString string) {
	var loopState LoopState
	switch stateString {
	case "Off":
		loopState = LoopOff
	case "Track":
		loopState = LoopTrack
	case "Queue":
		loopState = LoopQueue
	default:
		loopState = LoopOff
	}

	q.Mode = loopState
}
