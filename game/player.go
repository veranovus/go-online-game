package game

import (
	"online-game/client"
	"online-game/server"
)

const (
	PlayerTypeUndefined = iota
	PlayerTypeServer
	PlayerTypeClient
)

type PlayerType int

type Player struct {
	Server *server.Server
	Client *client.Client
	Type   PlayerType
}

func NewPlayer(s *server.Server, c *client.Client) *Player {
	return &Player{
		Server: s,
		Client: c,
		Type:   PlayerTypeUndefined,
	}
}
