package routers

import (
	bookcontroller "BookStore/book_controller"
	"BookStore/shared"
	"fmt"
	"net/http"
)

func Run() {
	bookC := &bookcontroller.BookController{}
	r := http.NewServeMux()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
			
		case http.MethodGet:
			query := r.URL.Query()
			if conditions, ok := query["id"]; ok && len(conditions) > 0 {
				bookC.GetByID(w, r, conditions[0])
			} else {
				bookC.GetAll(w, r)
			}

		case http.MethodPost:
			bookC.Create(w, r)

		case http.MethodDelete:
			bookC.Delete(w, r)

		case http.MethodPut:
			query := r.URL.Query()
			if conditions, ok := query["id"]; ok && len(conditions) > 0 {
				bookC.Update(w, r, conditions[0])
			} else {
				shared.NewErrorResponse(w,http.StatusBadRequest, "ID is required")
				return
			}

		default:
			shared.NewErrorResponse(w,http.StatusMethodNotAllowed, "Method not allowed")
			return
		}
	})

	fmt.Println("Server listen on http://127.0.0.1:8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}
