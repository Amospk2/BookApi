package book

type Book struct {
	Id           string  `json:"id,omitempty"`
	Title        string  `json:"title,omitempty"`
	Category     string  `json:"category,omitempty"`
	Subtitle     string  `json:"subtitle,omitempty"`
	Description  string  `json:"description,omitempty"`
	Release_date string  `json:"release_date,omitempty"`
	Publisher    string  `json:"publisher,omitempty"`
	Language     string  `json:"language,omitempty"`
	Author       string  `json:"author,omitempty"`
	Page_number  int32   `json:"page_number,omitempty"`
	Imagem       string  `json:"imagem,omitempty"`
	Rate         float32 `json:"rate,omitempty"`
	Owner        string  `json:"owner,omitempty"`
}

func NewAuthor(
	title string,
	category string,
	subtitle string,
	description string,
	release_date string,
	publisher string,
	language string,
	author string,
	page_number int32,
	imagem string,
	rate float32,
	owner string,
) *Book {
	return &Book{
		Title:        title,
		Category:     category,
		Subtitle:     subtitle,
		Description:  description,
		Release_date: release_date,
		Language:     language,
		Publisher:    publisher,
		Author:       author,
		Page_number:  page_number,
		Imagem:       imagem,
		Rate:         rate,
		Owner:        owner,
	}

}
