package bookcontroller

import (
	"BookStore/model"
	"BookStore/shared"
	"encoding/json"
	"net/http"
	"strconv"
)

func (c *BookController) Update(w http.ResponseWriter, r *http.Request, id string) {
	idi, err := strconv.Atoi(id)
	if err != nil {
		shared.NewErrorResponse(w,http.StatusBadRequest, "Invalid ID")
		return
	}

	var updatedBook model.Book
	err = json.NewDecoder(r.Body).Decode(&updatedBook)
	if err != nil {
		shared.NewErrorResponse(w,http.StatusBadRequest, "Invalid request payload")
		return
	}

	updated, err := model.Update(idi, &updatedBook)
	if err != nil {
		shared.NewErrorResponse(w,http.StatusInternalServerError, err.Error())
		return
	}

	if updated {
		shared.NewResponse(w,http.StatusOK, "Book updated successfully", updatedBook)
		return
	}
	shared.NewErrorResponse(w,http.StatusNotFound, "Book not found")
}