package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/csvitor-dev/go.dev/internal/config/env"
	"github.com/csvitor-dev/go.dev/internal/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

var scopeToken *jwt.Token
var currentJti string

func CreateToken(user models.User, duration time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"user_id":    user.Id,
		"authorized": true,
		"jti":        uuid.NewString(),
		"exp":        time.Now().Add(duration).Unix(),
	}

	if jti, ok := claims["jti"].(string); ok {
		currentJti = jti
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	scopeToken = token

	return token.SignedString(env.Env.SECRET_KEY)
}

func ValidateToken(token string) error {
	refinedToken, err := jwt.Parse(
		token,
		func(t *jwt.Token) (any, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("auth: signing method mismatch: %v", t.Method.Alg())
			}
			return env.Env.SECRET_KEY, nil
		})

	if err != nil {
		return err
	}
	claims, ok := refinedToken.Claims.(jwt.MapClaims)

	if !ok || !refinedToken.Valid {
		return errors.New("auth: invalid token")
	}
	jti, _ := claims["jti"].(string)

	if !isCurrentToken(jti) {
		return errors.New("auth: token has been invalidate")
	}
	return nil
}

func isCurrentToken(jti string) bool {
	return jti == currentJti
}

func InvalidateToken() {
	currentJti = ""
}
