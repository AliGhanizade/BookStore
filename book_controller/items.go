package bookcontroller

import (
	"BookStore/model"
	"encoding/json"
	"net/http"
)


func (c *BookController)GetAllBooks(w http.ResponseWriter , r *http.Request)  {
	if r.Method == "GET" {

	}

	err := json.NewEncoder(w).Encode(model.Books)
	if err != nil {
		return
	}
	
}