package main

import (
	"go-rest-api/controllers"
	"go-rest-api/models"
	"go-rest-api/routes"
	"go-rest-api/services"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {

	controllers.Books = models.Books

	services.PopulateData()

	// init router
	r := mux.NewRouter()
	// call to set up route handlers
	routes.RegisterRoutes(r)

	// Add a listener on port 8000, which needs to use our router
	log.Fatal(http.ListenAndServe(":8000", r))
}
