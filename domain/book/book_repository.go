package book

type BookRepository interface {
	Get() ([]Book, error)
	GetById(id string) (Book, error)
	Update(data Book) error
	Create(data Book) error
	Delete(id string) error
}
