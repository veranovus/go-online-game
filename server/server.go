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

func getMessageType(msg string) int {
	if len(msg) == 0 {
		return ncom.MessageTypeUnknown
	}

	msgType, err := strconv.Atoi(string(msg[0]))
	if err != nil {
		return ncom.MessageTypeUnknown
	}

	return msgType
}

func closeSession(sess *ncom.Session) error {
	err := sess.Conn.Close()
	sess.Active = false

	if err != nil {
		return err
	}

	return nil
}

func authenticateUser(
	user *ncom.User,
	msgChannel chan ncom.Message,
	msg string,
	authSent *bool,
) (bool, error) {

	if !user.Authenticated {
		if *authSent {
			err := closeSession(user.Session)
			user.Authenticated = false

			if err != nil {
				return false, err
			}
			return false, nil
		}

		message := ncom.Message{
			User: user,
			Message: &ncom.UserAuthenticationEvent{
				Password: msg,
			},
		}

		msgChannel <- message
		*authSent = true

		return false, nil
	}

	return true, nil
}

func generateName() string {

	name := fmt.Sprintf("USER-%06d", rand.Intn(100000)+1)

	return name
}

func handleConnection(conn net.Conn, msgChannel chan ncom.Message) error {

	log.Println("[SERVER] Connection successful.")

	buff := make([]byte, 4096)
	authSent := false

	session := &ncom.Session{
		Conn:   conn,
		Active: true,
	}

	user := &ncom.User{
		Name:          generateName(),
		Session:       session,
		Authenticated: false,
	}

	msgChannel <- ncom.Message{
		User:    user,
		Message: &ncom.UserJoinedEvent{},
	}

	for {
		// Check if session is still active
		if !session.Active {
			break
		}

		// Read the message
		n, err := conn.Read(buff)
		if err != nil {
			return err
		}

		// Check the message length
		if n < 2 {
			log.Printf(
				"[CLIENT] Zero bytes. Closing the connection with '%s'.\n",
				user.Name)
			return nil
		}

		// Remove the artifacts
		msg := string(buff)
		msg = msg[:n-2]

		// Authenticate the user if use is not authenticated
		if ok, err := authenticateUser(user, msgChannel, msg, &authSent); !ok {
			if err != nil {
				return err
			}
			continue
		}

		// Get and validate the message type
		msgType := getMessageType(msg)
		if msgType == ncom.MessageTypeUnknown {
			log.Printf("[%s] Unknown type of message, or message layout.\n", user.Name)
			continue
		}

		msg = msg[1:] // Remove the type from message

		// Create the message
		var message ncom.Message

		switch msgType {
		case ncom.MessageTypeUserDisconnected:
			err = closeSession(session)

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

		// Send the message
		msgChannel <- message

		// Process leftover errors
		if err != nil {
			return err
		}
	}

	return nil
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
			if err := handleConnection(conn, server.Channel); err != nil {
				log.Println("[SERVER] Error while handling the connection.", err)
				return
			}
		}()
	}
}

func (server *Server) StartGameLoop() {
	for input := range server.Channel {
		switch event := input.Message.(type) {
		case *ncom.UserJoinedEvent:
			log.Printf("[%s] Connected.", input.User.Name)

			resp := fmt.Sprintf(
				"[SERVER] '%s' connection accepted vaiting for authentication.",
				input.User.Name)

			_ = input.User.Session.WriteLine(resp)
		case *ncom.UserDisconnectedEvent:
			log.Printf("[%s] Disconnected.", input.User.Name)

			server.Guest = nil
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
			log.Printf("[%s] `%s`\n", input.User.Name, event.Message)

			resp := fmt.Sprintf("[SERVER] ECHO - `%s`", event.Message)
			_ = input.User.Session.WriteLine(resp)
		default:
			log.Printf("[SERVER] Error! Unknown event `%v`.\n", event)
		}
	}
}
