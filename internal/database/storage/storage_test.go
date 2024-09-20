package storage_test

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/suite"
	gomock "go.uber.org/mock/gomock"

	"beaver/internal/database/storage"
	"beaver/internal/database/storage/engine"
)

//go:generate mockgen -source=storage.go -destination=storage_mock_test.go -package=storage_test

type StorageSuite struct {
	suite.Suite
	ctrl   *gomock.Controller
	engine *MockengineStore
	logger *slog.Logger

	storage *storage.Storage
}

func TestStorageSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(StorageSuite))
}

func (s *StorageSuite) SetupTest() {
	s.ctrl = gomock.NewController(s.T())
	s.engine = NewMockengineStore(s.ctrl)
	s.logger = slog.New(slog.NewTextHandler(io.Discard, nil))

	var err error
	s.storage, err = storage.New(storage.NewOptions(s.engine, s.logger))
	s.Require().NoError(err)
}

func (s *StorageSuite) TearDownTest() {
	s.ctrl.Finish()
}

func (s *StorageSuite) Test_Get_Success() {
	// Arrange
	const (
		key   = "some-key"
		value = "some-value"
	)

	s.engine.EXPECT().
		Get(gomock.Any(), key).
		Return(value, nil)

	// Action
	result, err := s.storage.Get(context.Background(), key)

	// Assert
	s.Require().NoError(err)

	s.Equal(value, result)
}

func (s *StorageSuite) Test_Get_NotFoundErr() {
	// Arrange
	const (
		key   = "some-key"
		value = "some-value"
	)

	s.engine.EXPECT().
		Get(gomock.Any(), key).
		Return("", engine.ErrNotKeyFound)

	// Action
	result, err := s.storage.Get(context.Background(), key)

	// Assert
	s.Require().Error(err)

	s.True(storage.IsKeyNotFountError(fmt.Errorf("wrapped: %w", err)))
	s.Empty(result)
}

func (s *StorageSuite) Test_Set_Success() {
	// Arrange
	const (
		key   = "some-key"
		value = "some-value"
	)

	s.engine.EXPECT().
		Set(gomock.Any(), key, value).
		Return(nil)

	// Action
	err := s.storage.Set(context.Background(), key, value)

	// Assert
	s.Require().NoError(err)
}

func (s *StorageSuite) Test_Set_Error() {
	// Arrange
	const (
		key   = "some-key"
		value = "some-value"
	)

	s.engine.EXPECT().
		Set(gomock.Any(), key, value).
		Return(errors.New("unexpected"))

	// Action
	err := s.storage.Set(context.Background(), key, value)

	// Assert
	s.Require().Error(err)
}

func (s *StorageSuite) Test_Del_Success() {
	// Arrange
	const (
		key = "some-key"
	)

	s.engine.EXPECT().
		Del(gomock.Any(), key).
		Return(nil)

	// Action
	err := s.storage.Del(context.Background(), key)

	// Assert
	s.Require().NoError(err)
}

func (s *StorageSuite) Test_Del_Error() {
	// Arrange
	const (
		key = "some-key"
	)

	s.engine.EXPECT().
		Del(gomock.Any(), key).
		Return(errors.New("unexpected"))

	// Action
	err := s.storage.Del(context.Background(), key)

	// Assert
	s.Require().Error(err)
}
