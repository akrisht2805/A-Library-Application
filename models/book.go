package models

// Model of Book
type Book struct {
	Id       int    `json:"id" gorm:"int;not null;primary_key"` //Id of book
	BookName string `json:"bookname" gorm:"not null"`           //BookName is the name of book
}

// User     []User `gorm:"foreignKey:BookId"`
