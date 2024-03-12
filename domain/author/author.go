package author

type Author struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

func NewAuthor(
	name string,
) *Author {
	return &Author{
		Name: name,
	}

}
