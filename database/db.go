package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const DB_USERNAME = "Saurav"
const DB_PASSWORD = "Saurav11"
const DB_NAME = "test"
const DB_HOST = "127.0.0.1"
const DB_PORT = "3306"

var Db *gorm.DB
//declaring an instance of the GORM database
func InitDb() *gorm.DB {
	Db = connectDB()
	return Db
}

func connectDB() *gorm.DB {
	var err error
	dsn := DB_USERNAME + ":" + DB_PASSWORD + "@tcp" + "(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?" + "parseTime=true&loc=Local"
        // DSN  Data Source Name (DSN) string
	fmt.Println("dsn : ", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
         //mysql.Open(dsn) is become possible because of mysql driver 
        //&gorm.Config{} is struct that is used to specify various configuration options 
        // for the behavior of GORM when interacting with the database.we are passing with default setting 
	if err != nil {
		//fmt.Println("Error connecting to database : error=%v", err)
		return nil
	}

	return db
}
