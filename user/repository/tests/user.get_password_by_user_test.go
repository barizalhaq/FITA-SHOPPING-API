package repository_test

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/barizalhaq/fita_shopping_api/domain"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func (s *Suite) Test_repository_Get_Password_By_User() {
	var user = domain.User{
		ID:       1,
		Username: "user_test",
	}

	s.T().Run("Success getting password", func(t *testing.T) {
		s.mock.ExpectQuery(
			regexp.QuoteMeta(`SELECT password FROM users WHERE id = $1`),
		).WillReturnRows(sqlmock.NewRows([]string{"password"}).AddRow("some_password"))

		pass, err := s.repository.GetPasswordByUser(&user)

		assert.Nil(s.T(), err)
		assert.NotEmpty(s.T(), pass)
	})

	s.T().Run("Requested user password is not found", func(t *testing.T) {
		s.mock.ExpectQuery(
			regexp.QuoteMeta(`SELECT password FROM users WHERE id = $1`),
		).WillReturnError(gorm.ErrRecordNotFound)

		pass, err := s.repository.GetPasswordByUser(&user)

		assert.Nil(s.T(), err)
		assert.Empty(s.T(), pass)
	})
}
