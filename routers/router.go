package routers

import (
	bookcontroller "BookStore/book_controller"
	"net/http"
)

func Run() {
	bookC := &bookcontroller.BookController{}
	r := http.NewServeMux()

	r.HandleFunc("/book/create", bookC.Create)
	r.HandleFunc("/book/items" , bookC.GetAllBooks)


	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
	

}
