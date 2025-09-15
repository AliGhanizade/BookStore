package model

import (
	cnf "BookStore/config"
	"encoding/json"
	"os"
)

type Book struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Author     string `json:"author"`
	CreateYear string `json:"create_year"`
}

//	{
//		"name" : "book1",
//		"author" : "ali",
//		"create_year" : "2012"
//	}
var Books []Book

func Create(newBook Book) error {
	if len(Books) == 0 {
		newBook.ID = 1
	} else {

		newBook.ID = Books[len(Books)-1].ID + 1
	}

	file, err := os.OpenFile(cnf.FilePath, os.O_RDWR|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	Books = append(Books, newBook)
	jsonBooks, err := json.Marshal(Books)
	if err != nil {
		return err
	}
	_, err = file.Write(jsonBooks)
	if err != nil {
		return err
	}
	return nil
}

func Delete(Id int) (bool,error) {
	found := false
	for i, book := range Books {
		if book.ID == Id {
			Books = append(Books[:i], Books[i+1:]...)
			found = true
			break
		} 
	}
	if !found {
		return found,nil
	}
	file, err := os.OpenFile(cnf.FilePath, os.O_RDWR|os.O_TRUNC, 0644)
	if err != nil {
		return found,err
	}
	defer file.Close()
	jsonBooks, err := json.Marshal(Books)
	if err != nil {
		return found,err
	}
	_, err = file.Write(jsonBooks)
	if err != nil {
		return found,err
	}
	return found,nil
}
