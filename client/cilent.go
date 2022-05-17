package client

import (
	"bufio"
	"log"
	"net"
	"online-game/ncom"
	"strconv"
	"strings"
)

type Client struct {
	Server        *ncom.User
	Channel       chan ncom.Message
	Authenticated bool
}

func (c *Client) processMessage(msgType int, msg string) error {

	// Create message
	var message ncom.Message

	switch msgType {
	case ncom.MessageTypeUserAccepted:
		c.Authenticated = true

		message = ncom.Message{
			User:    c.Server,
			Message: &ncom.UserAcceptedEvent{},
		}
		break
	case ncom.MessageTypeUserDenied:
		message = ncom.Message{
			User: c.Server,
			Message: &ncom.UserDeniedEvent{
				Message: msg,
			},
		}
		break
	case ncom.MessageTypeUserMessage:
		message = ncom.Message{
			User: c.Server,
			Message: &ncom.UserMessageEvent{
				Message: msg,
			},
		}
		break
	case ncom.MessageTypeServerMessage:
		message = ncom.Message{
			User: c.Server,
			Message: &ncom.ServerMessageEvent{
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
			User: c.Server,
			Message: &ncom.UserReadyEvent{
				State: state,
			},
		}
		break
	case ncom.MessageTypeUserSelected:
		message = ncom.Message{
			User:    c.Server,
			Message: &ncom.UserDisconnectedEvent{},
		}
		break
	}

	c.Channel <- message

	return nil
}

func (c *Client) authenticateClient(password string) error {
	err := c.Server.Session.WriteLine(password)
	return err
}

func (c *Client) ConnectToServer(address, password string) error {

	conn, err := net.Dial("tcp", address)
	if err != nil {
		return err
	}

	buff := make([]byte, 4096)

	session := &ncom.Session{
		Conn:   conn,
		Active: true,
	}

	server := &ncom.User{
		Name:    "Server",
		Session: session,
	}

	c.Server = server

	connBuffer := bufio.NewReader(conn)

	for {
		if !c.Authenticated {
			err = c.authenticateClient(password)
			if err != nil {
				err = c.Server.Session.CloseSession()
				return err
			}
		}

		n, err := connBuffer.Read(buff)
		if err != nil {
			return err
		}

		// Message queue
		var msgQueue []struct {
			int
			string
		}

		if n < 2 {
			log.Printf(
				"[SERVER] Zero bytes. Closing the connection with '%s'.\n",
				server.Name)
			return nil
		}

		// Convert buffer to string
		strBuff := string(buff)

		for {
			if index := strings.Index(strBuff, "\r\n"); index != -1 {
				// Get the full message
				tempMsg := make([]byte, len(strBuff[:index]))
				copy(tempMsg, strBuff[:index])

				msg := string(tempMsg)

				// Get message type
				msgType := ncom.GetMessageType(msg)
				if msgType == ncom.MessageTypeUnknown {
					log.Println("[CLIENT] Unknown type of message, or message layout.")
					break
				}

				// Remove the type
				msg = msg[1:]

				// Push the message
				msgQueue = append(
					msgQueue, struct {
						int
						string
					}{int: msgType, string: msg},
				)

				// Move the string buffer
				if index+2 < len(strBuff) {
					strBuff = strBuff[index+2:]
				} else {
					break
				}
			} else {
				break
			}
		}

		for _, msg := range msgQueue {

			// Process the message
			err = c.processMessage(msg.int, msg.string)

			// Leftover errors
			if err != nil {
				return err
			}
		}
	}
}

func (c *Client) StartGameLoop() {
	for input := range c.Channel {

		switch event := input.Message.(type) {
		case *ncom.UserAcceptedEvent:
			log.Println("[SERVER] Authentication successful.")
			break
		case *ncom.UserDeniedEvent:
			log.Println("[SERVER] Authentication failed.")
			break
		case *ncom.ServerMessageEvent:
			log.Printf("[SERVER] %s", event.Message)
			break
		default:
			log.Printf("[CLIENT] Error! Unknown event `%v`.\n", event)
			break
		}
	}
}
