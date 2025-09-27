package music

import (
	"os/exec"
	"strconv"
)

const (
	FRAME_RATE int = 48_000
	CHANNEL    int = 2
)

func FfmpegCommand() *exec.Cmd {
	return exec.Command(
		"ffmpeg",
		"-i",
		"pipe:0",
		"-f",
		"s16le",
		"-ar",
		strconv.Itoa(FRAME_RATE),
		"-ac",
		strconv.Itoa(CHANNEL),
		"pipe:1",
	)
}
