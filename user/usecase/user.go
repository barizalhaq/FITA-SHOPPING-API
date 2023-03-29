package usecase

import "github.com/barizalhaq/fita_shopping_api/domain"

type userUsecase struct {
	userRepository domain.UserRepositoryInterface
}

func NewUserUsecase(uRepo domain.UserRepositoryInterface) *userUsecase {
	return &userUsecase{uRepo}
}
