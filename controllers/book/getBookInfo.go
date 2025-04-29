package book

import (
	GetBookInfoLogic "basic-crud/logics/book"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetBookInfoHandler(w http.ResponseWriter, r *http.Request){
	id := mux.Vars(r)["id"]	;
	if id == "" {
		http.Error(w,"NO ID GIVEN",http.StatusBadRequest)
		return
	}
	book, err := GetBookInfoLogic.GetBookInfo(id)
	if err!=nil{
		http.Error(w, "INTERNAL_SERVER_ERROR",http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type","application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}