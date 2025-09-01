package routers

import (
	bookcontroller "BookStore/book_controller"
	"fmt"
	"net/http"
)

func Run() {
	bookC := &bookcontroller.BookController{}
	r := http.NewServeMux()

	r.HandleFunc("/book/create", bookC.Create)
	r.HandleFunc("/book/item", bookC.GetAllBooks)
	r.HandleFunc("/book/delete", bookC.Delete)

	fmt.Println("Server listen on 127.0.0.1:8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}
