package models

type DB interface {

	//For books
	FindAllBooks() ([]Book, error)
	FindBookById(id int) (Book, error)
	AddBookInStock(book Book) error
	UpdateBookDetails(id int, newBookName string) error
	// DeleteBookFromLibrary(id int) error

	// //For User
	FindAllActiveUsers() ([]User, error)
	FindUserById(id int) (User, error)
	// AddUserInLibrary(user User) error
	// UpdateUserDetails(id int, bookId int, bookName string) error
	// DeleteUser(id int) error
}
