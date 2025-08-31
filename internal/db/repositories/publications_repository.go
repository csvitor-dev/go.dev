package repositories

import (
	"database/sql"
	"errors"

	"github.com/csvitor-dev/social-media/internal/models"
	errs "github.com/csvitor-dev/social-media/pkg/errors"
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
		return models.Publication{}, errs.ErrModelNotFound
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

func (repo *Publications) SearchPubsByUserId(userId uint64) ([]models.Publication, error) {
	rows, err := repo.db.Query(`
		SELECT DISTINCT p.*, u.nickname
		FROM publications p
		INNER JOIN users u
		ON p.author_id = u.id
		LEFT JOIN followers f
		ON p.author_id = f.user_id
		WHERE p.author_id = ? OR f.follower_id = ?
		ORDER BY 1 DESC;`,
		userId,
		userId,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	publications := make([]models.Publication, 0)

	for rows.Next() {
		var pub models.Publication

		if err = rows.Scan(
			&pub.Id,
			&pub.Title,
			&pub.Content,
			&pub.Likes,
			&pub.AuthorId,
			&pub.CreatedOn,
			&pub.UpdatedOn,
			&pub.AuthorNickName,
		); err != nil {
			return nil, err
		}
		publications = append(publications, pub)
	}
	return publications, nil
}

func (repo *Publications) FilterPubsByUserId(userId uint64) ([]models.Publication, error) {
	rows, err := repo.db.Query(`
		SELECT p.*, u.nickname
		FROM publications p
		INNER JOIN users u
		ON p.author_id = u.id
		WHERE p.author_id = ?
		ORDER BY 1 DESC;`,
		userId,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	publications := make([]models.Publication, 0)

	for rows.Next() {
		var pub models.Publication

		if err = rows.Scan(
			&pub.Id,
			&pub.Title,
			&pub.Content,
			&pub.Likes,
			&pub.AuthorId,
			&pub.CreatedOn,
			&pub.UpdatedOn,
			&pub.AuthorNickName,
		); err != nil {
			return nil, err
		}
		publications = append(publications, pub)
	}
	return publications, nil
}

func (repo *Publications) IsAuthorOfPub(userId, pubId uint64) error {
	pub, err := repo.FindById(pubId)

	if err != nil {
		return err
	}

	if pub.AuthorId != userId {
		return errors.New("controllers: you are not the author of this publication")
	}
	return nil
}

func (repo *Publications) Update(pubId uint64, publication models.Publication) error {
	statement, err := repo.db.Prepare(
		"UPDATE publications SET title = ?, content = ? WHERE id = ?;",
	)

	if err != nil {
		return err
	}
	defer statement.Close()
	_, err = statement.Exec(
		publication.Title,
		publication.Content,
		pubId,
	)

	return err
}

func (repo *Publications) Delete(pubId uint64) error {
	statement, err := repo.db.Prepare(
		"DELETE FROM publications WHERE id = ?;",
	)

	if err != nil {
		return err
	}
	defer statement.Close()
	_, err = statement.Exec(pubId)

	return err
}

func (repo *Publications) Like(pubId uint64) error {
	statement, err := repo.db.Prepare(
		"UPDATE publications SET likes = likes + 1 WHERE id = ?;",
	)

	if err != nil {
		return err
	}
	defer statement.Close()
	result, err := statement.Exec(pubId)

	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()

	if rowsAffected == 0 {
		return errs.ErrModelNotFound
	}
	return err
}

func (repo *Publications) Dislike(pubId uint64) error {
	statement, err := repo.db.Prepare(`
		UPDATE publications
		SET likes = 
		CASE
			WHEN likes > 0
			THEN likes - 1
			ELSE 0
		END
		WHERE id = ?;
	`)

	if err != nil {
		return err
	}
	defer statement.Close()
	result, err := statement.Exec(pubId)

	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()

	if rowsAffected == 0 {
		return errs.ErrModelNotFound
	}
	return err
}
