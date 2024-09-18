package database_test

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/suite"
	gomock "go.uber.org/mock/gomock"

	"beaver/internal/database"
	"beaver/internal/database/compute"
)

//go:generate mockgen -source=database.go -destination=database_mock_test.go -package=database_test

type DatabaseSuite struct {
	suite.Suite
	ctrl    *gomock.Controller
	parser  *MockcomputeParser
	storage *Mockstorage
	logger  *slog.Logger

	db *database.Database
}

func TestDatabaseSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(DatabaseSuite))
}

func (s *DatabaseSuite) SetupTest() {
	s.ctrl = gomock.NewController(s.T())
	s.parser = NewMockcomputeParser(s.ctrl)
	s.storage = NewMockstorage(s.ctrl)
	s.logger = slog.New(slog.NewTextHandler(io.Discard, nil))

	var err error
	s.db, err = database.New(database.NewOptions(s.parser, s.storage, s.logger))
	s.Require().NoError(err)
}

func (s *DatabaseSuite) TearDownTest() {
	s.ctrl.Finish()
}

func (s *DatabaseSuite) Test_ParseQuery_Error() {
	// Arrange
	q := "GET some"

	err := errors.New("unexpected parse err")

	s.parser.EXPECT().
		Parse(q).
		Return(compute.Query{}, err)

	res := s.db.ParseQuery(context.Background(), q)

	s.Require().Contains(res, err.Error())
	s.Contains(res, database.ErrPrefix)
}

func (s *DatabaseSuite) Test_GetCMDError() {
	// Arrange
	const key = "some-key"

	err := errors.New("unexpected get err")
	query := compute.NewQuery(compute.GetCommandID, []string{key})

	s.parser.EXPECT().
		Parse(gomock.Any()).
		Return(query, nil)

	s.storage.EXPECT().
		Get(gomock.Any(), key).
		Return("", err)

	res := s.db.ParseQuery(context.Background(), "")

	s.Require().Contains(res, err.Error())
	s.Contains(res, database.ErrPrefix)
}

func (s *DatabaseSuite) Test_GetCMDSuccess() {
	// Arrange
	const (
		key   = "some-key"
		value = "some value"
	)

	query := compute.NewQuery(compute.GetCommandID, []string{key})

	s.parser.EXPECT().
		Parse(gomock.Any()).
		Return(query, nil)

	s.storage.EXPECT().
		Get(gomock.Any(), key).
		Return(value, nil)

	res := s.db.ParseQuery(context.Background(), "")

	s.Require().Contains(res, value)
}

func (s *DatabaseSuite) Test_SetCMDError() {
	// Arrange
	const (
		key   = "some-key"
		value = "some value"
	)

	err := errors.New("unexpected set err")
	query := compute.NewQuery(compute.SetCommandID, []string{key, value})

	s.parser.EXPECT().
		Parse(gomock.Any()).
		Return(query, nil)

	s.storage.EXPECT().
		Set(gomock.Any(), key, value).
		Return(err)

	res := s.db.ParseQuery(context.Background(), "")

	s.Require().Contains(res, err.Error())
	s.Contains(res, database.ErrPrefix)
}

func (s *DatabaseSuite) Test_SetCMDSuccess() {
	// Arrange
	const (
		key   = "some-key"
		value = "some value"
	)

	query := compute.NewQuery(compute.SetCommandID, []string{key, value})

	s.parser.EXPECT().
		Parse(gomock.Any()).
		Return(query, nil)

	s.storage.EXPECT().
		Set(gomock.Any(), key, value).
		Return(nil)

	res := s.db.ParseQuery(context.Background(), "")

	s.Require().Contains(res, database.SuccessMsg)
}

func (s *DatabaseSuite) Test_DelCMDError() {
	// Arrange
	const (
		key = "some-key"
	)

	query := compute.NewQuery(compute.DelCommandID, []string{key})

	err := errors.New("unexpected del error")

	s.parser.EXPECT().
		Parse(gomock.Any()).
		Return(query, nil)

	s.storage.EXPECT().
		Del(gomock.Any(), key).
		Return(err)

	res := s.db.ParseQuery(context.Background(), "")

	s.Require().Contains(res, err.Error())
	s.Contains(res, database.ErrPrefix)
}

func (s *DatabaseSuite) Test_DelCMDSuccess() {
	// Arrange
	const (
		key = "some-key"
	)

	query := compute.NewQuery(compute.DelCommandID, []string{key})

	s.parser.EXPECT().
		Parse(gomock.Any()).
		Return(query, nil)

	s.storage.EXPECT().
		Del(gomock.Any(), key).
		Return(nil)

	res := s.db.ParseQuery(context.Background(), "")

	s.Require().Contains(res, database.SuccessMsg)
}

func (s *DatabaseSuite) Test_UnknownCMDError() {
	// Arrange
	const (
		key = "some-key"
	)

	query := compute.NewQuery(compute.UnknownCommandID, []string{key})

	s.parser.EXPECT().
		Parse(gomock.Any()).
		Return(query, nil)

	res := s.db.ParseQuery(context.Background(), "")

	s.Require().Contains(res, database.ErrPrefix)
	s.Require().Contains(res, database.InternalErrText)
}
