package database

import (
	"api/domain/author"
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type AuthorRepository struct {
	pool *pgxpool.Pool
}

func (db *AuthorRepository) Get() ([]author.Author, error) {

	authors := make([]author.Author, 0)

	rows, err := db.pool.Query(context.Background(),
		`select id, name
		from public.authors`,
	)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var author author.Author

		err = rows.Scan(
			&author.Id,
			&author.Name,
		)

		if err != nil {
			log.Fatal(err)
		}

		authors = append(authors, author)
	}

	return authors, nil
}

func (db *AuthorRepository) GetById(id string) (author.Author, error) {

	var authorFinded author.Author

	err := db.pool.QueryRow(
		context.Background(),
		`select 
			id, name
		from 
			public.authors 
		where
			id=$1`,
		id,
	).Scan(
		&authorFinded.Id,
		&authorFinded.Name,
	)

	if err != nil {
		return author.Author{}, err
	}

	return authorFinded, nil
}

func (db *AuthorRepository) Update(data author.Author) error {
	_, err := db.pool.Exec(
		context.Background(),
		`UPDATE authors 
		SET name = $1
		WHERE id = $2`,
		data.Name, data.Id,
	)

	if err != nil {
		return err
	}

	return nil
}

func (db *AuthorRepository) Create(data author.Author) error {
	_, err := db.pool.Exec(
		context.Background(), "INSERT INTO authors VALUES($1,$2)",
		data.Id, data.Name,
	)

	if err != nil {
		return err
	}

	return nil
}

func (db *AuthorRepository) Delete(id string) error {

	_, err := db.pool.Exec(context.Background(), "DELETE FROM authors WHERE id=$1", id)

	if err != nil {
		return err
	}

	return nil
}

func NewAuthorRepository(pool *pgxpool.Pool) *AuthorRepository {
	return &AuthorRepository{pool: pool}
}
