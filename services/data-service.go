package services

import (
	"go-rest-api/models"
)

func PopulateData() {
	// mock data
	// Init books var as a slice nook struct
	models.Books = append(models.Books, models.Book{ID: "1", Isbn: "122323", Title: "My Cookbook", Author: &models.Author{Firstname: "John", Lastname: "Deer"}})
	models.Books = append(models.Books, models.Book{ID: "2", Isbn: "334175", Title: "Martian Stew", Author: &models.Author{Firstname: "Wendy", Lastname: "Goatnose"}})
	models.Books = append(models.Books, models.Book{ID: "3", Isbn: "580112", Title: "Flies", Author: &models.Author{Firstname: "Gus", Lastname: "Beard"}})
}
