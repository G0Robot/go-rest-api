package controllers

import (
	"encoding/json"
	"go-rest-api/models"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var Books []models.Book

// all route methods need to have a reponse and a request as arguments
// - why is r a pointer?
// Get Books
func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Books)

}

// Get a Book
func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// get parameters
	params := mux.Vars(r)

	// iterate through books for id
	for _, item := range Books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&models.Book{})
}

// Create Book
func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book models.Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(100000))
	Books = append(Books, book)
	json.NewEncoder(w).Encode(book)
}

// Update Book
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range Books {
		if item.ID == params["id"] {
			Books = append(Books[:index], Books[index+1:]...)

			var book models.Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			Books = append(Books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(Books)
}

// Delete Book
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range Books {
		if item.ID == params["id"] {
			Books = append(Books[:index], Books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(Books)
}
