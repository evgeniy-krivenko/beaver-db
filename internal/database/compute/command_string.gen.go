// Code generated by "stringer -type=CommandID -linecomment -output=command_string.gen.go"; DO NOT EDIT.

package compute

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[UnknownCommandID-0]
	_ = x[GetCommandID-1]
	_ = x[SetCommandID-2]
	_ = x[DelCommandID-3]
}

const _CommandID_name = "UnknownCommandIDGETSETDEL"

var _CommandID_index = [...]uint8{0, 16, 19, 22, 25}

func (i CommandID) String() string {
	if i < 0 || i >= CommandID(len(_CommandID_index)-1) {
		return "CommandID(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _CommandID_name[_CommandID_index[i]:_CommandID_index[i+1]]
}
