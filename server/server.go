package server

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"online-game/ncom"
	"strconv"
)

type Server struct {
	Guest    *ncom.User
	Password string
	Channel  chan ncom.Message
}

func CheckValidConnection(user *ncom.User) bool {
	if !user.Authenticated {
		user.Session.Active = false
		return false
	}
	return true
}

func GenerateName() string {

	name := fmt.Sprintf("USER-%06d", rand.Intn(100000)+1)

	return name
}

func HandleConnection(conn net.Conn, messageChannel chan ncom.Message) error {

	log.Println("[SERVER] Connection successful.")

	buff := make([]byte, 4096)

	session := &ncom.Session{
		Conn:   conn,
		Active: true,
	}

	user := &ncom.User{
		Name:          GenerateName(),
		Session:       session,
		Authenticated: false,
	}

	messageChannel <- ncom.Message{
		User:    user,
		Message: &ncom.UserJoinedEvent{},
	}

	for {
		if !session.Active {
			err := session.Conn.Close()
			if err != nil {
				log.Printf(
					"[%s] Error while closing connection.\n",
					user.Name)
				log.Printf("[%s] Error: %v", user.Name, err)
			}

			return nil
		}

		n, err := conn.Read(buff)
		if err != nil {
			return err
		}

		if n < 2 { // Old code was `n == 0`
			log.Printf(
				"[CLIENT] Zero bytes. Closing the connection with '%s'.\n",
				user.Name)
			return nil
		}

		msg := string(buff)
		msg = msg[:n-2]

		messageType, err := strconv.Atoi(string(msg[0]))
		if err != nil {
			log.Println("[CLIENT] Error unknown message layout.")
		}

		msg = msg[1:]
		var message ncom.Message

		switch messageType {
		case ncom.MessageTypeUserDisconnected:
			message = ncom.Message{
				User:    user,
				Message: &ncom.UserDisconnectedEvent{},
			}
			break
		case ncom.MessageTypeUserAuthentication:
			message = ncom.Message{
				User: user,
				Message: &ncom.UserAuthenticationEvent{
					Password: msg,
				},
			}
			break
		case ncom.MessageTypeUserMessage:
			message = ncom.Message{
				User: user,
				Message: &ncom.UserMessageEvent{
					Message: msg,
				},
			}
			break
		case ncom.MessageTypeUserReadyState:
			state, err := strconv.ParseBool(msg)
			if err != nil {
				log.Printf(
					"[CLIENT] Error expected bool for UserReadyState, found `%s`.\n",
					msg)
			}

			message = ncom.Message{
				User: user,
				Message: &ncom.UserReadyEvent{
					State: state,
				},
			}
			break
		case ncom.MessageTypeUserSelected:
			message = ncom.Message{
				User:    user,
				Message: &ncom.UserDisconnectedEvent{},
			}
			break
		}

		messageChannel <- message
	}
}

func (server *Server) StartServer(address, password string) error {
	log.Println("[SERVER] Starting.")

	ln, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	server.Password = password

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("[SERVER] Failed to accept the connection.", err)
			continue
		}

		if conn != nil && server.Guest != nil {
			sess := &ncom.Session{Conn: conn}

			err = sess.WriteLine(
				fmt.Sprintf("%cAnother user is connected.", ncom.MessageTypeUserDenied))
			if err != nil {
				return err
			}

			err = sess.Conn.Close()
			if err != nil {
				log.Println("[SERVER] Error while denying the incoming connection.")
			}

			continue
		}

		go func() {
			if err := HandleConnection(conn, server.Channel); err != nil {
				log.Println("[SERVER] Error while handling the connection.", err)
				return
			}
		}()
	}
}

func (server *Server) StartGameLoop() {
	for input := range server.Channel {
		fmt.Println("CALLED!")
		switch event := input.Message.(type) {
		case *ncom.UserJoinedEvent:
			log.Printf("[%s] Connected.", input.User.Name)

			resp := fmt.Sprintf(
				"[SERVER] '%s' connection accepted vaiting for authentication.",
				input.User.Name)

			_ = input.User.Session.WriteLine(resp)
		case *ncom.UserDisconnectedEvent:
			if !CheckValidConnection(input.User) {
				continue
			}

			log.Printf("[%s] Disconnected.", input.User.Name)

			input.User.Session.Active = false
		case *ncom.UserAuthenticationEvent:
			if event.Password == server.Password {
				input.User.Authenticated = true
				server.Guest = input.User

				log.Printf("[%s] Authenticated.\n", input.User.Name)

				resp := fmt.Sprintf("[SERVER] Authentication succesful.")
				_ = input.User.Session.WriteLine(resp)
			} else {
				log.Printf("[%s] Failed to authenticate.\n", input.User.Name)
			}
		case *ncom.UserMessageEvent:
			if !CheckValidConnection(input.User) {
				continue
			}

			log.Printf("[%s] `%s`\n", input.User.Name, event.Message)

			resp := fmt.Sprintf("[SERVER] ECHO - `%s`", event.Message)
			_ = input.User.Session.WriteLine(resp)
		default:
			if !CheckValidConnection(input.User) {
				continue
			}

			log.Printf("[SERVER] Error! Unknown event `%v`.\n", event)
		}
	}
}
