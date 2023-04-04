package constants

const (
	EMPTY_BODY     = "BODY IS EMPTY"
	ERROR          = "Error"
	BOOK_ID_EXIST  = "BookId already exist"
	BOOK_NOT_FOUND = "Book with ID %d not found. Please check the ID and try again."
	USER_NOT_FOUND = "User with ID %d not found. Please check the ID and try again."
)

// These constants define the error and success messages returned by the API handlers. They are encoded as JSON in the response.
const (
	NO_DATA_INSIDE_JSON  = "No data inside JSON"        // Error message when no data is present in the JSON request body.
	SEND_SOME_DATA       = "Please send some data"      // Error message when the JSON request body is empty.
	BOOK_ADDED           = "Book Added Successfully!"   // Success message when a book is added to the library.
	BOOK_DETAILS_UPDATED = "Book Details Updated!"      // Success message when the book details are updated in the library.
	BOOK_DELETED         = "Book Deleted Successfully!" // Success message when a book is deleted from the library.
	USER_ADDED           = "User Added Successful!"     //Success message when a user is added to the library.
	USER_DETAILS_UPDATED = "User Details Updated!"      //Success message when the user details are updated in the library.
	USER_DELETED         = "User Deleted Successfully!" //Success message when a user is going out from the library.
)
