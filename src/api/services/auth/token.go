package auth

import (
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
