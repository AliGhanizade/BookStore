package bookcontroller

import (
	"BookStore/model"
	"encoding/json"
	"net/http"
)

func (c *BookController) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	for _, book := range model.Books {
		err := json.NewEncoder(w).Encode(book)
		if err != nil {
			http.Error(w,err.Error(),http.StatusBadRequest)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	

}
