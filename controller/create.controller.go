package controller

import (
	"book-api/model"
	"encoding/json"
	"fmt"
	"net/http"
)


func create(w http.ResponseWriter, r *http.Request,) {
	if r.Method == http.MethodPost {
		fmt.Println("Request received")
		if err := model.CreateBook(value.ID, value.Name, value.Author); err != nil {
			w.Write([]byte("Sorry, an error occurred "))
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(value)

	} else if r.Method == http.MethodGet {
			books, err := model.ReadAll()
			if err != nil {
				w.Write([]byte(err.Error()))
				return
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(books)
	} 
}

