package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// /*
// This function retrieves all active users from the database and returns them as a JSON array.
// It also logs the execution time of the function.
// */
func (server *Server) GetAllUsers(c *gin.Context) {

	start := time.Now()
	users, err := server.DB.FindAllActiveUsers()
	if err != nil {
		logrus.Error(err.Error())
		c.IndentedJSON(http.StatusNotFound, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, users)
	duration := time.Since(start)
	logrus.Info("Execution time:", duration)

}

// /*
// This function retrieves a user from the database based on the ID provided in the request parameter and returns it in the response.
// It also logs the execution time of the function.
// */
func (server *Server) GetUserById(c *gin.Context) {

	start := time.Now()
	id := c.Param("id")
	idInInt, _ := strconv.Atoi(id)
	user, err := server.DB.FindUserById(idInInt)
	if err != nil {
		logrus.Error(err.Error())
		c.IndentedJSON(http.StatusNotFound, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, user)
	duration := time.Since(start)
	logrus.Info("Execution time:", duration)

}

// /*
// This function creates a new user and assigns a book to them.
// It first checks if the book is available and then assigns it to the user.
// If successful, it returns a message confirming the user addition.
// */
// func (server *Server) CreateUser(c *gin.Context) {
// 	start := time.Now()

// 	var user models.User
// 	if err := c.BindJSON(&user); err != nil {
// 		logrus.Error(constants.NO_DATA_INSIDE_JSON)
// 		c.IndentedJSON(http.StatusBadRequest, constants.NO_DATA_INSIDE_JSON)
// 		return
// 	}

// 	//Check for book avalibility
// 	book, err := server.db.FindBookById(user.AssignedBookId)
// 	if err != nil {
// 		logrus.Error(err.Error())
// 		c.IndentedJSON(http.StatusNotFound, err.Error())
// 		return
// 	}
// 	user.AssignedBookName = book.BookName

// 	//Assigning books
// 	err = server.db.AddUserInLibrary(user)
// 	if err != nil {
// 		logrus.Error(err.Error())
// 		c.IndentedJSON(http.StatusNotFound, err.Error())
// 		return
// 	}

// 	//Remove from stock
// 	err = server.db.DeleteBookFromLibrary(book.Id)
// 	if err != nil {
// 		logrus.Error(err.Error())
// 		c.IndentedJSON(http.StatusNotFound, err.Error())
// 		return
// 	}

// 	c.IndentedJSON(http.StatusOK, constants.USER_ADDED)

// 	duration := time.Since(start)
// 	logrus.Info("Execution time:", duration)
// }

// /*
// This function updates the details of a user and the book assigned to them.
// It first finds the user by the given id, adds their old book back to the library, checks if the new book is available, assigns it to the user, and removes it from the library.
// Finally, it responds with a success message and logs the execution time.
// */
// func (server *Server) UpdateUser(c *gin.Context) {

// 	start := time.Now()
// 	id := c.Param("id")
// 	idInInt, _ := strconv.Atoi(id)

// 	var user models.User

// 	// what if: body is empty
// 	if err := c.BindJSON(&user); err != nil {
// 		logrus.Error(constants.NO_DATA_INSIDE_JSON)
// 		c.IndentedJSON(http.StatusBadRequest, constants.NO_DATA_INSIDE_JSON)
// 		return
// 	}

// 	newBookId := user.AssignedBookId
// 	//finding user details to add old book
// 	user, err := server.db.FindUserById(idInInt)
// 	if err != nil {
// 		logrus.Error(err.Error())
// 		c.IndentedJSON(http.StatusNotFound, err.Error())
// 		return
// 	}

// 	var book models.Book
// 	book.Id = user.AssignedBookId
// 	book.BookName = user.AssignedBookName
// 	//Storing book in library
// 	err = server.db.AddBookInStock(book)
// 	if err != nil {
// 		logrus.Error(err.Error())
// 		c.IndentedJSON(http.StatusNotFound, err.Error())
// 		return
// 	}

// 	//Check for book avalibility of new BookId
// 	book, err = server.db.FindBookById(newBookId)
// 	if err != nil {
// 		logrus.Error(err.Error())
// 		c.IndentedJSON(http.StatusNotFound, err.Error())
// 		return
// 	}

// 	//Assigning new book to user
// 	err = server.db.UpdateUserDetails(user.Id, book.Id, book.BookName)
// 	if err != nil {
// 		logrus.Error(err.Error())
// 		c.IndentedJSON(http.StatusNotFound, err.Error())
// 		return
// 	}

// 	//Delete from stors
// 	err = server.db.DeleteBookFromLibrary(book.Id)
// 	if err != nil {
// 		logrus.Error(err.Error())
// 		c.IndentedJSON(http.StatusNotFound, err.Error())
// 		return
// 	}

// 	c.IndentedJSON(http.StatusOK, err.Error())
// 	duration := time.Since(start)
// 	logrus.Info("Execution time:", duration)
// }

// /*
// This function deletes a user by ID, adds the assigned book back to the library and returns a success message. It also logs the execution time.
// */
// func (server *Server) DeleteUser(c *gin.Context) {
// 	start := time.Now()
// 	id := c.Param("id")
// 	idInInt, _ := strconv.Atoi(id)

// 	//finding user details to add old book
// 	user, err := server.db.FindUserById(idInInt)
// 	if err != nil {
// 		logrus.Error(err.Error())
// 		c.IndentedJSON(http.StatusNotFound, err.Error())
// 		return
// 	}

// 	var book models.Book
// 	book.Id = user.AssignedBookId
// 	book.BookName = user.AssignedBookName
// 	//Storing book in library
// 	err = server.db.AddBookInStock(book)
// 	if err != nil {
// 		logrus.Error(err.Error())
// 		c.IndentedJSON(http.StatusNotFound, err.Error())
// 		return
// 	}

// 	//Removing user
// 	err = server.db.DeleteUser(idInInt)
// 	if err != nil {
// 		logrus.Error(err.Error())
// 		c.IndentedJSON(http.StatusNotFound, err.Error())
// 		return
// 	}

// 	c.IndentedJSON(http.StatusOK, constants.USER_DELETED)
// 	duration := time.Since(start)
// 	logrus.Info("Execution time:", duration)
// }
