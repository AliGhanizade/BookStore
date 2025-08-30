package model

type Book struct{
	Id int `json:"id"`
	Name string `json:"name"`
	Author string `json:"author"`
	CreateYear string `json:"create_year"`
	// CreateTime string `json:"create_time"`
}
// {
// 	"id" : "1",
// 	"name" : "book1",
// 	"author" : "ali",
// 	"create_year" : "2012",
// }
var Books []Book

func Create(newbook Book)  {
	Books = append(Books, newbook) 
}
