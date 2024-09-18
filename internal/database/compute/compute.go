package compute

import (
	"errors"
	"fmt"
	"log/slog"
	"strings"
)

var (
	errWrongNumArgs   = errors.New("wrong number of arguments")
	errUnknownCommand = errors.New("unknown command")
	errInvalidQuery   = errors.New("invalid query")
)

type Compute struct {
	logger *slog.Logger
}

func New(logger *slog.Logger) *Compute {
	return &Compute{logger: logger}
}

func (c *Compute) Parse(query string) (Query, error) {
	q := slog.String("query", query)

	tokens := strings.Fields(query)
	if len(tokens) == 0 {
		c.logger.Debug("empty tokens", q)
		return Query{}, errInvalidQuery
	}

	commandID := getCommandID(tokens[0])
	if commandID == UnknownCommandID {
		c.logger.Debug("wrong command", q)
		return Query{}, errUnknownCommand
	}

	args := tokens[1:]

	if nums := numsArgs(commandID); len(args) != nums {
		c.logger.Debug("wrong number of arguments", q)
		return Query{}, fmt.Errorf(
			"%w for command %s: needs %d",
			errWrongNumArgs,
			commandID,
			nums,
		)
	}

	return NewQuery(commandID, args), nil
}

func numsArgs(cmd CommandID) int {
	switch cmd {
	case GetCommandID:
		return getCommandArgsNumber
	case SetCommandID:
		return setCommandArgsNumber
	case DelCommandID:
		return delCommandArgsNumber
	}

	return 0
}
