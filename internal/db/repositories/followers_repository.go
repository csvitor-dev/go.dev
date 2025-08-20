package repositories

import "database/sql"

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
