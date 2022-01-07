package handler

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

// File ファイル情報_構造体
type File struct {
	Path string `json:"path"`
	Size int64  `json:"size"`
}

// Upload 画像をアップロードし、imagesディレクトリに保存
func Upload(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["file"]

	// UUID を取得
	uuid := c.PostForm("uuid")

	for _, file := range files {
		// ファイル名にUUIDを付与
		err := c.SaveUploadedFile(file, "images/"+uuid+".png")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
	}

	c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
}

// Delete imagesディレクトリから画像を削除する
func Delete(c *gin.Context) {
	uuid := c.Param("uuid")
	err := os.Remove(fmt.Sprintf("images/%s.png", uuid))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("id: %s is deleted!", uuid)})
}

// List imagesディレクトリから一覧を取得する
func List(c *gin.Context) {
	files, err := dirwalk("./images")
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, files)
}

// dirwalk ディレクトリ内検索
func dirwalk(dir string) (files []File, err error) {
	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		path = strings.Replace(path, "images/", "http://localhost:8888/", 1)
		size := info.Size()
		f := File{
			Path: path,
			Size: size,
		}
		files = append(files, f)
		return nil
	})
	if err != nil {
		return
	}
	files = files[1:]
	return
}
