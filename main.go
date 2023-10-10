package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sauravranapng/modules/controllers"
	"net/http"
)

func main() {
	r := setupRouter()
	_ = r.Run(":8080")
}

func setupRouter() *gin.Engine {
	r := gin.Default()
         //This line initializes a Gin router using the default settings
	//The r variable now holds an instance of the Gin router with default configurations, including middleware.
	r.GET("ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})
//defines an anonymous function (closure) that takes a single parameter c, 
//which is a context object of type *gin.Context.
//The context object contains information about the incoming HTTP request and provides methods for handling and responding to the request.
	userRepo := controllers.New()
	//It calls the New function of a controllers package (or type) to create this instance.
	r.POST("books", userRepo.CreateUser)
	r.GET("books", userRepo.GetUsers)
	r.GET("books/:id", userRepo.GetUser)
	r.PUT("books/:id", userRepo.UpdateUser)
	r.DELETE("books/:id", userRepo.DeleteUser)

	return r
}



