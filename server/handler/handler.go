package handler

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

// Upload は画像をアップ後に「/images」ディレクトリに保存する。
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

// Delete は「/images」ディレクトリから削除する
func Delete(c *gin.Context) {
	uuid := c.Param("uuid")
	err := os.Remove(fmt.Sprintf("images/%s.png", uuid))
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("id: %s is deleted!", uuid)})
}

// File はファイル情報を持っている
type File struct {
	Path string `json:"path"`
	Size int64  `json:"size"`
}

func dirwalk(dir string) (files []File, err error) {
	// 全然意味がわかりません
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

// List は「/images」ディレクトリから一覧取得
func List(c *gin.Context) {
	files, err := dirwalk("./images")
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, files)
}
