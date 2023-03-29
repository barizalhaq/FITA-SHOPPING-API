package usecase_test

import (
	"testing"

	"github.com/barizalhaq/fita_shopping_api/domain"
	"github.com/barizalhaq/fita_shopping_api/domain/mocks"
	"github.com/barizalhaq/fita_shopping_api/graph/gModel"
	"github.com/barizalhaq/fita_shopping_api/user/usecase"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestAuthenticate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockUserRepositoryInterface(ctrl)
	userUC := usecase.NewUserUsecase(m)

	dummyInput := gModel.AuthenticateInput{
		Username: "username_test",
		Password: "user_test_password",
	}

	t.Run("User successfully authenticated", func(t *testing.T) {
		m.EXPECT().GetUserByUsername(dummyInput.Username).
			Return(&domain.User{
				ID:       1,
				Username: dummyInput.Username,
			}, nil).Times(1)

		hashedPass, _ := bcrypt.GenerateFromPassword([]byte(dummyInput.Password), 10)
		m.EXPECT().GetPasswordByUser(&domain.User{
			ID:       1,
			Username: dummyInput.Username,
		}).Return(string(hashedPass), nil).Times(1)

		signaturedToken, err := userUC.Authenticate(dummyInput)
		assert.Nil(t, err)
		assert.NotEmpty(t, signaturedToken)
	})
}
