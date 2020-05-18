package main

import (
	//"fmt"
	"book-api/controller"
	"book-api/model"
	"fmt"
	"log"
	"net/http"
	


func main() {
	mux := controller.Register()
	db := model.Createconnect()
	defer db.Close()
	fmt.Println("Server is up and running...")
	log.Fatal(http.ListenAndServe(":3000",mux))
}