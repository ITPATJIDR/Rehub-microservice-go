package routers

import (
	uploadController "Rehub_Microservice/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowCredentials = true
	config.AllowHeaders = append(config.AllowHeaders)

	r := gin.Default()
	r.Use(cors.New(config))

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Message": "I'm Alive...",
		})
	})

	upload := r.Group("/upload")
	{
		upload.POST("/uploadFile", uploadController.UploadFile)
		upload.POST("/downloadFile", uploadController.DownloadFile)
		upload.POST("/uploadVideo", uploadController.UploadVideo)
	}

	return r
}
