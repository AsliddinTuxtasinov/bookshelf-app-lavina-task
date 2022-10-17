package handlers

import (
	"bookshelf-app/middleware"
	"log"

	"github.com/gin-gonic/gin"
)

func Run() {
	router := gin.New()

	router.POST("/books", CreateBook)
	router.GET("/books", middleware.Auth, GetAllBooks)
	router.PATCH("/books/:book_id", UpdateBook)
	router.DELETE("/books/:book_id", DeleteBook)
	router.POST("/signup", CreateUser)
	router.GET("/myself", middleware.Auth, GetUser)

	if err := router.Run(":8080"); err != nil {
		log.Fatalln(err.Error())
	}
}
