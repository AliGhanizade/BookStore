package bookcontroller

import (
	"BookStore/model"
	"encoding/json"
	"net/http"
)

func (c *BookController) Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var requestBody struct {
		ID string `json:"id"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "invalid request body ", http.StatusBadRequest)
		return
	}
	bookFound := false
	var updateBooks []model.Book

	for _, book := range model.Books {
		if book.ID == requestBody.ID {
			continue
		}
		updateBooks = append(updateBooks, book)
	}
	if !bookFound {
		http.Error(w, "Book not found ", http.StatusNotFound)
		return
	}
	model.Books = updateBooks

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message":"Book Deleted successfully"})
}
