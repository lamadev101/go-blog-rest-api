package controllers

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func HandleFileUpload(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	files := form.File["files"]
	for _, file := range files {
		err := saveFile(file)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Files uploaded successfully"})
}

func saveFile(fileHeader *multipart.FileHeader) error {
	fmt.Println(fileHeader.Filename)
	// ext := strings.Split(fileHeader.Filename, ".")
	// if ext[1] != "png" {
	// 	return errors.New("pass a valid png")
	// }

	src, err := fileHeader.Open()
	if err != nil {
		return err
	}

	defer src.Close()

	// Create a new file in the desired destination folder
	// dstPath := filepath.Join("../files", ext[0]+strconv.Itoa(key)+"."+ext[1])
	dstPath := filepath.Join("./files", fileHeader.Filename)

	dst, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer dst.Close()

	return nil
}
