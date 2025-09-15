package bookcontroller

import (
	"BookStore/model"
	"encoding/json"
	"strconv"
	"net/http"
)

func (c *BookController) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	for _, book := range model.Books {
		err := json.NewEncoder(w).Encode(book)
		if err != nil {
			http.Error(w,err.Error(),http.StatusBadRequest)
		}
	}	

}

func (c *BookController) GetByID(w http.ResponseWriter, r *http.Request, id string) {
	idint, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	for _, book := range model.Books {
		if idint == book.ID {
			err := json.NewEncoder(w).Encode(book)
			if err != nil {
				http.Error(w,err.Error(),http.StatusBadRequest)
			}
			return
		}
	}
	http.Error(w, "Book not found", http.StatusNotFound)
}