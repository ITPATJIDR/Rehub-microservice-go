package uploadController

import (
	"Rehub_Microservice/model"
  "os"
	"net/http"
  "io"

	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context) {
  file, err := c.FormFile("file")  
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{
      "error": err.Error(),
    })
    return
  }

  filename := "uploads/" + file.Filename
  err = c.SaveUploadedFile(file,filename)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{
      "error": err.Error(),
    })
    return
  }

  c.JSON(http.StatusOK, gin.H{
    "message": "File uploaded successfully",
    "filename": filename,
  })
}

func DownloadFile(c *gin.Context) {
  var filePath fileModel.FileSturct

  if err:= c.ShouldBindJSON(&filePath); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
  }

  file, err := os.Open(filePath.Filepath)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": "Error opening file"})
    return
  }
  defer file.Close()

  c.Header("Content-Description", "File Transfer")
  c.Header("Content-Transfer-Encoding", "binary")
  c.Header("Content-Disposition", "attachment; filename="+filePath.Filepath)
  c.Header("Content-Type", "application/octet-stream")
  c.Header("Expires", "0")
  c.Header("Cache-Control", "must-revalidate")
  c.Header("Pragma", "public")

  _, err = io.Copy(c.Writer, file)
      if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error copying file content to response"})
        return
      }

  c.Writer.Flush()
}
