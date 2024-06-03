package handlers

import (
	"bytes"
	"encoding/json"
	"go_api/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetBooks(t *testing.T) {
	books = []models.Book{
		{ID: 1, Title: "Book 1", ISBN: "111", Author: "Author 1", Year: 2001},
		{ID: 2, Title: "Book 2", ISBN: "222", Author: "Author 2", Year: 2002},
	}

	req, err := http.NewRequest("GET", "/books", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(BooksHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var responseBooks []models.Book
	if err := json.Unmarshal(rr.Body.Bytes(), &responseBooks); err != nil {
		t.Fatalf("could not parse response: %v", err)
	}

	if !compareBooks(responseBooks, books) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			responseBooks, books)
	}
}

func TestCreateBook(t *testing.T) {
	var jsonStr = []byte(`{"title":"New Book", "isbn":"12345", "author":"New Author", "year":2020}`)
	req, err := http.NewRequest("POST", "/books", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(BooksHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	var createdBook models.Book
	json.Unmarshal(rr.Body.Bytes(), &createdBook)
	if createdBook.Title != "New Book" {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), `{"id":1,"title":"New Book","isbn":"12345","author":"New Author","year":2020}`)
	}
}

func TestUpdateBook(t *testing.T) {
	books = []models.Book{
		{ID: 1, Title: "Book 1", ISBN: "111", Author: "Author 1", Year: 2001},
	}
	var jsonStr = []byte(`{"title":"Updated Book", "isbn":"111", "author":"Author 1", "year":2001}`)
	req, err := http.NewRequest("PUT", "/books/1", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(BookHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var updatedBook models.Book
	json.Unmarshal(rr.Body.Bytes(), &updatedBook)
	if updatedBook.Title != "Updated Book" {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), `{"id":1,"title":"Updated Book","isbn":"111","author":"Author 1","year":2001}`)
	}
}

func TestDeleteBook(t *testing.T) {
	books = []models.Book{
		{ID: 1, Title: "Book 1", ISBN: "111", Author: "Author 1", Year: 2001},
	}
	req, err := http.NewRequest("DELETE", "/books/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(BookHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := map[string]string{"message": "Book deleted"}
	var response map[string]string
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Fatalf("could not parse response: %v", err)
	}

	if response["message"] != expected["message"] {
		t.Errorf("handler returned unexpected body: got %v want %v",
			response, expected)
	}

	if len(books) != 0 {
		t.Errorf("book was not deleted: remaining books: %v", books)
	}
}

func compareBooks(a, b []models.Book) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
