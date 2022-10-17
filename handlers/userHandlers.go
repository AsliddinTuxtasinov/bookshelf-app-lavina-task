package handlers

import (
	"bookshelf-app/initializers"
	"bookshelf-app/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var reqBody struct {
		Name   string `json:"name"`
		Key    string `json:"key"`
		Secret string `json:"secret"`
	}

	// Get date of request body
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	fmt.Println(reqBody)
	user := models.User{Name: reqBody.Name, Key: reqBody.Key, Secret: reqBody.Secret}
	if tx := initializers.DB.Create(&user); tx.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"isOk":    false,
			"message": tx.Error,
		})
		return
	}
	// Response
	c.JSON(http.StatusCreated, gin.H{
		"data":    user,
		"isOk":    true,
		"message": "ok",
	})
}

func GetUser(c *gin.Context) {
	user := c.Keys["user"]
	// Response
	c.JSON(http.StatusCreated, gin.H{
		"data":    user,
		"isOk":    true,
		"message": "ok",
	})
}
