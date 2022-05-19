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

	MessageTypePick
)
