package bookcontroller

import (
	"BookStore/model"
	"encoding/json"
	"net/http"
)

func (c *BookController) Delete(w http.ResponseWriter, r *http.Request) {
	var requestBody struct {
		ID int `json:"id"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "invalid request body ", http.StatusBadRequest)
		return
	}

	stauts, err := model.Delete(requestBody.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if !stauts {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message":"Book Deleted successfully"})
}
