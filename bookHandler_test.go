package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"cognologix.com/main/models"
	"cognologix.com/main/models/mockdb"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetAllBooks(t *testing.T) {
	var server Server

	mockdb := mockdb.New()
	server.DB = mockdb

	books := []models.Book{{
		Id:       1,
		BookName: "Programming Fundamentals",
	}}
	mockdb.On("FindAllBooks").Return(books, nil)

	gin.SetMode(gin.TestMode)
	r := gin.Default()

	r.GET("/books", server.GetAllBooks)
	req, err := http.NewRequest("GET", "/books", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	json.Unmarshal(w.Body.Bytes(), &books)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, books)

	// if status := w.Code; status != http.StatusOK {
	// 	t.Errorf("handler returned wrong status code: got %v want %v",
	// 		status, http.StatusOK)
	// }

}

func TestGetBookById(t *testing.T) {
	var server Server

	mockdb := mockdb.New()
	server.DB = mockdb

	book := models.Book{
		Id:       1,
		BookName: "Programming",
	}

	mockdb.On("FindBookById", book.Id).Return(book, nil)

	gin.SetMode(gin.TestMode)
	r := gin.Default()

	r.GET("/book/:id", server.GetBookById)
	req, _ := http.NewRequest("GET", "/book/"+fmt.Sprint(book.Id), nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// var books []models.Book
	json.Unmarshal(w.Body.Bytes(), &book)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, book)
}

func TestCreateBook(t *testing.T) {
	var server Server

	mockdb := mockdb.New()
	server.DB = mockdb

	book := models.Book{
		Id:       3,
		BookName: "Anything",
	}
	mockdb.On("AddBookInStock", book).Return(nil)

	gin.SetMode(gin.TestMode)
	r := gin.Default()

	r.POST("/book", server.CreateBook)

	jsonValue, err := json.Marshal(book)
	if err != nil {
		fmt.Print(err)
	}
	// fmt.Println(string(jsonValue))
	req := httptest.NewRequest("POST", "/book", bytes.NewBuffer(jsonValue))
	// req := httptest.NewRequest("POST", "/book", strings.NewReader(string(jsonValue)))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}
func TestUpdateBook(t *testing.T) {
	var server Server

	mockdb := mockdb.New()
	server.DB = mockdb

	book := models.Book{
		Id:       1,
		BookName: "Programming 2",
	}
	mockdb.On("UpdateBookDetails", book.Id, book.BookName).Return(nil)

	gin.SetMode(gin.TestMode)
	r := gin.Default()

	r.PUT("/book/:id", server.UpdateBook)

	jsonValue, err := json.Marshal(book)
	if err != nil {
		fmt.Print(err)
	}

	req := httptest.NewRequest("PUT", "/book/"+fmt.Sprint(book.Id), bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

// var tests = []struct {
// 	name               string
// 	method             string
// 	mockOn             string
// 	args               models.Book
// 	handler            gin.HandlerFunc
// 	expectedStatusCode int
// 	expectedError      error
// }{
// 	{"Valid Book ID", "GET", "FindBookById", models.Book{
// 		Id:       1,
// 		BookName: "Let Us C",
// 	}, server.GetBookById, 404, errors.New("error")},
// 	{"InValid Book ID", "GET", "FindBookById", models.Book{
// 		Id:       3,
// 		BookName: "Hello",
// 	}, server.GetBookById, 404, errors.New("error")},
// }
