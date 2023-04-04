package mysql

import (
	"errors"
	"fmt"

	"cognologix.com/main/constants"
	"cognologix.com/main/models"
)

/*
FindAllBooks retrieves all books from the database and returns them as a slice of Book structs.
If an error occurs while querying the database, the function returns an empty slice and the error.
*/
func (db *Mysql) FindAllBooks() ([]models.Book, error) {

	var books []models.Book
	result := db.Client.Find(&books)

	return books, result.Error
}

/*
FindBookById retrieves a book from the database based on the given book ID.
If the book is found, it is returned along with a nil error. If the book is not found,
an error is returned with a Book Not Found message indicating that the book was not found.
*/
func (db *Mysql) FindBookById(bookId int) (models.Book, error) {

	var book models.Book

	// SELECT * FROM books WHERE id = bookId;
	row := db.Client.First(&book, bookId)

	if row.Error != nil {
		// if Id not found : no row is found
		customError := fmt.Sprintf(constants.BOOK_NOT_FOUND, bookId)
		row.Error = errors.New(customError)
	}

	return book, row.Error
}

/*
AddBookInStock adds a new book to the database stock.
If the book is added successfully, a nil error is returned, else return a error.
*/
func (db *Mysql) AddBookInStock(book models.Book) error {

	var b models.Book
	db.Client.First(&b, book.Id)

	if b.Id != 0 {
		return errors.New(constants.BOOK_ID_EXIST)
	}

	//INSERT INTO `books` (`id`,`bookName`) VALUES (book.bookId,book.bookName)
	result := db.Client.Create(&book)

	return result.Error
}

/*
UpdateBookDetails updates the name of a book in the database based on its ID.
If the update is successful, a nil error is returned.Else it is returned as database error.
*/
func (db *Mysql) UpdateBookDetails(id int, name string) error {

	var book models.Book
	// update data
	result := db.Client.Model(&book).Where("id=?", id).Update("book_name", name)

	return result.Error
}

/*
DeleteBookFromLibrary deletes a book from the database based on its ID.
If the book is deleted successfully, a nil error is returned.
If the book with the given ID is not found in the database, Book with given id is not found error returned.
*/
func (db *Mysql) DeleteBookFromLibrary(bookId int) error {
	var book models.Book

	// delete data
	result := db.Client.Delete(&book, bookId)

	if result.RowsAffected == 0 {
		// if Id not found : err : no row is found
		customError := fmt.Sprintf(constants.BOOK_NOT_FOUND, bookId)
		result.Error = errors.New(customError)
	}

	return result.Error
}
