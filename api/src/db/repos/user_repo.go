package repos

import (
	"database/sql"

	"github.com/csvitor-dev/social-media/api/src/models"
)

// Users: ...
type Users struct {
	db *sql.DB
}

// NewUserRepo: ...
func NewUserRepo(db *sql.DB) *Users {
	return &Users{
		db,
	}
}

// CreateUser: ...
func (repo *Users) CreateUser(user models.User) (uint64, error) {
	statement, err := repo.db.Prepare(
		"INSERT INTO users(name, nickname, email, password) values(?, ?, ?, ?)",
	)

	if err != nil {
		return 0, err
	}
	defer statement.Close()
	result, err := statement.Exec(user.Name, user.Nickname, user.Email, user.Password)

	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}
	return uint64(id), nil
}
