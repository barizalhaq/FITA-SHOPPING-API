package auth

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"

	"github.com/gin-gonic/gin"
)

type contextKey int

const (
	AuthenticationContext    contextKey = iota
	AuthenticatedUserContext contextKey = iota
	AuthorizationCookieKey   string     = "Authorization"
)

type AuthCookieAccess struct {
	GinContext *gin.Context
}

func WriteSigned(token string, secret []byte) string {
	hash := hmac.New(sha256.New, secret)
	hash.Write([]byte(AuthorizationCookieKey))
	hash.Write([]byte(token))

	signature := hash.Sum(nil)

	signatured := base64.URLEncoding.EncodeToString([]byte(string(signature) + token))

	return signatured
}

func GetCookieAccess(ctx context.Context) *AuthCookieAccess {
	return ctx.Value(AuthenticationContext).(*AuthCookieAccess)
}
