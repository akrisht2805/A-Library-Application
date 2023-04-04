package main

import (
	"net/http"
	"strconv"
	"time"

	"cognologix.com/main/constants"
	"cognologix.com/main/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// // GetAllBooks retrieves all books from the repository layer FindAllBooks function and return all books details.
func (server *Server) GetAllBooks(c *gin.Context) {

	start := time.Now()
	books, err := server.DB.FindAllBooks()
	if err != nil {
		logrus.Error(err.Error())
		c.IndentedJSON(http.StatusNotFound, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, books)
	duration := time.Since(start)
	logrus.Info("Execution time:", duration)
}

/*
GetBookById is a handler function that retrieves a book by ID from the repository,
and returns it as a JSON response. It also logs the execution time of the function.
*/
func (server *Server) GetBookById(c *gin.Context) {

	start := time.Now()
	// extracting the employee id from http request
	id := c.Param("id")

	idInInt, _ := strconv.Atoi(id)
	book, err := server.DB.FindBookById(idInInt)
	if err != nil {
		logrus.Error(err.Error())
		c.IndentedJSON(http.StatusNotFound, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, book)
	duration := time.Since(start)
	logrus.Info("Execution time:", duration)
}

/*
This function creates a new book in the library and returns a success status if successful.
If there is no data inside JSON or if there is an error in adding the book, it returns an error status.
It also logs the execution time.
*/
func (server *Server) CreateBook(c *gin.Context) {
	start := time.Now()
	var book models.Book

	//What if BookName is Empty
	if err := c.BindJSON(&book); err != nil {
		logrus.Error(constants.NO_DATA_INSIDE_JSON)
		c.IndentedJSON(http.StatusBadRequest, constants.NO_DATA_INSIDE_JSON)
		return
	}

	err := server.DB.AddBookInStock(book)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}
	logrus.Info(constants.BOOK_ADDED)
	c.IndentedJSON(http.StatusCreated, constants.BOOK_ADDED)
	duration := time.Since(start)
	logrus.Info("Execution time:", duration)
}

// // /*
// // UpdateBook updates the name of a book with the given ID and returns success message if the
// // update is successful. If the book with the given ID is not found, it returns error message.
// // */
func (server *Server) UpdateBook(c *gin.Context) {
	start := time.Now()

	id := c.Param("id")
	idInInt, _ := strconv.Atoi(id)

	var book models.Book

	if err := c.BindJSON(&book); err != nil {
		logrus.Error(err)
		c.IndentedJSON(http.StatusBadRequest, constants.NO_DATA_INSIDE_JSON)
		return
	}

	err := server.DB.UpdateBookDetails(idInInt, book.BookName)
	if err != nil {
		logrus.Error(err.Error())
		c.IndentedJSON(http.StatusNotFound, err.Error())
		return
	}

	logrus.Info(constants.BOOK_DETAILS_UPDATED)
	c.IndentedJSON(http.StatusOK, constants.BOOK_DETAILS_UPDATED)
	duration := time.Since(start)
	logrus.Info("Execution time:", duration)
}

// // /*
// // This function deletes a book from the library based on the ID passed in the request parameters.
// // It also logs the execution time and returns a success or error message as a JSON response.
// // */
// func (server *Server) DeleteBook(c *gin.Context) {
// 	start := time.Now()

// 	id := c.Param("id")
// 	idInInt, _ := strconv.Atoi(id)

// 	err := server.db.DeleteBookFromLibrary(idInInt)

// 	if err != nil {
// 		logrus.Error(err.Error())
// 		c.IndentedJSON(http.StatusNotFound, err.Error())
// 		return
// 	}

// 	logrus.Info(constants.BOOK_DELETED)
// 	c.IndentedJSON(http.StatusOK, constants.BOOK_DELETED)
// 	duration := time.Since(start)
// 	logrus.Info("Execution time:", duration)
// }
