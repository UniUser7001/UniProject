package handlers

import (
	"encoding/json"
	"go_api/models"
	"go_api/utils"
	"net/http"
	"strconv"
	"strings"
)

var books []models.Book
var nextID = 1

func BooksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getBooks(w, r)
	case http.MethodPost:
		createBook(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func BookHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/books/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case http.MethodGet:
		getBookByID(w, r, id)
	case http.MethodPut:
		updateBook(w, r, id)
	case http.MethodDelete:
		deleteBook(w, r, id)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	utils.RespondJSON(w, http.StatusOK, books)
}

func getBookByID(w http.ResponseWriter, r *http.Request, id int) {
	for _, book := range books {
		if book.ID == id {
			utils.RespondJSON(w, http.StatusOK, book)
			return
		}
	}
	http.NotFound(w, r)
}

func createBook(w http.ResponseWriter, r *http.Request) {
	var newBook models.Book
	if err := json.NewDecoder(r.Body).Decode(&newBook); err != nil {
		utils.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	newBook.ID = nextID
	nextID++
	books = append(books, newBook)
	utils.RespondJSON(w, http.StatusCreated, newBook)
}

func updateBook(w http.ResponseWriter, r *http.Request, id int) {
	var updatedBook models.Book
	if err := json.NewDecoder(r.Body).Decode(&updatedBook); err != nil {
		utils.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	for i, book := range books {
		if book.ID == id {
			books[i] = updatedBook
			books[i].ID = book.ID
			utils.RespondJSON(w, http.StatusOK, books[i])
			return
		}
	}
	http.NotFound(w, r)
}

func deleteBook(w http.ResponseWriter, r *http.Request, id int) {
	for i, book := range books {
		if book.ID == id {
			books = append(books[:i], books[i+1:]...)
			utils.RespondJSON(w, http.StatusOK, map[string]string{"message": "Book deleted"})
			return
		}
	}
	http.NotFound(w, r)
}
