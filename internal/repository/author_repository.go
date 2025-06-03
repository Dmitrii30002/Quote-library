package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/Dmitrii30002/Quote-library/internal/models"
)

type AuthorRepository struct {
	db *sql.DB
}

func NewAuthorRepository(db *sql.DB) *AuthorRepository {
	return &AuthorRepository{db: db}
}

func (r *AuthorRepository) Create(author *models.Author) error {
	query := `
		INSERT INTO Author (author_name)
		VALUES ($1)
	`

	_, err := r.db.Exec(query,
		author.Name,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *AuthorRepository) GetAll(page int, limit int) ([]*models.Author, error) {
	offset := (page - 1) * limit

	query := `
        SELECT id, author_name, created_at 
        FROM Author
        ORDER BY created_at DESC
        LIMIT $1 OFFSET $2
    `

	rows, err := r.db.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var authors []*models.Author
	for rows.Next() {
		var author models.Author
		err := rows.Scan(&author.ID, &author.Name, author.Created_at)
		if err != nil {
			return nil, err
		}
		authors = append(authors, &author)
	}

	return authors, nil
}

func (r *AuthorRepository) GetByID(id int) (*models.Author, error) {
	query := `
		SELECT id, author_name, created_at
		FROM Author
		WHERE id = $1
	`
	var author models.Author

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&author.ID,
		&author.Name,
		&author.Created_at,
	)

	if err != nil {
		return nil, err
	}

	return &author, nil
}

func (r *AuthorRepository) GetByName(name string) (*models.Author, error) {
	query := `
		SELECT id, author_name, created_at
		FROM Author
		WHERE author_name= $1
	`
	var author models.Author

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := r.db.QueryRowContext(ctx, query, name).Scan(
		&author.ID,
		&author.Name,
		&author.Created_at,
	)

	if err != nil {
		return nil, err
	}

	return &author, nil
}

func (r *AuthorRepository) Update(author models.Author) error {
	query := `
		UPDATE Author
		SET id = $1
			author_name = $2
			created_at = $3
	`

	_, err := r.db.Exec(query, author.ID, author.Name, author.Created_at)

	if err != nil {
		return err
	}

	return nil
}

func (r *AuthorRepository) Delete(id int) error {
	query := `
		DELETE FROM Author
		WHERE id = $1
	`

	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
