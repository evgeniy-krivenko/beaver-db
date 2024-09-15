package storage

import (
	"context"
	"log/slog"

	"beaver/internal/database/storage/engine"
)

type engineStore interface {
	Get(key string) (string, error)
	Set(key string, value string) error
	Del(key string) error
}

//go:generate options-gen -out-filename=storage_options.gen.go -from-struct=Options
type Options struct {
	engine engineStore  `option:"mandatory" validate:"required"`
	logger *slog.Logger `option:"mandatory" validate:"required"`
}

type Storage struct {
	Options
}

func New(opts Options) (*Storage, error) {
	return &Storage{opts}, opts.Validate()
}

func (s *Storage) Get(_ context.Context, key string) (string, error) {
	return s.engine.Get(key)
}

func (s *Storage) Set(_ context.Context, key string, value string) error {
	return s.engine.Set(key, value)
}

func (s *Storage) Del(_ context.Context, key string) error {
	return s.engine.Del(key)
}

func IsKeyNotFountError(err error) bool {
	return engine.IsKeyNotFountError(err)
}
