package handlers

import (
	"bookshelf-app/middleware"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Run() {
	router := gin.New()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data":    "asliddin",
		})
	})
	router.POST("/books", middleware.Auth, CreateBook)
	router.GET("/books", middleware.Auth, GetAllBooks)
	router.PATCH("/books/:book_id", middleware.Auth, UpdateBook)
	router.DELETE("/books/:book_id", middleware.Auth, DeleteBook)
	router.POST("/signup", CreateUser)
	router.GET("/myself", middleware.Auth, GetUser)
	router.GET("/cleanup", CleanUp)


	if err := router.Run(":8000"); err != nil {
		log.Fatalln(err.Error())
	}
}
