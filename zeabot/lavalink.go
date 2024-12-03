package zeabot

import (
	"context"
	"log/slog"

	"github.com/disgoorg/disgolink/v3/disgolink"
	"github.com/disgoorg/disgolink/v3/lavalink"
)

func (z *Zeabot) onTrackEnd(player disgolink.Player, event lavalink.TrackEndEvent) {
	if !event.Reason.MayStartNext() {
		return
	}

	queue := z.Manager.Get(event.GuildID())
	var (
		nextTrack *lavalink.Track
		ok        bool = true
	)

	switch queue.Mode {
	case LoopOff:
		nextTrack, ok = queue.Next()
	case LoopTrack:
		nextTrack = &event.Track
	case LoopQueue:
		queue.Add(event.Track)
		nextTrack, ok = queue.Next()
	}

	if !ok {
		return
	}
	if err := player.Update(context.TODO(), lavalink.WithTrack(*nextTrack)); err != nil {
		slog.Error("Failed to play next track", slog.Any("err", err))
	}
}
