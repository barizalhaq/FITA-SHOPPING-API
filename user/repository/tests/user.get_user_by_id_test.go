package repository_test

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func (s *Suite) Test_repository_Get_User_By_Id() {
	var userID int = 1

	s.T().Run("Succes getting user by ID", func(t *testing.T) {
		s.mock.ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "users" WHERE "users"."id" = $1 ORDER BY "users"."id" LIMIT 1`),
		).WithArgs(userID).
			WillReturnRows(sqlmock.NewRows([]string{"id", "username"}).AddRow(userID, "some username"))

		res, err := s.repository.GetUserByID(uint64(userID))

		assert.Nil(s.T(), err)
		assert.NotNil(s.T(), res)
	})

	s.T().Run("Requested user ID not found", func(t *testing.T) {
		s.mock.ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "users" WHERE "users"."id" = $1 ORDER BY "users"."id" LIMIT 1`),
		).WithArgs(999).
			WillReturnError(gorm.ErrRecordNotFound)

		res, err := s.repository.GetUserByID(999)

		assert.Nil(s.T(), err)
		assert.Nil(s.T(), res)
	})
}
