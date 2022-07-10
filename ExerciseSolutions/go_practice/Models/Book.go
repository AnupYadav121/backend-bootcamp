package Models

type Book struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Author string `json:"author"`
	Title  string `json:"title"`
}

type InputBook struct {
	Author string `json:"author"`
	Title  string `json:"title"`
}

type UpdatedBook struct {
	Author string `json:"author"`
	Title  string `json:"title"`
}
