package music

import (
	"fmt"
	"os/exec"
)

type Track struct {
	Thumbnail  string `json:"thumbnail"`
	Title      string `json:"title"`
	Url        string `json:"webpage_url"`
	Channel    string `json:"channel"`
	ChannelUrl string `json:"channel_url"`
}

func (t *Track) FormatDiscordString() string {
	return fmt.Sprintf("[%s](<%s>) by [%s](<%s>)", t.Title, t.Url, t.Channel, t.ChannelUrl)
}

func (t *Track) YtdlpCommand() *exec.Cmd {
	return exec.Command("yt-dlp", "-f", "ba", "--no-playlist", "-o", "-", t.Url)
}
