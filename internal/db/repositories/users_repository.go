package repositories

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/csvitor-dev/social-media/internal/models"
	"github.com/csvitor-dev/social-media/pkg/errors"
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

// FindAll: retrieves all users from the database
func (repo *Users) FindAll() ([]models.User, error) {
	rows, err := repo.db.Query("SELECT id, name, nickname, email, created_on, updated_on FROM users;")

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users = []models.User{}

	for rows.Next() {
		var user models.User

		if err = rows.Scan(
			&user.Id,
			&user.Name,
			&user.Nickname,
			&user.Email,
			&user.CreatedOn,
			&user.UpdatedOn,
		); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// FindById: retrieves a user by its id.
// It returns an error if not found
func (repo *Users) FindById(id uint64) (models.User, error) {
	rows, err := repo.db.Query(
		"SELECT id, name, nickname, email, created_on, updated_on FROM users WHERE id = ?;", id,
	)

	if err != nil {
		return models.User{}, err
	}
	defer rows.Close()

	if !rows.Next() {
		return models.User{}, errors.ErrUserNotFound
	}
	var user models.User

	err = rows.Scan(
		&user.Id,
		&user.Name,
		&user.Nickname,
		&user.Email,
		&user.CreatedOn,
		&user.UpdatedOn,
	)
	return user, err
}

func (repo *Users) FindByEmail(email string) (models.User, error) {
	rows, err := repo.db.Query(
		"SELECT id, password FROM users WHERE email = ?;", email,
	)

	if err != nil {
		return models.User{}, err
	}
	defer rows.Close()

	if !rows.Next() {
		return models.User{}, errors.ErrUserNotFound
	}
	var user models.User

	err = rows.Scan(
		&user.Id,
		&user.Password,
	)
	return user, err
}

// Create: inserts a new user into the database
func (repo *Users) Create(user models.User) (uint64, error) {
	statement, err := repo.db.Prepare(
		"INSERT INTO users(name, nickname, email, password) VALUES (?, ?, ?, ?);",
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

	return uint64(id), err
}

// Update: updates a user by its id.
func (repo *Users) Update(id uint64, user models.User) error {
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
	_, err = statement.Exec(fields...)

	return err
}

// buildQueryWithValidFields: builds the SQL query and fields for updating a user.
func (repo *Users) buildQueryWithValidFields(user models.User) (string, []any, error) {
	var (
		partials    []string
		validFields []any
	)
	fields := user.ToMap([]string{"name", "nickname", "email"})

	for key, field := range fields {
		if field != "" {
			partials = append(partials, fmt.Sprintf("%s = ?", key))
			validFields = append(validFields, field)
		}
	}

	if len(validFields) == 0 {
		return "", nil, errors.ErrNoFieldsToUpdate
	}
	partials = append(partials, "updated_on = CURRENT_TIMESTAMP()")
	return fmt.Sprintf("UPDATE users SET %s WHERE id = ?;", strings.Join(partials, ", ")), validFields, nil
}

func (repo *Users) Delete(id uint64) error {
	statement, err := repo.db.Prepare("DELETE FROM users WHERE id = ?;")

	if err != nil {
		return err
	}
	defer statement.Close()
	_, err = statement.Exec(id)

	return err
}

func (repo *Users) FindPasswordFromUser(id uint64) (string, error) {
	rows, err := repo.db.Query("SELECT password FROM users WHERE id = ?;", id)

	if err != nil {
		return "", err
	}
	defer rows.Close()

	if !rows.Next() {
		return "", errors.ErrUserNotFound
	}
	var password string
	err = rows.Scan(&password)

	return password, err
}

func (repo *Users) RefreshPasswordFromUser(id uint64, newPassword string) error {
	statement, err := repo.db.Prepare("UPDATE users SET password = ? WHERE id = ?;")

	if err != nil {
		return err
	}
	defer statement.Close()
	_, err = statement.Exec(newPassword, id)

	return err
}
