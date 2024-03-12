package author

type AuthorRepository interface {
	Get() ([]Author, error)
	GetById(id string) (Author, error)
	Update(data Author) error
	Create(data Author) error
	Delete(id string) error
}
