package bookcontroller

import (
	"BookStore/model"
	"BookStore/shared"
	"net/http"
	"strconv"
)

func (c *BookController) GetAll(w http.ResponseWriter, r *http.Request) {
	err := model.GetAll()

	if err != nil {
		shared.NewErrorResponse(w,http.StatusInternalServerError, err.Error())
		return
	}

	shared.NewResponse(w,http.StatusOK, "Books retrieved successfully", model.Books)
	
}

func (c *BookController) GetByID(w http.ResponseWriter, r *http.Request, id string) {
	idi, err := strconv.Atoi(id)
	if err != nil {
		shared.NewErrorResponse(w,http.StatusBadRequest, "Invalid ID")
		return
	}

	book, found := model.GetByID(idi)

	if found {
		shared.NewResponse(w,http.StatusOK, "Book found", book)
		return
	}

	shared.NewErrorResponse(w,http.StatusNotFound, "Book not found")
}