package auth

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

func GetUserIdFromToken() (uint64, error) {
	claims, err := validateClaims()

	if err != nil {
		return 0, err
	}
	id, ok := claims["user_id"].(uint64)

	if !ok {
		return 0, err
	}
	return id, nil
}

func validateClaims() (jwt.MapClaims, error) {
	claims, ok := scopeToken.Claims.(jwt.MapClaims)

	if !ok {
		return jwt.MapClaims{}, errors.New("auth: invalid token formulation")
	}
	return claims, nil
}
