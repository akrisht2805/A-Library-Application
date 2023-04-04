package main

import (
	"fmt"
	"net/http"
	"os"

	"cognologix.com/main/models"
	"cognologix.com/main/models/mysql"
	"github.com/sirupsen/logrus"
)

// Log File Name
const (
	logFile = "logFile.txt"
)

type Server struct {
	DB     models.DB
	Routes http.Handler
}

func newServer() *Server {
	s := Server{}
	return &s
}

// Main Function
func main() {

	server := newServer()
	db, err := mysql.ConnectDB()
	if err != nil {
		logrus.Error("Connection Failed")
	}

	server.DB = db

	// Open the log file for writing, creating it if it doesn't exist and appending to the end if it does.
	file, err := os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)

	// If there was an error opening the file, print an error message and panic.
	if err != nil {
		fmt.Println("Failed to create logfile " + logFile)
		panic(err)
	}

	// Defer the closing of the log file until the function returns.
	defer file.Close()

	// Set the output of the logrus logger to the opened file.
	logrus.SetOutput(file)

	// Set the log level to Debug.
	logrus.SetLevel(logrus.DebugLevel)

	// initializes the HTTP routing for a web application.
	server.RouterHandlers()
}
