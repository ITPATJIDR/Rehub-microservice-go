package uploadController

import (
	fileModel "Rehub_Microservice/model"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context) {

	var uploadData fileModel.UploadFileSturct

	bytes := make([]byte, 10)
	_, err := rand.Read(bytes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	if err := c.ShouldBindJSON(&uploadData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	if _, err := os.Stat(`C:\Users\CoSI-Lenovite\Desktop\code\Rehub-microservice-go\uploads\` + uploadData.Foldername); os.IsNotExist(err) {

		err := os.MkdirAll(`C:\Users\CoSI-Lenovite\Desktop\code\Rehub-microservice-go\uploads\`+uploadData.Foldername, os.ModePerm)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}

		filename := `C:\Users\CoSI-Lenovite\Desktop\code\Rehub-microservice-go\uploads\` + uploadData.Foldername + "/" + uploadData.Foldername + "_" + hex.EncodeToString(bytes) + ".json"
		err = ioutil.WriteFile(filename, []byte(uploadData.Physicalrawreport), os.ModePerm)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message":  "File uploaded successfully",
			"Filepath": filename,
		})

	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

	} else {
		filename := `C:\Users\CoSI-Lenovite\Desktop\code\Rehub-microservice-go\uploads\` + uploadData.Foldername + "/" + uploadData.Foldername + "_" + hex.EncodeToString(bytes) + ".json"
		err := ioutil.WriteFile(filename, []byte(uploadData.Physicalrawreport), os.ModePerm)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message":  "File uploaded successfully",
			"Filepath": filename,
		})
	}
}

func DownloadFile(c *gin.Context) {
	var filePath fileModel.FileSturct

	if err := c.ShouldBindJSON(&filePath); err != nil {
		fmt.Print("HI")
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

func UploadVideo(c *gin.Context) {
	file, err := c.FormFile("videoBlob")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	filename := "videos/" + file.Filename + ".webm"
	err = c.SaveUploadedFile(file, filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": file,
	})
}
