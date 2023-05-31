package main

import (
	
  "net/http"


  "github.com/gin-gonic/gin"
)


type book struct{
	Id string 
	Title string
	Author string
	Price  int 
}

var books = []book{
	 {Id:"1",Title: "Rich Dad Poor Dad",Author: " Robert T. Kiyosaki" ,Price: 202,},
	 {Id:"2",Title: "Ikigai",Author: "  Francesc Miralles " ,Price: 368,},
	 {Id:"3",Title: "Atomic Habits",Author: " James Clear " ,Price: 423,},
	 {Id:"4",Title: "The Psychology of Money ",Author: "  Morgan Housel " ,Price: 210,},
	
}

func addBooks(c *gin.Context) {
    var newbook book

    //it will check if taken input from client is matching with book object format or not ,if not then value of 
	//err will not be nil hence it will return from there itself
    if err := c.BindJSON(&newbook); err != nil {
        return 
    }

    // Add the new book to the list.
    books= append(books, newbook)
	//sending to client 
    c.IndentedJSON(http.StatusCreated, newbook)
}
func getBooks(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, books)
}
func getBookById(c*gin.Context){
	//since context have all the things user has sent so we are retrieving id from the path using Param("id")
id:=c.Param("id")

for i,a:=range books{
if a.Id==id {
c.IndentedJSON(http.StatusOK,books[i])
return
}
}
c.IndentedJSON(http.StatusNotFound,gin.H{"message":"book not by of this id"})

}
func updateBookById(c*gin.Context){
	id:=c.Param("id")
	var newbook book
	for i,a:=range books {
		if a.Id==id {
			if err := c.BindJSON(&newbook); err != nil {
				c.IndentedJSON(http.StatusBadRequest,gin.H{"message":"Book is not updated"})
				return 
			}
         books[i]=newbook
		 c.IndentedJSON(http.StatusOK,gin.H{"message":"Book is updated"})
		}
	}

}
func deleteBookById(c*gin.Context){
	id:=c.Param("id")
	for i,a:=range books {
		if a.Id==id {
			books = append(books[:i], books[i+1:]...)
        c.IndentedJSON(http.StatusAccepted,gin.H{"message":"book deleted"})
			return
		}}

c.IndentedJSON(http.StatusNotFound,gin.H{"message":"book not found"})
}
func main() {
  r := gin.Default()
  r.GET("/books", getBooks)
  r.POST("/books",addBooks)
  r.GET("/getbyid/:id",getBookById)//  :(colon) indicate id is a path parameter
  r.PATCH("/update/:id",updateBookById)
  r.DELETE("/delete/:id",deleteBookById)
  r.Run("localhost:8080") //it is my desktop
}
