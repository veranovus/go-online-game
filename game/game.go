package game

type Game struct {
	Player *Player
	Client *Client
	Server *Server
}

func NewGame() *Game {
	g := &Game{
		Player: NewPlayer(),
	}

	g.Client = NewClient(g.Player)
	g.Server = NewServer(g.Player)

	return g
}
