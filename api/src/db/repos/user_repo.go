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

func (repo *Users) GetUsers() ([]models.User, error) {
	rows, err := repo.db.Query("SELECT id, name, nickname, email, created_on FROM users;")

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []models.User

	for rows.Next() {
		var user models.User
		
		if err = rows.Scan(&user.Id, &user.Name, &user.Nickname, &user.Email, &user.CreatedOn); err != nil {
			return nil, err
		}

		users = append(users, user)
	}
	return users, nil
}

// CreateUser: ...
func (repo *Users) CreateUser(user models.User) (uint64, error) {
	statement, err := repo.db.Prepare(
		"INSERT INTO users(name, nickname, email, password) VALUES(?, ?, ?, ?)",
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
