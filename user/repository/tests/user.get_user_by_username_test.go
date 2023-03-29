package repository_test

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func (s *Suite) Test_repository_Get_User_By_Username() {
	var username string = "user_test"

	s.T().Run("Succes getting user by username", func(t *testing.T) {
		s.mock.ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "users" WHERE username = $1 ORDER BY "users"."id" LIMIT 1`),
		).WithArgs(username).
			WillReturnRows(sqlmock.NewRows([]string{"id", "username"}).AddRow(1, username))

		res, err := s.repository.GetUserByUsername(username)

		assert.Nil(s.T(), err)
		assert.NotNil(s.T(), res)
	})

	s.T().Run("Requested username not found", func(t *testing.T) {
		username = "user_not_found"
		s.mock.ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "users" WHERE username = $1 ORDER BY "users"."id" LIMIT 1`),
		).WithArgs(username).
			WillReturnError(gorm.ErrRecordNotFound)

		res, err := s.repository.GetUserByUsername(username)

		assert.Nil(s.T(), err)
		assert.Nil(s.T(), res)
	})
}
