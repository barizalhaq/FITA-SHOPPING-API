package usecase

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/barizalhaq/fita_shopping_api/config"
	"github.com/barizalhaq/fita_shopping_api/graph/gModel"
	"github.com/barizalhaq/fita_shopping_api/user/delivery/auth"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func (u *userUsecase) Authenticate(input gModel.AuthenticateInput) (string, error) {
	user, err := u.userRepository.GetUserByUsername(input.Username)
	if err != nil {
		log.Fatalf("[userUsecase.Authenticate] userRepository.GetUserByUsername Error: %v", err)
		return "", err
	}
	if user == nil {
		return "", errors.New("the requested username or password is invalid")
	}

	hashedPass, err := u.userRepository.GetPasswordByUser(user)
	if err != nil {
		log.Fatalf("[userUsecase.Authenticate] userRepository.GetPasswordByUser Error: %v", err)
		return "", err
	}
	if len(hashedPass) == 0 {
		return "", errors.New("the requested username or password is invalid")
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(input.Password))
	if err != nil {
		return "", errors.New("the requested username or password is invalid")
	}

	tokenExp := time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     tokenExp,
	})

	stringToken, err := token.SignedString([]byte(config.C.AppSecret))
	if err != nil {
		fmt.Println("HERE", config.C.AppSecret)
		log.Fatalf("[userUsecase.Authenticate] Error: %v", err)
		return "", err
	}

	signaturedToken := auth.WriteSigned(stringToken, []byte(config.C.AppSecret))

	// AuthAccess.GinContext.SetSameSite(http.SameSiteLaxMode)
	// AuthAccess.GinContext.SetCookie(auth.AuthorizationCookieKey, signaturedToken, 3600, "", "", false, true)

	return signaturedToken, nil
}
