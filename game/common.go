package game

const (
	MessageTypeMessage = iota

	MessageTypeUserPassword

	MessageTypeUserAuthFailed // TODO : Add functionality.
	MessageTypeUserAuthSuccessful

	MessageTypeServerDisconnect
	MessageTypeUserDisconnect

	MessageTypeSetReady
	MessageTypePick
)
