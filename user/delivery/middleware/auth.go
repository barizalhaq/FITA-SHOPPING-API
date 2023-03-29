package middleware

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/barizalhaq/fita_shopping_api/config"
	"github.com/barizalhaq/fita_shopping_api/domain"
	"github.com/barizalhaq/fita_shopping_api/user/delivery/auth"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func InitMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cookieAccess := auth.AuthCookieAccess{
			GinContext: ctx,
		}

		authCtx := context.WithValue(ctx.Request.Context(), auth.AuthenticationContext, &cookieAccess)
		ctx.Request = ctx.Request.WithContext(authCtx)

		ctx.Next()
	}
}

func Authenticated(userRepo domain.UserRepositoryInterface) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cookieVal, err := ctx.Cookie(auth.AuthorizationCookieKey)
		if err != nil {
			if errors.Is(http.ErrNoCookie, err) {
				ctx.Next()
				return
			}
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		signedTokenString, err := base64.URLEncoding.DecodeString(cookieVal)
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if len(signedTokenString) < sha256.Size {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		signature := signedTokenString[:sha256.Size]
		tokenString := string(signedTokenString[sha256.Size:])

		hash := hmac.New(sha256.New, []byte(config.C.AppSecret))
		hash.Write([]byte(auth.AuthorizationCookieKey))
		hash.Write([]byte(tokenString))
		expectedSignature := hash.Sum(nil)

		if !hmac.Equal([]byte(signature), expectedSignature) {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header)
			}

			return []byte(config.C.AppSecret), nil
		})
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if float64(time.Now().Unix()) > claims["exp"].(float64) {
				ctx.AbortWithStatus(http.StatusUnauthorized)
				return
			}

			userID := uint64(claims["user_id"].(float64))
			user, err := userRepo.GetUserByID(userID)
			if err != nil {
				ctx.AbortWithError(http.StatusInternalServerError, err)
				return
			}
			if user == nil {
				ctx.AbortWithStatus(http.StatusUnauthorized)
				return
			}

			ctxWithAuthenticatedUser := context.WithValue(ctx.Request.Context(), auth.AuthenticatedUserContext, user)
			ctx.Request = ctx.Request.WithContext(ctxWithAuthenticatedUser)
			ctx.Next()
			return
		}

		ctx.AbortWithStatus(http.StatusUnauthorized)
	}
}
