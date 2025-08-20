package repositories

import (
	"database/sql"

	"github.com/csvitor-dev/social-media/internal/models"
)

// Followers: Followers repository interface
type Followers struct {
	users *Users
	db    *sql.DB
}

// NewFollowersRepository: creates a new instance of Followers repository
func NewFollowersRepository(db *sql.DB) *Followers {
	return &Followers{
		users: NewUsersRepository(db),
		db:    db,
	}
}

func (repo *Followers) Follow(userId, followerId uint64) error {
	_, err := repo.users.FindById(userId)

	if err != nil {
		return err
	}
	statement, err := repo.db.Prepare(
		"INSERT IGNORE INTO followers(user_id, follower_id) VALUES (?, ?);",
	)

	if err != nil {
		return err
	}
	defer statement.Close()
	_, err = statement.Exec(userId, followerId)

	return err
}

func (repo *Followers) Unfollow(userId, followerId uint64) error {
	_, err := repo.users.FindById(userId)

	if err != nil {
		return err
	}
	statement, err := repo.db.Prepare(
		"DELETE FROM followers WHERE (user_id, follower_id) = (?, ?);",
	)

	if err != nil {
		return err
	}
	defer statement.Close()
	_, err = statement.Exec(userId, followerId)

	return err
}

func (repo *Followers) FindFollowersByUserId(userId uint64) ([]models.User, error) {
	rows, err := repo.db.Query(`
		SELECT u.id, u.name, u.nickname, u.email, u.created_on, u.updated_on
		FROM users u
		INNER JOIN followers f
		ON u.id = f.follower_id
		WHERE f.user_id = ?;
	`, userId)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var followers = []models.User{}

	for rows.Next() {
		var follower models.User

		if err = rows.Scan(
			&follower.Id,
			&follower.Name,
			&follower.Nickname,
			&follower.Email,
			&follower.CreatedOn,
			&follower.UpdatedOn,
		); err != nil {
			return nil, err
		}
		followers = append(followers, follower)
	}
	return followers, nil
}

func (repo *Followers) FindFollowingByUserId(userId uint64) ([]models.User, error) {
	rows, err := repo.db.Query(`
		SELECT u.id, u.name, u.nickname, u.email, u.created_on, u.updated_on
		FROM users u
		INNER JOIN followers f
		ON u.id = f.user_id
		WHERE f.follower_id = ?;
	`, userId)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var following = []models.User{}

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
		following = append(following, user)
	}
	return following, nil
}
