package mockdb

import (
	"errors"
	"fmt"

	"cognologix.com/main/models"
	"github.com/stretchr/testify/mock"
)

func New() *MockDB {
	return &MockDB{}
}

type MockDB struct {
	mock.Mock
}

func (mock *MockDB) FindAllBooks() ([]models.Book, error) {
	rets := mock.Called()
	return rets.Get(0).([]models.Book), rets.Error(1)
}

func (mock *MockDB) FindBookById(bookId int) (models.Book, error) {
	book := models.Book{}
	if bookId == 1 {
		book = models.Book{
			Id:       1,
			BookName: "Programming Fundamentals",
		}
		return book, nil
	}
	return book, errors.New("Id not exist")
}

func (mock *MockDB) AddBookInStock(book models.Book) error {
	rets := mock.Called(book)
	fmt.Println(book)
	return rets.Error(0)
}

func (mock *MockDB) UpdateBookDetails(id int, newBookName string) error {
	rets := mock.Called(id, newBookName)
	return rets.Error(0)
}

func (mock *MockDB) DeleteBook(book models.Book) error {
	return errors.New("Not Implemented")
}

// User
func (mock *MockDB) FindAllActiveUsers() ([]models.User, error) {
	rets := mock.Called()
	return rets.Get(0).([]models.User), rets.Error(1)
}

func (mock *MockDB) FindUserById(userId int) (models.User, error) {
	user := models.User{}
	if userId == 1 {
		user = models.User{
			Id:               1,
			UserName:         "Akrisht",
			AssignedBookId:   2,
			AssignedBookName: "Hello",
		}
		return user, nil
	}
	return user, errors.New("Id not exist")
}



func (mock *MockDB) AddUserInLibrary(user models.User) error {
	return errors.New("Not Implemented")
}
func (mock *MockDB) UpdateUserDetails(user models.User) error {
	return errors.New("Not Implemented")
}
func (mock *MockDB) DeleteUser(user models.User) error {
	return errors.New("Not Implemented")
}
