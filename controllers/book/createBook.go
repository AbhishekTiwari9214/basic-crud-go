package book

import (
	bookLogic "basic-crud/logics/book"
	"basic-crud/models"
	"encoding/json"
	"net/http"
)

func CreateBookHandler(w http.ResponseWriter, r *http.Request) {

	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "INVALID_INPUT", http.StatusBadRequest)
		return
	}
	createdBook, err := bookLogic.CreateBook(book)
	if err != nil {
		http.Error(w, "Error creating book: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type","application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdBook)
}
