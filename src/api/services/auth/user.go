package auth

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/dgrijalva/jwt-go"
)

func GetUserId() (uint64, error) {
	claims, ok := scopeToken.Claims.(jwt.MapClaims)

	if !ok {
		return 0, errors.New("auth: invalid token formulation")
	}
	id, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["user_id"]), 10, 64)

	if err != nil {
		return 0, err
	}
	return id, nil
}
