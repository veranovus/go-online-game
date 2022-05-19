package game

const (
	PlayerTypeUndefined = iota
	PlayerTypeServer
	PlayerTypeClient
)

type PlayerType int

type Player struct {
	Type       PlayerType
	OtherReady bool
	Ready      bool
}

func NewPlayer() *Player {
	return &Player{
		Type:       PlayerTypeUndefined,
		OtherReady: false,
		Ready:      false,
	}
}
