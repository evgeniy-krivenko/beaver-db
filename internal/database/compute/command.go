package compute

//go:generate stringer -type=CommandID -linecomment -output=command_string.gen.go

type CommandID int

const (
	UnknownCommandID CommandID = iota
	GetCommandID               // GET
	SetCommandID               // SET
	DelCommandID               // DEL
)

const (
	setCommandArgsNumber = 2
	getCommandArgsNumber = 1
	delCommandArgsNumber = 1
)

func getCommandID(cmd string) CommandID {
	switch cmd {
	case GetCommandID.String():
		return GetCommandID
	case SetCommandID.String():
		return SetCommandID
	case DelCommandID.String():
		return DelCommandID
	default:
		return UnknownCommandID
	}
}
