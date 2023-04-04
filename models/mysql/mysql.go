package mysql

import (
	"encoding/json"
	"os"

	"cognologix.com/main/models"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// const filePath = "/home/akshrits/GitHub/Akrisht_Yadav/Demonstrate/Golng Iteration 5 demo/libraryApp/configuration/config.json"
const filePath = "././configuration/config.json"

type Configuration struct {
	DataSourceName string
}

type Mysql struct {
	Client *gorm.DB
}

/*
dbInitilizer reads the configuration file at the specified filePath,and returns the DataSourceName field of the resulting Configuration object.
The DataSourceName field is a string containing the configuration details needed to connect to a MySQL database.
*/
func dbInitilizer() string {

	file, err := os.Open(filePath)
	if err != nil {
		logrus.Fatal(err)
	}

	defer file.Close()

	configuration := new(Configuration)
	json.NewDecoder(file).Decode(configuration)

	return configuration.DataSourceName
}

/*
ConnectDB establishes a connection to a MySQL database using GORM and returns the database connection.
It uses the configuration provided by dbInitilizer() to set up the database connection.
*/
func ConnectDB() (*Mysql, error) {

	dataSourceName := dbInitilizer()

	client, err := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})

	mysql := &Mysql{
		Client: client,
	}

	// Set up the database tables for the Book and User models using db.AutoMigrate()
	client.AutoMigrate(&models.Book{})
	client.AutoMigrate(&models.User{})

	return mysql, err
}
