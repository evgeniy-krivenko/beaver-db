package storage

import (
	"context"
	"log/slog"

	"beaver/internal/database/storage/engine"
)

type engineStore interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value string) error
	Del(ctx context.Context, key string) error
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

func (s *Storage) Get(ctx context.Context, key string) (string, error) {
	return s.engine.Get(ctx, key)
}

func (s *Storage) Set(ctx context.Context, key string, value string) error {
	return s.engine.Set(ctx, key, value)
}

func (s *Storage) Del(ctx context.Context, key string) error {
	return s.engine.Del(ctx, key)
}

func IsKeyNotFountError(err error) bool {
	return engine.IsKeyNotFountError(err)
}
