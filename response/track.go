package response

import (
	"fmt"

	"github.com/disgoorg/disgolink/v3/lavalink"
)

func FormatTrack(track *lavalink.Track) string {
	return fmt.Sprintf("[%s - %s](<%s>)", track.Info.Author, track.Info.Title, *track.Info.URI)
}
