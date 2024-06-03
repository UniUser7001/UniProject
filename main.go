package main

import (
	"fmt"
	"go_api/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/books", handlers.BooksHandler)
	http.HandleFunc("/books/", handlers.BookHandler)

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
