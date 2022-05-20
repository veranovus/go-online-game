package game

const (
	MessageTypeMessage = iota

	MessageTypeUserPassword

	MessageTypeUserAuthFailed
	MessageTypeUserAuthSuccessful

	MessageTypeServerDisconnect
	MessageTypeUserDisconnect

	MessageTypeSetReady
	MessageTypeSetGameProperties
	MessageTypeStartGame

	MessageTypePick
)

const (
	CardTypeNone = iota
	CardTypeRock
	CardTypePaper
	CardTypeScissor
)
