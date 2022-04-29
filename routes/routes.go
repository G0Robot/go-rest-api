package routes

import (
	"go-rest-api/controllers"
	"github.com/gorilla/mux"
)

var RegisterRoutes = func(r *mux.Router) {
	// Route Handlers / Endpoints
	// HandleFunc sets up the routers and which methods will handle them
	r.HandleFunc("/api/books", controllers.GetBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", controllers.GetBook).Methods("GET")
	r.HandleFunc("/api/books", controllers.CreateBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", controllers.UpdateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", controllers.DeleteBook).Methods("DELETE")
}
