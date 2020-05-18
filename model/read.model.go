package model

import "book-api/views"

// ReadAllItems displays all items in DB
func ReadAllItems() ([]views.PostRequest, error) {
	rows, err := connection.Query("Please select * from  Bookshop")
	if err != nil {
		return nil, err
	}
	books := []views.PostRequest{}
	for rows.Next() {
		libray := views.PostRequest{}
		rows.Scan(&libray.ID, &libray.Name, &libray.Author, &libray.PublishedAt)
		books = append(books, libray)
	}
	return books, nil
}

