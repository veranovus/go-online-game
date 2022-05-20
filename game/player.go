package game

const (
	PlayerTypeUndefined = iota
	PlayerTypeServer
	PlayerTypeClient
)

type PlayerType int

type Player struct {
	Name string
	Type PlayerType

	Password string

	OtherReady bool
	Ready      bool
	StartGame  bool

	GameLength int32
	GameTime   int32

	OtherCard int32
	Card      int32

	OtherScore int32
	Score      int32
}

func NewPlayer() *Player {
	return &Player{
		Type:       PlayerTypeUndefined,
		OtherReady: false,
		Ready:      false,
		StartGame:  false,
		GameLength: 3,
		GameTime:   30,
		Card:       0,
		OtherCard:  0,
		OtherScore: 0,
		Score:      0,
	}
}

func (p *Player) Reset() {
	p.Type = PlayerTypeUndefined

	p.OtherReady = false
	p.Ready = false
	p.StartGame = false

	p.GameLength = 3
	p.GameTime = 30

	p.Card = 0
	p.OtherCard = 0

	p.OtherScore = 0
	p.Score = 0
}
