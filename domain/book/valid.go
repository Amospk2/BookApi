package book

import (
	"fmt"
	"time"
)

func (book Book) Valid() bool {

	if len(book.Title) == 0 || book.Title == "" {
		return false
	}

	if len(book.Category) == 0 || book.Category == "" {
		return false
	}

	if len(book.Subtitle) == 0 || book.Subtitle == "" {
		return false
	}

	if len(book.Description) == 0 || book.Description == "" {
		return false
	}

	if _, err := time.Parse("2006-01-02", book.Release_date); err != nil {
		return false
	}

	if len(book.Publisher) == 0 || book.Publisher == "" {
		return false
	}

	if len(book.Language) == 0 || book.Language == "" {
		return false
	}

	if len(book.Author) == 0 || book.Author == "" {
		return false
	}

	if fmt.Sprintf("%T", book.Page_number) != "int" && book.Page_number < 0 {
		return false
	}

	if fmt.Sprintf("%T", book.Rate) != "float" && book.Rate < 0 {
		return false
	}

	if len(book.Owner) == 0 || book.Owner == "" {
		return false
	}

	return true
}
