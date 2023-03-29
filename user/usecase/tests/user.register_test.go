package usecase_test

import (
	"testing"

	"github.com/barizalhaq/fita_shopping_api/domain"
	"github.com/barizalhaq/fita_shopping_api/domain/mocks"
	"github.com/barizalhaq/fita_shopping_api/graph/gModel"
	"github.com/barizalhaq/fita_shopping_api/user/usecase"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockUserRepositoryInterface(ctrl)
	userUC := usecase.NewUserUsecase(m)

	dummyInput := gModel.RegisterInput{
		Username: "dummy_username",
		Password: "dummy_password",
	}

	t.Run("User successfully registered", func(t *testing.T) {
		m.EXPECT().GetUserByUsername(dummyInput.Username).
			Return(nil, nil).
			Times(1)

		m.EXPECT().Create(dummyInput.Username, gomock.Any()).
			Return(&domain.User{
				ID:       1,
				Username: dummyInput.Username,
			}, nil).Times(1)

		user, err := userUC.Register(dummyInput)

		assert.Nil(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, dummyInput.Username, user.Username)
	})

	t.Run("Duplicate username", func(t *testing.T) {
		m.EXPECT().GetUserByUsername(dummyInput.Username).
			Return(&domain.User{
				ID:       1,
				Username: dummyInput.Username,
			}, nil).
			Times(1)

		user, err := userUC.Register(dummyInput)

		assert.NotNil(t, err)
		assert.Nil(t, user)
		assert.Equal(t, err.Error(), "cannot proceed, the username has been taken")
	})
}
