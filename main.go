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

	r.GET("ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})

	userRepo := controllers.New()
	r.POST("books", userRepo.CreateUser)
	r.GET("books", userRepo.GetUsers)
	r.GET("books/:id", userRepo.GetUser)
	r.PUT("books/:id", userRepo.UpdateUser)
	r.DELETE("books/:id", userRepo.DeleteUser)

	return r
}



