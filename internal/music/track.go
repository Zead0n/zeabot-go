package music

import (
	"context"
	"fmt"
	"os/exec"
	"strconv"
)

type Track struct {
	media     string
	url       string
	title     string
	thumbnail string
}

func (t *Track) FfmpegCmd(ctx context.Context) *exec.Cmd {
	return exec.CommandContext(
		ctx,
		"ffmpeg",
		"-i",
		t.media,
		"-f",
		"s16le",
		"-ar",
		strconv.Itoa(48_000),
		"-ac",
		strconv.Itoa(2),
		"pipe:1",
	)
}

func (t *Track) FormatString() string {
	return fmt.Sprintf("[%s](<%s>)", t.title, t.url)
}
