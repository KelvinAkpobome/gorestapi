package model

import (
	"database/sql"
	"fmt"
	"log"
	
)

var connection *sql.DB   

// Createconnect creates the connection to the database 
func Createconnect() *sql.DB {
	db, err := sql.Open("mysql", "kelvin:akpos18%.?1@tcp(localhost:3306)/learn_api")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err) 
	}
	fmt.Println("Connected to database Successfully")
	connection = db
	return db
}