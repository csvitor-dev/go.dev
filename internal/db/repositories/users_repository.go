package repos

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/csvitor-dev/social-media/internal/models"
)

// Users: Users repository interface
type Users struct {
	db *sql.DB
}

// NewUsersRepository: creates a new instance of Users repository
func NewUsersRepository(db *sql.DB) *Users {
	return &Users{
		db,
	}
}

// GetUsers: retrieves all users from the database
func (repo *Users) GetUsers() ([]models.User, error) {
	rows, err := repo.db.Query("SELECT id, name, nickname, email, created_on FROM users;")

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users = []models.User{}

	for rows.Next() {
		var user models.User

		if err = rows.Scan(&user.Id, &user.Name, &user.Nickname, &user.Email, &user.CreatedOn); err != nil {
			return nil, err
		}

		users = append(users, user)
	}
	return users, nil
}

// GetById: retrieves a user by its id.
// It returns an error if not found
func (repo *Users) GetById(id uint64) (models.User, error) {
	rows, err := repo.db.Query(
		"SELECT id, name, nickname, email, created_on FROM users WHERE id = ?", id,
	)

	if err != nil {
		return models.User{}, err
	}
	var user models.User

	if rows.Next() {
		if err := rows.Scan(
			&user.Id,
			&user.Name,
			&user.Nickname,
			&user.Email,
			&user.CreatedOn,
		); err != nil {
			return models.User{}, err
		}
	}
	return user, nil
}

// CreateUser: inserts a new user into the database
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

// UpdateUserById: updates a user by its id.
func (repo *Users) UpdateUserById(id uint64, user models.User) error {
	query, fields, err := repo.buildQueryWithValidFields(user)

	if err != nil {
		return err
	}
	statement, err := repo.db.Prepare(query)
	fields = append(fields, id)

	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(fields...); err != nil {
		return err
	}
	return nil
}

// buildQueryWithValidFields: builds the SQL query and fields for updating a user.
func (repo *Users) buildQueryWithValidFields(user models.User) (string, []any, error) {
	var (
		partials    []string
		validFields []any
	)
	fields := user.ToMap()

	for key, field := range fields {
		if key != "password" && field != "" {
			partials = append(partials, fmt.Sprintf("%s = ?", key))
			validFields = append(validFields, field)
		}
	}
	if len(validFields) == 0 {
		return "", nil, errors.New("no fields to update")
	}

	return fmt.Sprintf("UPDATE users SET %s WHERE id = ?;", strings.Join(partials, ", ")), validFields, nil
}
