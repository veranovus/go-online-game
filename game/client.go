package game

import (
	"github.com/shrainu/gnet"
	"log"
	"strconv"
)

type Client struct {
	Player *Player
	Client *gnet.Client
}

func (c *Client) Shuffle(key uint64) uint64 {
	return key
}

func NewClient(p *Player) *Client {
	c := &Client{
		Player: p,
	}

	c.Client = gnet.NewClient(c)

	return c
}

func (c *Client) WaitForConnection() bool {
	if c.Client.Session == nil || !c.Client.Connected() {
		return false
	}
	return true
}

func (c *Client) ProcessMessages() {
	for !c.WaitForConnection() {
		continue
	}

	for c.Client.Connected() {
		for msg := range c.Client.Channel {

			switch msg.Type {
			case MessageTypeMessage:
				log.Printf("[SERVER] `%s`\n", msg.Content)
				break
			case MessageTypeServerDisconnect:
				log.Println("[SERVER] Disconnected.")
				c.Client.Session.Close()
				return
			case MessageTypeSetReady:
				b, err := strconv.ParseBool(msg.Content)
				if err != nil {
					log.Println(err)
				} else {
					c.Player.OtherReady = b
				}
				break
			default:
				log.Println("[CLIENT] Unhandled, or unknown message type.")
				break
			}
		}
	}
}
