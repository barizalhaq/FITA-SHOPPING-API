package usecase

import (
	"errors"
	"log"

	"github.com/barizalhaq/fita_shopping_api/domain"
	"github.com/barizalhaq/fita_shopping_api/graph/gModel"
	"golang.org/x/crypto/bcrypt"
)

func (u *userUsecase) Register(input gModel.RegisterInput) (*domain.User, error) {
	existingUser, err := u.userRepository.GetUserByUsername(input.Username)
	if err != nil {
		log.Fatalf("[userUsecase.Register] userRepository.GetUserByUsername Error: %v", err)
		return nil, err
	}

	if existingUser != nil {
		return nil, errors.New("cannot proceed, the username has been taken")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), 10)
	if err != nil {
		log.Fatalf("[userUsecase.Register] Error: %v", err)
		return nil, err
	}

	return u.userRepository.Create(input.Username, string(hashedPassword))
}
