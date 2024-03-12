package category

type Category struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

func NewCategory(
	name string,
) *Category {
	return &Category{
		Name: name,
	}

}
