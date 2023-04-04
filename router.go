package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) RouterHandlers() {

	router := gin.Default()

	//Serve Home Route
	router.GET("/", ServeHome)

	//Routing for Book

	// localhost:3000/books
	router.GET("/books", server.GetAllBooks)

	//localhost:3000/book/{id}
	router.GET("/book/:id", server.GetBookById)

	//localhost:3000/book
	router.POST("/book", server.CreateBook)

	//localhost:3000/book/{id}
	router.PUT("/book/:id", server.UpdateBook)

	// //localhost:3000/book/{id}
	// router.DELETE("/book/:id", server.DeleteBook)

	// //Routing For User

	//localhost:3000/users
	router.GET("/users", server.GetAllUsers)

	//localhost:3000/user/{id}
	router.GET("/user/:id", server.GetUserById)

	// //localhost:3000/user
	// router.POST("/user", server.CreateUser)

	// //localhost:3000/user/{id}
	// router.PUT("/user/:id", server.UpdateUser)

	// //localhost:3000/user/{id}
	// router.DELETE("/user/:id", server.DeleteUser)

	router.Run("localhost:3000")

}

// serve home route
func ServeHome(c *gin.Context) {
	c.JSON(http.StatusOK, "Welcome to Library Application using API")
}
