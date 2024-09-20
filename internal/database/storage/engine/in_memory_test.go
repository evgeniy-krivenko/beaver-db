package engine

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInMemory_Get(t *testing.T) {
	t.Run("success get", func(t *testing.T) {
		t.Parallel()
		const (
			key   = "test_key"
			value = "value"
		)
		db := NewInMemory()
		db.s[key] = value

		res, err := db.Get(context.Background(), key)

		require.NoError(t, err)
		require.Equal(t, value, res)
	})

	t.Run("key not found", func(t *testing.T) {
		t.Parallel()
		db := NewInMemory()

		res, err := db.Get(context.Background(), "random-key")

		require.ErrorIs(t, err, ErrNotKeyFound)
		assert.Empty(t, res)
	})
}

func TestInMemory_SetSuccess(t *testing.T) {
	t.Parallel()

	const (
		key   = "test_key"
		value = "value"
	)

	db := NewInMemory()

	err := db.Set(context.Background(), key, value)

	require.NoError(t, err)

	assert.NotNil(t, db.s)
	assert.Len(t, db.s, 1)
	assert.Equal(t, value, db.s[key])
}

func TestInMemory_DelSuccess(t *testing.T) {
	t.Parallel()

	const (
		key   = "test_key"
		value = "value"
	)

	db := NewInMemory()

	db.s[key] = value

	err := db.Del(context.Background(), key)

	require.NoError(t, err)

	assert.NotNil(t, db.s)
	assert.Len(t, db.s, 0)
	assert.Empty(t, db.s)
}
