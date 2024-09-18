package engine

import (
	"context"
	"errors"
)

var ErrNotKeyFound = errors.New("key not found")

type InMemory struct {
	s map[string]string
}

func NewInMemory() *InMemory {
	return &InMemory{
		s: make(map[string]string),
	}
}

func (e *InMemory) Get(_ context.Context, key string) (string, error) {
	value, ok := e.s[key]
	if !ok {
		return "", ErrNotKeyFound
	}
	return value, nil
}

func (e *InMemory) Set(_ context.Context, key string, value string) error {
	e.s[key] = value
	return nil
}

func (e *InMemory) Del(_ context.Context, key string) error {
	delete(e.s, key)
	return nil
}

func IsKeyNotFountError(err error) bool {
	return errors.Is(err, ErrNotKeyFound)
}
