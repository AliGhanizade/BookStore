package bookcontroller

import (
	"BookStore/model"
	"BookStore/shared"
	"encoding/json"
	"net/http"
)

func (c *BookController) Create(w http.ResponseWriter, r *http.Request) {
	var newBook model.Book

	err := json.NewDecoder(r.Body).Decode(&newBook)
	if err != nil {
		shared.NewErrorResponse(w,http.StatusBadRequest, "Invalid request payload")
		return
	}

	err = model.Create(&newBook)
	if err != nil {
		shared.NewErrorResponse(w,http.StatusInternalServerError, err.Error())
		return
	}

	shared.NewResponse(w,http.StatusCreated, "Book created successfully", newBook)

}
