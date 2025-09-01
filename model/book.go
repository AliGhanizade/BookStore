package model

type Book struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Author     string `json:"author"`
	CreateYear string `json:"create_year"`
	// CreateTime string `json:"create_time"`
}

// {
// 	"id" : "1",
// 	"name" : "book1",
// 	"author" : "ali",
// 	"create_year" : "2012"
// }
var Books []Book

func Create(newBook Book) {
	if len(Books) == 0 {
		newBook.ID = 1
	} else {

		newBook.ID = Books[len(Books)-1].ID + 1
	}

	Books = append(Books, newBook)
}
