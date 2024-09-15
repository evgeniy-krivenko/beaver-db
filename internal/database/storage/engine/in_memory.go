package engine

import "errors"

var errNotKeyFound = errors.New("key not found")

type InMemory struct {
	s map[string]string
}

func NewInMemory() *InMemory {
	return &InMemory{
		s: make(map[string]string),
	}
}

func (e *InMemory) Get(key string) (string, error) {
	value, ok := e.s[key]
	if !ok {
		return "", errNotKeyFound
	}
	return value, nil
}

func (e *InMemory) Set(key, value string) error {
	e.s[key] = value
	return nil
}

func (e *InMemory) Del(key string) error {
	delete(e.s, key)
	return nil
}

func IsKeyNotFountError(err error) bool {
	return errors.Is(err, errNotKeyFound)
}
