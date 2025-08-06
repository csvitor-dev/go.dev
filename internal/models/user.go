package models

import (
	"errors"
	"strings"
	"time"
)

// User: the model represents the 'User' entity mapped from the database
type User struct {
	Id        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nickname  string    `json:"nickname,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedOn time.Time `json:"created_on,omitzero"`
}

// Prepare: Prepare makes validation and formatation of data
func (u *User) Prepare(isRegisterStage bool) []error {
	if errs := u.validate(isRegisterStage); errs != nil {
		return errs
	}

	u.format()
	return nil
}

func (u *User) ToMap() map[string]any {
	return map[string]any{
		"name":     u.Name,
		"nickname": u.Nickname,
		"email":    u.Email,
		"password": u.Password,
	}
}

func (u *User) validate(isRegisterStage bool) []error {
	if !isRegisterStage {
		return []error{}
	}
	var errs []error

	if u.Name == "" {
		errs = append(errs, errors.New("the name field is required"))
	}
	if u.Nickname == "" {
		errs = append(errs, errors.New("the nickname field is required"))
	}
	if u.Email == "" {
		errs = append(errs, errors.New("the email field is required"))
	}
	if u.Password == "" {
		errs = append(errs, errors.New("the password field is required"))
	}
	return errs
}

func (u *User) format() {
	u.Name = strings.TrimSpace(u.Name)
	u.Nickname = strings.TrimSpace(u.Nickname)
	u.Email = strings.TrimSpace(u.Email)
}
