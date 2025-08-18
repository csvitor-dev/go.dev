package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/csvitor-dev/social-media/internal/config"
	"github.com/dgrijalva/jwt-go"
)

func CreateToken(userId uint64) (string, error) {
	claims := jwt.MapClaims{
		"user_id":    userId,
		"authorized": true,
		"exp":        time.Now().Add(time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(config.Env.SECRET_KEY)
}

func ValidateToken(token string) error {
	refinedToken, err := jwt.Parse(
		token,
		func(t *jwt.Token) (any, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("auth: signing method mismatch: %v", t.Method.Alg())
			}
			return config.Env.SECRET_KEY, nil
		})

	if err != nil {
		return err
	}

	if _, ok := refinedToken.Claims.(jwt.MapClaims); !ok || !refinedToken.Valid {
		return errors.New("auth: invalid token")
	}
	return nil
}
