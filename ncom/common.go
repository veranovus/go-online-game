package ncom

import "net"

const (
	MessageTypeUnknown = iota

	MessageTypeUserJoined
	MessageTypeUserDisconnected

	MessageTypeUserAuthentication

	MessageTypeUserDenied
	MessageTypeUserAccepted

	MessageTypeUserMessage
	MessageTypeUserReadyState
	MessageTypeUserSelected
)

type Session struct {
	Conn   net.Conn
	Active bool
}

func (s *Session) WriteLine(str string) error {
	_, err := s.Conn.Write([]byte(str + "\r\n"))
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
}

type UserAcceptedEvent struct {
}

type UserMessageEvent struct {
	Message string
}

type UserReadyEvent struct {
	State bool
}
