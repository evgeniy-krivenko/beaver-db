package database_test

import (
	"io"
	"log/slog"
	"testing"

	"beaver/internal/database"
	"github.com/stretchr/testify/suite"
	gomock "go.uber.org/mock/gomock"
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

func (s *DatabaseSuite) Test_ParseQuery() {
}
