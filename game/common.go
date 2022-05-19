package game

const (
	MessageTypeMessage = iota

	MessageTypeServerDisconnect
	MessageTypeUserDisconnect

	MessageTypeSetReady
	MessageTypePick
)
