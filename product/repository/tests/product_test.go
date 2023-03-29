package repository_test

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/barizalhaq/fita_shopping_api/domain"
	"github.com/barizalhaq/fita_shopping_api/product/repository"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Suite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	repository domain.ProductRepositoryInterface
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	dialector := postgres.New(postgres.Config{
		DriverName: "postgres",
		DSN:        "sqlmock_db",
		Conn:       db,
	})

	s.DB, err = gorm.Open(dialector, &gorm.Config{})
	require.NoError(s.T(), err)

	s.DB.Logger.LogMode(logger.Info)

	s.repository = repository.NewPGProductRepository(s.DB)
}

func (s *Suite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}
