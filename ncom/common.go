package ncom

import (
	"fmt"
	"net"
	"strconv"
)

const (
	MessageTypeUnknown = iota

	MessageTypeUserJoined
	MessageTypeUserDisconnected

	MessageTypeUserAuthentication

	MessageTypeUserDenied
	MessageTypeUserAccepted

	MessageTypeServerMessage
	MessageTypeUserMessage

	MessageTypeUserReadyState
	MessageTypeUserSelected
)

func GetMessageType(msg string) int {
	if len(msg) == 0 {
		return MessageTypeUnknown
	}

	msgType, err := strconv.Atoi(string(msg[0]))
	if err != nil {
		return MessageTypeUnknown
	}

	return msgType
}

type Session struct {
	Conn   net.Conn
	Active bool
}

func (s *Session) CloseSession() error {
	err := s.Conn.Close()
	s.Active = false

	if err != nil {
		return err
	}

	return nil
}

func (s *Session) WriteLine(str string) error {
	_, err := s.Conn.Write([]byte(str + "\r\n"))
	return err
}

func (s *Session) SendMessage(t int32, m string) error {
	err := s.WriteLine(fmt.Sprintf("%d%s", t, m))
	return err
}

type User struct {
	Name          string
	Session       *Session
	Authenticated bool
}

type Event interface{}

type Message struct {
	User    *User
	Message Event
}

type UserJoinedEvent struct {
}

type UserDisconnectedEvent struct {
}

type UserAuthenticationEvent struct {
	Password string
}

type UserDeniedEvent struct {
	Message string
}

type UserAcceptedEvent struct {
}

type ServerMessageEvent struct {
	Message string
}

type UserMessageEvent struct {
	Message string
}

type UserReadyEvent struct {
	State bool
}
