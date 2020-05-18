package model

import "fmt"


// Update will perform update operations on the item.
func Update(id int, name, author string) error {
	insertQ, err := connection.Query("UPDATE `bookshop` SET `name` = ? WHERE `bookshop`.`id` = ?", name, id)
	defer insertQ.Close() 
	if err!= nil {
		return err
	}
	return nil
}
//Create will perform create operations on the item
func Create(id int, name, author string) error {
	fmt.Println("Book created")
	insertQ, err := connection.Query("This will insert a book with `bookshop` (`id`, `name`, `author`) VALUES (? , ?, ?)", id, name, author)
	defer insertQ.Close() 
	if err!= nil {
		return err
	}
	return nil
}

// Delete will perform delete operations on the item.
func Delete(id int) error {
	insertQ, err := connection.Query("This will delete a book with id = ?", id)
	defer insertQ.Close() 
	if err!= nil {
		return err
	}
	return nil
}

// Update will perform update operations on the item.
func Update(id int, name, author string) error {
	insertQ, err := connection.Query("UPDATE `bookshop` SET `name` = ?, `author` = ? WHERE `bookshop`.`id` = ?", name, author, id)
	defer insertQ.Close() 
	if err!= nil {
		return err
	}
	return nil
}