package models

import "time"

// User: the model represents the 'User' entity mapped from the database
type User struct {
	ID        uint64 `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Nickname  string `json:"nickname,omitempty"`
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
	CreatedOn time.Time `json:"created_on,omitempty"`
}