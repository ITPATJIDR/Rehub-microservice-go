package routers

import (
	"github.com/gin-gonic/gin"
  "Rehub_Microservice/controllers"
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
  }


  return r
}
