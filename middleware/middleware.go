package middleware

import (
	"bookshelf-app/initializers"
	"bookshelf-app/models"
	"bookshelf-app/utils"
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {

	if len(c.Request.Header["Key"]) == 0 || len(c.Request.Header["Sign"]) == 0 {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    "INTERNAL_SERVER_ERR",
			"message": "access denied",
		})
		return
	}

	var user models.User
	key := c.Request.Header["Key"][0]
	sign := c.Request.Header["Sign"][0]

	result := initializers.DB.Where("key=?", key).First(&user)
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code":    "UNAUTHORIZED",
			"message": "access denied",
		})
		return
	}

	reqBodyByte, err := ioutil.ReadAll(c.Request.Body)
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(reqBodyByte))
	if err != nil {
		reqBodyByte = []byte("")
	}
	hash := utils.CreateHash(c.Request.Method + "http://" + c.Request.Host + c.Request.RequestURI + string(reqBodyByte) + user.Secret)
	if sign != hash {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code":    "UNAUTHORIZED",
			"message": "access denied",
		})
		return
	}

	c.Set("user", user)
	c.Next()
}
