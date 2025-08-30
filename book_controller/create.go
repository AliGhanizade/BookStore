package bookcontroller

import (
	"BookStore/model"
	"encoding/json"
	"net/http"
)

func (c *BookController) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	var newbook model.Book

	err := json.NewDecoder(r.Body).Decode(&newbook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	model.Create(newbook)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Book Added in slice"})

}
