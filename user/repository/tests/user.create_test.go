package repository_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/barizalhaq/fita_shopping_api/domain"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func (s *Suite) Test_repository_Create() {
	var (
		username = "user_test"
		password = "password_test"
	)

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 10)

	s.T().Run("Success creating user", func(t *testing.T) {
		s.mock.ExpectQuery(
			regexp.QuoteMeta(`INSERT INTO users (username, password) VALUES ($1,$2) RETURNING id,username`),
		).
			WithArgs(username, string(hashedPassword)).
			WillReturnRows(sqlmock.NewRows([]string{"id", "username"}).AddRow(uint64(1), username))

		res, err := s.repository.Create(username, string(hashedPassword))

		assert.NoError(s.T(), err)
		assert.NotNil(s.T(), res)
		assert.Equal(s.T(), &domain.User{
			ID:       1,
			Username: username,
		}, res)
	})

	s.T().Run("Fail create user", func(t *testing.T) {
		s.mock.ExpectQuery(
			regexp.QuoteMeta(`INSERT INTO users (username, password) VALUES ($1,$2) RETURNING id,username`),
		).
			WithArgs(username, string(hashedPassword)).
			WillReturnError(fmt.Errorf("dummy error"))

		res, err := s.repository.Create(username, string(hashedPassword))
		assert.NotNil(s.T(), err)
		assert.Nil(s.T(), res)
	})
}
