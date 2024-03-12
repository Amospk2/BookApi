package mock

import (
	"api/domain/author"
	"errors"
)

type AuthorRepositoryMock struct {
	datas []author.Author
}

func (db *AuthorRepositoryMock) Get() ([]author.Author, error) {
	return db.datas, nil
}

func (db *AuthorRepositoryMock) GetById(id string) (author.Author, error) {
	var idx int = -1

	for index, content := range db.datas {
		if id == content.Id {
			idx = index
		}
	}

	if idx < 0 {
		return author.Author{}, errors.New("NOT FOUND")
	}

	return db.datas[idx], nil
}

func (db *AuthorRepositoryMock) GetByEmail(email string) (author.Author, error) {

	return author.Author{}, nil
}

func (db *AuthorRepositoryMock) Update(data author.Author) error {
	var idx int = -1

	for index, content := range db.datas {
		if data.Id == content.Id {
			idx = index
		}
	}

	if idx < 0 {
		return errors.New("NOT FOUND")
	}

	db.datas[idx] = data

	return nil
}

func (db *AuthorRepositoryMock) Create(data author.Author) error {
	db.datas = append(db.datas, data)

	return nil
}

func (db *AuthorRepositoryMock) Delete(id string) error {
	var idx int = -1

	for index, content := range db.datas {
		if id == content.Id {
			idx = index
		}
	}

	if idx < 0 {
		return errors.New("NOT FOUND")
	}

	db.datas = append(db.datas[:idx], db.datas[idx+1:]...)

	return nil
}

func NewAuthorRepositoryMock(products []author.Author) *AuthorRepositoryMock {
	return &AuthorRepositoryMock{datas: products}
}
