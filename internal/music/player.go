package music

type PlayerState int

const (
	PlayerStateEnded PlayerState = iota
	PlayerStatePlaying
	PlayerStatePaused
)

type Player struct {
	state PlayerState
}
