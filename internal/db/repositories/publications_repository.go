package repositories

import (
	"database/sql"

	"github.com/csvitor-dev/social-media/internal/models"
)

type Publications struct {
	db *sql.DB
}

// NewPublicationsRepository: creates a new instance of Publications repository
func NewPublicationsRepository(db *sql.DB) *Publications {
	return &Publications{
		db,
	}
}

func (repo *Publications) Create(publication models.Publication) (uint64, error) {
	statement, err := repo.db.Prepare(
		"INSERT INTO publications(title, content, author_id) VALUES (?, ?, ?);",
	)

	if err != nil {
		return 0, err
	}
	defer statement.Close()
	result, err := statement.Exec(
		publication.Title,
		publication.Content,
		publication.AuthorId,
	)

	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()

	return uint64(id), err
}
