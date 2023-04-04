package models

// Model of User
type User struct {
	Id               int    `json:"id" gorm:"int;not null;primary_key;auto_incremented"` //Is is the id of user
	UserName         string `json:"username" gorm:"varchar(35)"`                         // UserName is the name of the user.
	AssignedBookId   int    `json:"assignedbookid" gorm:"int;not null"`                  // AssignedBookId is the ID of the book assigned to the user.
	AssignedBookName string `json:"assignedbookname" grom:"varchar(25);not null"`        // AssignedBookName is the name of the book assigned to the user.
}

// BookId           int
