package mock

import (
	"api/domain/book"
	"errors"
)

type BookRepositoryMock struct {
	datas []book.Book
}

func (db *BookRepositoryMock) Get() ([]book.Book, error) {
	return db.datas, nil
}

func (db *BookRepositoryMock) GetById(id string) (book.Book, error) {
	var idx int = -1

	for index, content := range db.datas {
		if id == content.Id {
			idx = index
		}
	}

	if idx < 0 {
		return book.Book{}, errors.New("NOT FOUND")
	}

	return db.datas[idx], nil
}

func (db *BookRepositoryMock) GetByEmail(email string) (book.Book, error) {

	return book.Book{}, nil
}

func (db *BookRepositoryMock) Update(data book.Book) error {
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

func (db *BookRepositoryMock) Create(data book.Book) error {
	db.datas = append(db.datas, data)

	return nil
}

func (db *BookRepositoryMock) Delete(id string) error {
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

func NewBookRepositoryMock(products []book.Book) *BookRepositoryMock {
	return &BookRepositoryMock{datas: products}
}
