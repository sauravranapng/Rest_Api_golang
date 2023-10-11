package controllers

import (
	"errors"
	//for HTTP status codes
	"net/http"
	"github.com/sauravranapng/modules/database"
	"github.com/sauravranapng/modules/models"


	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type books_db struct {

	Db *gorm.DB
//Db is become the database connection object from the GORM library.
}

func New() *books_db {

	db := database.InitDb()
// This line initializes a database connection using the InitDb function from the database package
//The db variable now holds a reference to the database connection.
	db.AutoMigrate(&models.Book{})
//This line performs an automatic migration of the database schema.
// It ensures that the database schema matches the Go struct representation of a book, 
//which is typically defined in the models package as models.Book.
	return &books_db{Db: db}
//creates a new books_db controller instance and sets its Db field to the previously initialized database connection (db).
}

//create user
func (Books_db *books_db) CreateUser(c *gin.Context) {
//(Books_db *books_db): This part is the receiver, which specifies the type on which the function is defined.
//func RECEIVER_NAME.FUNCTION_NAME(PARAMETERS) RETURN_TYPE {}
//(c *gin.Context) :This part is the function's parameter list. It specifies the input parameters that the function expects.

	var book models.Book

	c.BindJSON(&book)
//c.BindJSON(&book) is used to bind or parse the JSON data from the HTTP request body
// and populate the book variable with it.
//This is a common pattern in web applications to extract data sent by clients in JSON format 
//and map it to a Go struct for further processing
	err := models.CreateUser(Books_db.Db, &book)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
//If an error is detected, it responds to the client with an HTTP status code 500 (Internal Server Error)
// and includes a JSON response containing an error message
		return
	}
	c.JSON(http.StatusOK, book)
}

//get users
func (Books_db *books_db) GetUsers(c *gin.Context) {
	var book []models.Book
	err := models.GetUsers(Books_db.Db, &book)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, book)
}

//get user by id
func (Books_db *books_db) GetUser(c *gin.Context) {
	id:= c.Param("id")
	var book models.Book
	err := models.GetUser(Books_db.Db, &book, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, book)
}

// update user
func (Books_db *books_db) UpdateUser(c *gin.Context) {
	id:= c.Param("id")
	var book models.Book
	err := models.GetUser(Books_db.Db, &book, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
		return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.BindJSON(&book)
	err = models.UpdateUser(Books_db.Db, &book)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, book)
}

// delete user
func (Books_db *books_db) DeleteUser(c *gin.Context) {
	var book models.Book
	id:= c.Param("id")
	err := models.DeleteUser(Books_db.Db, &book, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
