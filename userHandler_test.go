package main

import (
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

func TestGetAllUsers(t *testing.T) {
	var server Server

	mockdb := mockdb.New()
	server.DB = mockdb

	users := []models.User{{
		Id:               1,
		UserName:         "Akrisht",
		AssignedBookId:   2,
		AssignedBookName: "Hello C",
	}}
	mockdb.On("FindAllActiveUsers").Return(users, nil)

	gin.SetMode(gin.TestMode)
	r := gin.Default()

	r.GET("/users", server.GetAllUsers)
	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	json.Unmarshal(w.Body.Bytes(), &users)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, users)

}

func TestGetUserById(t *testing.T) {
	var server Server

	mockdb := mockdb.New()
	server.DB = mockdb

	user := models.User{
		Id:               1,
		UserName:         "Akrisht",
		AssignedBookId:   2,
		AssignedBookName: "Hello C",
	}
	mockdb.On("FindUserById", user.Id).Return(user, nil)

	gin.SetMode(gin.TestMode)
	r := gin.Default()

	r.GET("/user/:id", server.GetUserById)
	req, _ := http.NewRequest("GET", "/user/"+fmt.Sprint(user.Id), nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// var books []models.Book
	json.Unmarshal(w.Body.Bytes(), &user)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, user)
}
