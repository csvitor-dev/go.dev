package models

import (
	"strings"
	"time"

	"github.com/csvitor-dev/go.dev/internal/errors"
	"github.com/csvitor-dev/go.dev/internal/security"
)

// User: the model represents the 'User' entity mapped from the database
type User struct {
	Id        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nickname  string    `json:"nickname,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedOn time.Time `json:"created_on,omitzero"`
	UpdatedOn time.Time `json:"updated_on,omitzero"`
}

// NewUser: creates a new user instance, hashing the given password
func NewUser(name, nickname, email string, password ...string) (User, error) {
	if len(password) > 1 {
		return User{}, errors.ErrTooManyFields
	}
	now := time.Now()
	user := User{
		Name:      strings.TrimSpace(name),
		Nickname:  strings.TrimSpace(nickname),
		Email:     strings.TrimSpace(email),
		CreatedOn: now,
		UpdatedOn: now,
	}

	if len(password) == 0 {
		return user, nil
	}
	hashedPassword, err := security.Cryptify(password[0])

	if err != nil {
		return User{}, err
	}
	user.Password = hashedPassword

	return user, nil
}

func (u *User) ToMap(fields []string) map[string]any {
	hook := map[string]any{
		"id":         u.Id,
		"name":       u.Name,
		"nickname":   u.Nickname,
		"email":      u.Email,
		"password":   u.Password,
		"created_on": u.CreatedOn,
		"updated_on": u.UpdatedOn,
	}
	result := make(map[string]any)

	for _, field := range fields {
		if _, exists := hook[field]; exists {
			result[field] = hook[field]
		}
	}
	return result
}
