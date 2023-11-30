package routers

import (
	uploadController "Rehub_Microservice/controllers"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowCredentials = true
	config.AllowHeaders = append(config.AllowHeaders)
	config.AllowMethods = append(config.AllowMethods)

	r := gin.Default()
	r.Use(cors.New(config))
	r.Static("/videos", "./videos")

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Message": "I'm Alive...",
		})
	})

	upload := r.Group("/upload")
	{
		upload.POST("/uploadFile", uploadController.UploadFile)
		upload.POST("/uploadVideo", uploadController.UploadVideo)
	}

	download := r.Group("/download")
	{
		download.OPTIONS("/downloadFile", func(c *gin.Context) {
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
			c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			c.Status(http.StatusOK)
		})
		download.POST("/downloadFile", uploadController.DownloadFile)
	}

	return r
}
