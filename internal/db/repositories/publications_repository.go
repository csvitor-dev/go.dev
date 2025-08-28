package repositories

import (
	"database/sql"

	"github.com/csvitor-dev/social-media/internal/models"
	"github.com/csvitor-dev/social-media/pkg/errors"
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

func (repo *Publications) FindById(id uint64) (models.Publication, error) {
	rows, err := repo.db.Query(`
		SELECT p.*, u.nickname 
		FROM publications p
		INNER JOIN users u
		ON p.author_id = u.id
		WHERE p.id = ?;`, id,
	)

	if err != nil {
		return models.Publication{}, err
	}
	defer rows.Close()

	if !rows.Next() {
		return models.Publication{}, errors.ErrModelNotFound
	}
	var pub models.Publication

	err = rows.Scan(
		&pub.Id,
		&pub.Title,
		&pub.Content,
		&pub.Likes,
		&pub.AuthorId,
		&pub.CreatedOn,
		&pub.UpdatedOn,
		&pub.AuthorNickName,
	)
	return pub, err
}
