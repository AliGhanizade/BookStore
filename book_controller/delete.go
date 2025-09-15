package bookcontroller

import (
	"BookStore/model"
	"BookStore/shared"
	"encoding/json"
	"net/http"
)

func (c *BookController) Delete(w http.ResponseWriter, r *http.Request) {
	var requestBody struct {
		ID int `json:"id"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		shared.NewErrorResponse(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	stauts, err := model.Delete(requestBody.ID)
	if err != nil {
		shared.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	if !stauts {
		shared.NewErrorResponse(w, http.StatusNotFound, "Book not found")
		return
	}
	
	shared.NewResponse(w, http.StatusOK, "Book deleted successfully", nil)
}
