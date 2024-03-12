package database

import (
	"api/domain/book"
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type BookRepository struct {
	pool *pgxpool.Pool
}

func (db *BookRepository) Get() ([]book.Book, error) {

	books := make([]book.Book, 0)

	rows, err := db.pool.Query(context.Background(),
		`select id, name
		from public.books`,
	)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var book book.Book

		err = rows.Scan(
			&book.Id,
			&book.Title,
			&book.Category,
			&book.Subtitle,
			&book.Description,
			&book.Release_date,
			&book.Language,
			&book.Author,
			&book.Page_number,
			&book.Imagem,
			&book.Rate,
			&book.Owner,
		)

		if err != nil {
			log.Fatal(err)
		}

		books = append(books, book)
	}

	return books, nil
}

func (db *BookRepository) GetById(id string) (book.Book, error) {

	var bookFinded book.Book

	err := db.pool.QueryRow(
		context.Background(),
		`select 
			id, title, category, subtitle, description, release_date,
			publisher, language, author, page_number, imagem, rate, owner
		from 
			public.books 
		where
			id=$1`,
		id,
	).Scan(
		&bookFinded.Id,
		&bookFinded.Title,
		&bookFinded.Category,
		&bookFinded.Subtitle,
		&bookFinded.Description,
		&bookFinded.Release_date,
		&bookFinded.Language,
		&bookFinded.Author,
		&bookFinded.Page_number,
		&bookFinded.Imagem,
		&bookFinded.Rate,
		&bookFinded.Owner,
	)

	if err != nil {
		return book.Book{}, err
	}

	return bookFinded, nil
}

func (db *BookRepository) Update(data book.Book) error {
	_, err := db.pool.Exec(
		context.Background(),
		`UPDATE books 
		SET title = $1, category = $2, subtitle = $3, description = $4,
		release_date = $5, language = $6, author = $7, page_number = $8,
		imagem = $9, rate = $10, owner = $11
		WHERE id = $12`,
		&data.Title,
		&data.Category,
		&data.Subtitle,
		&data.Description,
		&data.Release_date,
		&data.Language,
		&data.Author,
		&data.Page_number,
		&data.Imagem,
		&data.Rate,
		&data.Owner,
		&data.Id,
	)

	if err != nil {
		return err
	}

	return nil
}

func (db *BookRepository) Create(data book.Book) error {
	_, err := db.pool.Exec(
		context.Background(), "INSERT INTO books VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)",
		&data.Id,
		&data.Title,
		&data.Category,
		&data.Subtitle,
		&data.Description,
		&data.Release_date,
		&data.Language,
		&data.Author,
		&data.Page_number,
		&data.Imagem,
		&data.Rate,
		&data.Owner,
	)

	if err != nil {
		return err
	}

	return nil
}

func (db *BookRepository) Delete(id string) error {

	_, err := db.pool.Exec(context.Background(), "DELETE FROM books WHERE id=$1", id)

	if err != nil {
		return err
	}

	return nil
}

func NewBookRepository(pool *pgxpool.Pool) *BookRepository {
	return &BookRepository{pool: pool}
}
