package routes

import (
	"basic-crud/controllers/book"
	"github.com/gorilla/mux"
)

func BookRoutes(r *mux.Router){
	r.HandleFunc("/book",book.CreateBookHandler).Methods("POST")
}