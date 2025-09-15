package model

import (
	cnf "BookStore/config"
	"BookStore/shared"
	"encoding/json"
	"os"

)

type Book struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Author     string `json:"author"`
	CreateYear string `json:"create_year"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

//	{
//		"name" : "book1",
//		"author" : "ali",
//		"create_year" : "2012"
//	}
var Books []Book

func GetAll() error {
	file, err := os.Open(cnf.FilePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	err = json.NewDecoder(file).Decode(&Books)
	if err != nil {
		return err
	}
	return nil
}

func GetByID(id int) (Book, bool) {
	err := GetAll()
	if err != nil {
		return Book{}, false
	}
	for _, book := range Books {
		if id == book.ID {
			return book, true
		}
	}
	return Book{}, false
}

func Create(newBook *Book) error {
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

	newBook.CreatedAt = shared.GetCurrentTime()
	newBook.UpdatedAt = shared.GetCurrentTime()

	Books = append(Books, *newBook)

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

func Update(id int,upBook *Book) (bool, error) {

	err := GetAll()
	if err != nil {
		return false, err
	}

	updated := false
	for i, book := range Books {
		if book.ID == id {
			upBook.ID = id
			upBook.UpdatedAt = shared.GetCurrentTime()
			upBook.CreatedAt = book.CreatedAt
			Books[i] = *upBook
			updated = true
			break
		}
	}
	if !updated {
		return false, nil
	}

	file, err := os.OpenFile(cnf.FilePath, os.O_RDWR|os.O_TRUNC, 0644)
	if err != nil {
		return false, err
	}
	defer file.Close()

	jsonBooks, err := json.Marshal(Books)
	if err != nil {
		return false, err
	}

	_, err = file.Write(jsonBooks)
	if err != nil {
		return false, err
	}

	return true, nil
}