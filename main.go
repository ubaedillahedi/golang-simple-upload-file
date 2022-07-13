package main

import (
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func main() {

	router := gin.Default()

	router.GET("/", func (c *gin.Context)  {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"message": "Welcome, REST API Upload File Golang!",
		})
	})

	router.POST("/upload", func (c *gin.Context)  {
		file, _ := c.FormFile("file")
		ext := filepath.Ext(file.Filename)
		target := "uploads/images/"
		newUUID := uuid.New()
		newFileName := newUUID.String() + ext
		c.SaveUploadedFile(file, target + newFileName)

		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"message": "Success, Uploaded File!",
		})

	})

	router.GET("list-image", func (c *gin.Context)  {
		var responseFiles []string

		files, _ := ioutil.ReadDir("./uploads/images/")
		for _, file := range files{
			responseFiles = append(responseFiles, file.Name())
		}



		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"message": "Success!",
			"baseUrl": "localhost:8888/images/",
			"data": responseFiles,
		})

	})



	router.Static("/images", "./uploads/images/")

	router.Run(":8888")

}