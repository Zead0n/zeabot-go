package bot

import "github.com/disgoorg/disgolink/v3/lavalink"

type LoopState uint8

const (
	LoopOff LoopState = iota
	LoopTrack
	LoopQueue
)

type Queue struct {
	Tracks    []lavalink.Track
	Mode      LoopState
	ChannelId uint
}
