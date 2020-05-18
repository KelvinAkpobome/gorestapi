package controller

import (
	"book-api/model"
	"book-api/views"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)
//controller for update
func updateDel(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
		fmt.Println("delete request recieved")
		id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/api/v1/books/"))
		if err != nil {
				log.Fatal(err)
			} 
		fmt.Println(id)
		if err := model.DeleteBook(id); err != nil {
			w.Write([]byte("An error occurred "))
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(struct {
			Status string `json:"status"`
		}{"Item Deleted"})
		
	}
}