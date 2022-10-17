package handlers

import (
	"bookshelf-app/initializers"
	"bookshelf-app/models"
	"bookshelf-app/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateBook(c *gin.Context) {
	var reqBody struct {
		ISBN string `json:"isbn"`
	}

	// Get date of request body
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	book := utils.GetBookInfoByIsbn(reqBody.ISBN)
	// Create book
	bookDb := models.Book{
		BookModel: models.BookModel{
			ISBN:      book.ISBN,
			Title:     book.Title,
			Author:    book.Author,
			Published: book.Published,
			Pages:     book.Pages,
		},
	}
	if tx := initializers.DB.Create(&bookDb); tx.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"isOk":    false,
			"message": tx.Error,
		})
		return
	}

	// Response
	c.JSON(http.StatusCreated, gin.H{
		"data": gin.H{
			"book": bookDb,
		},
		"isOk":    true,
		"message": "ok",
	})
}

func GetAllBooks(c *gin.Context) {
	var books []models.Book
	if tx := initializers.DB.Find(&books); tx.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"isOk":    false,
			"message": tx.Error,
		})
		return
	}

	// Response
	c.JSON(http.StatusOK, gin.H{
		"data":    books,
		"isOk":    true,
		"message": "ok",
	})
}

func UpdateBook(c *gin.Context) {
	var book_id = c.Param("book_id")
	var book models.Book
	var reqBody struct {
		Status uint `json:"status"`
	}

	// Get date of request body
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"isOk":    false,
			"message": err.Error,
		})
	}

	if tx := initializers.DB.Where("id = ?", book_id).First(&book); tx.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"isOk":    false,
			"message": tx.Error,
		})
		return
	}

	var statuses = map[uint]string{0: "new", 1: "read", 2: "finished"}
	if _, ok := statuses[reqBody.Status]; ok {
		fmt.Println("ok= true")
		book.Status = statuses[reqBody.Status]
		initializers.DB.Save(&book)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"isOk":    false,
			"message": "status is not defaine",
		})
		return
	}

	// Response
	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"book": book,
		},
		"isOk":    true,
		"message": "ok",
	})

}

func DeleteBook(c *gin.Context) {
	var book_id = c.Param("book_id")

	initializers.DB.Where("id = ?", book_id).Delete(&models.Book{})

	// Response
	c.JSON(http.StatusOK, gin.H{
		"data":    "Successfully deleted",
		"isOk":    true,
		"message": "ok",
	})
}
