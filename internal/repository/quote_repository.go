package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/Dmitrii30002/Quote-library/internal/models"
)

type QuoteRepository struct {
	db *sql.DB
}

func NewQuoteRepository(db *sql.DB) *QuoteRepository {
	return &QuoteRepository{db: db}
}

func (r *QuoteRepository) Create(quote *models.Quote) error {
	query := `
		INSERT INTO Quote (quote_text, author_id)
		VALUES ($1, $2)
	`

	_, err := r.db.Exec(query,
		quote.Text, quote.Author_ID,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *QuoteRepository) GetAll() ([]*models.Quote, error) {
	query := `
        SELECT id, quote_text, author_id, created_at 
        FROM Quote
        ORDER BY created_at DESC
    `

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var quotes []*models.Quote
	for rows.Next() {
		var quote models.Quote
		err := rows.Scan(&quote.ID, &quote.Text, &quote.Author_ID, &quote.Created_at)
		if err != nil {
			return nil, err
		}
		quotes = append(quotes, &quote)
	}

	return quotes, nil
}

func (r *QuoteRepository) GetByID(id int) (*models.Quote, error) {
	query := `
		SELECT id, quote_text, author_id, created_at
		FROM Quote
		WHERE id = $1
	`
	var quote models.Quote

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&quote.ID,
		&quote.Text,
		&quote.Author_ID,
		&quote.Created_at,
	)

	if err != nil {
		return nil, err
	}

	return &quote, nil
}

func (r *QuoteRepository) GetRandom() (*models.Quote, error) {
	query := `
		SELECT * 
		FROM Quote 
		WHERE id = (SELECT floor(random() * (SELECT max(id) FROM table_name)) + 1);
	`
	var quote models.Quote

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := r.db.QueryRowContext(ctx, query).Scan(
		&quote.ID,
		&quote.Text,
		&quote.Author_ID,
		&quote.Created_at,
	)

	if err != nil {
		return nil, err
	}

	return &quote, nil
}

func (r *QuoteRepository) GetByAuthorName(name string) ([]*models.Quote, error) {
	query := `
		SELECT id, quote_text, author_id, created_at
		FROM Quote
		LEFT JOIN Author on Author.id = Quote.id
		WHERE Author.name = $1
	`
	rows, err := r.db.Query(query, name)

	if err != nil {
		return nil, err
	}

	var quotes []*models.Quote
	for rows.Next() {
		var quote models.Quote
		err := rows.Scan(&quote.ID, &quote.Text, quote.Author_ID, quote.Created_at)
		if err != nil {
			return nil, err
		}
		quotes = append(quotes, &quote)
	}

	return quotes, nil
}

func (r *QuoteRepository) Update(quote models.Quote) error {
	query := `
		UPDATE Quote
		SET id = $1
			quote_text = $2
			author_id = $3
			created_at = $4
	`

	_, err := r.db.Exec(query, quote.ID, quote.Text, quote.Author_ID, quote.Created_at)

	if err != nil {
		return err
	}

	return nil
}

func (r *QuoteRepository) Delete(id int) error {
	query := `
		DELETE FROM Quote
		WHERE id = $1
	`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
