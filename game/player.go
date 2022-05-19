package game

const (
	PlayerTypeUndefined = iota
	PlayerTypeServer
	PlayerTypeClient
)

type PlayerType int

type Player struct {
	Type       PlayerType
	Password   string
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

func (p *Player) Reset() {
	p.Type = PlayerTypeUndefined
	p.OtherReady = false
	p.Ready = false
}
