package mysql

import (
	"errors"
	"fmt"

	"cognologix.com/main/constants"
	"cognologix.com/main/models"
)

/*
FindAllActiveUsers retrieves a list of all active users from the database.
If there are active users in the database, they are returned as a slice of models.User objects.
If no active users are found, an empty slice and a nil error are returned.
If an error occurs during the retrieval process, it is returned as an error.
*/
func (db *Mysql) FindAllActiveUsers() ([]models.User, error) {

	var users []models.User
	result := db.Client.Find(&users)
	return users, result.Error

}

/*
FindUserById retrieves a user from the database based on the given user ID.
If the user is found, it is returned along with a nil error.
If the user is not found, an error is returned with a custom message indicating that the user was not found.
*/
func (db *Mysql) FindUserById(userId int) (models.User, error) {

	var user models.User
	// SELECT * FROM users WHERE id = bookId;
	result := db.Client.First(&user, userId)
	if result.Error != nil {
		// if Id not found : no row is found
		customError := fmt.Sprintf(constants.USER_NOT_FOUND, userId)
		result.Error = errors.New(customError)
	}
	return user, result.Error

}

/*
AddUserInLibrary adds a new user to the library database.
It takes a user models as input and returns an error (if any).
*/
func (db *Mysql) AddUserInLibrary(user models.User) error {

	// INSERT INTO `books` (`id`,`bookName`) VALUES (book.bookId,book.bookName)
	result := db.Client.Create(&user)
	return result.Error

}

/*
UpdateUserDetails updates the assigned book details for a user with the given ID in the library database.
It takes the user's ID, new book ID, and new book name as input and returns an error (if any).
*/
func (db *Mysql) UpdateUserDetails(id int, bookId int, bookName string) error {

	var user models.User
	// update data
	result := db.Client.Model(&user).Where("id=?", id).Select("assigned_book_id", "assigned_book_name").Updates(models.User{AssignedBookId: bookId, AssignedBookName: bookName})
	return result.Error

}

/*
DeleteUser deletes a user with the given ID from the library database.
It takes the user's ID as input and returns error message.
*/
func (db *Mysql) DeleteUser(userId int) error {

	var user models.User
	result := db.Client.Delete(&user, userId)
	if result.RowsAffected == 0 {
		customError := fmt.Sprintf(constants.USER_NOT_FOUND, userId)
		result.Error = errors.New(customError)
	}
	return result.Error

}
