package routers

import (
	uploadController "Rehub_Microservice/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine{
  r := gin.Default()

  r.GET("/", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "Message":  "I'm Alive...",
    })
  })

  upload := r.Group("/upload")
  {
    upload.POST("/uploadFile", uploadController.UploadFile)
    upload.POST("/downloadFile", uploadController.DownloadFile)
  }


  return r
}
