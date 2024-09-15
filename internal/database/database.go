package database

import (
	"context"
	"fmt"
	"log/slog"

	"beaver/internal/database/compute"
)

type computeParser interface {
	Parse(string) (compute.Query, error)
}

type storage interface {
	Get(cxt context.Context, key string) (string, error)
	Set(cxt context.Context, key string, value string) error
	Del(cxt context.Context, key string) error
}

//go:generate options-gen -out-filename=database_options.gen.go -from-struct=Options
type Options struct {
	parser  computeParser `option:"mandatory" validate:"required"`
	storage storage       `option:"mandatory" validate:"required"`
	logger  *slog.Logger  `option:"mandatory" validate:"required"`
}

type Database struct {
	Options
}

func New(opts Options) (*Database, error) {
	return &Database{Options: opts}, opts.Validate()
}

func (d *Database) ParseQuery(ctx context.Context, query string) string {
	d.logger.Debug("handling query", slog.String("query", query))
	q, err := d.parser.Parse(query)
	if err != nil {
		return fmt.Sprintf("ERR: %s", err.Error())
	}

	switch q.CommandID() {
	case compute.GetCommandID:
		return d.handleGetCommand(ctx, q.Args())
	case compute.SetCommandID:
		return d.handleSetCommand(ctx, q.Args())
	case compute.DelCommandID:
		return d.handleDelCommand(ctx, q.Args())
	}

	d.logger.Error(
		"compute is incorrect",
		slog.Int("command_id", int(q.CommandID())),
	)

	return "ERR: internal error"
}

func (d *Database) handleGetCommand(ctx context.Context, args []string) string {
	key := args[0]
	res, err := d.storage.Get(ctx, key)
	if err != nil {
		return fmt.Sprintf("ERR: %v", err)
	}
	return res
}

func (d *Database) handleSetCommand(ctx context.Context, args []string) string {
	key, value := args[0], args[1]

	err := d.storage.Set(ctx, key, value)
	if err != nil {
		return fmt.Sprintf("ERR: %v", err)
	}
	return "OK"
}

func (d *Database) handleDelCommand(ctx context.Context, args []string) string {
	key := args[0]
	err := d.storage.Del(ctx, key)
	if err != nil {
		return fmt.Sprintf("ERR: %v", err)
	}
	return "OK"
}
