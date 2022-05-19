package game

import (
	"github.com/shrainu/gnet"
	"log"
	"strconv"
)

type Server struct {
	Player *Player
	Server *gnet.Server
}

func (s *Server) Shuffle(key uint64) uint64 {
	return key
}

func (s *Server) OnUserConnect(sess *gnet.Session) bool {
	log.Println("[CLIENT] Connection successful.")
	return true
}

func (s *Server) OnUserDisconnect(sess *gnet.Session) {
	log.Println("[CLIENT] Disconnected.")
}

func (s *Server) OnUserMessages(msg gnet.Message) {
	switch msg.Type {
	case MessageTypeMessage:
		log.Printf("[CLIENT] `%s`\n", msg.Content)
		break
	case MessageTypeUserDisconnect:
		log.Println("[CLIENT] Disconnected.")
		s.Server.CloseSession(msg.Sess)
		break
	case MessageTypeSetReady:
		b, err := strconv.ParseBool(msg.Content)
		if err != nil {
			log.Println(err)
		} else {
			s.Player.OtherReady = b
		}
	default:
		log.Println("[SERVER] Unhandled, or unknown message type.")
		break
	}
}

func NewServer(p *Player) *Server {
	s := &Server{
		Player: p,
	}

	s.Server = gnet.NewServer(s)

	return s
}
