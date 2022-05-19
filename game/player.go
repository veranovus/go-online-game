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
	GameLength int32
	GameTime   int32
}

func NewPlayer() *Player {
	return &Player{
		Type:       PlayerTypeUndefined,
		OtherReady: false,
		Ready:      false,
		GameLength: 3,
		GameTime:   30,
	}
}

func (p *Player) Reset() {
	p.Type = PlayerTypeUndefined
	p.OtherReady = false
	p.Ready = false
	p.GameLength = 3
	p.GameTime = 30
}
