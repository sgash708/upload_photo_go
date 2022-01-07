package main

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestMainSuccess(t *testing.T) {
	// 変数の定義
	r := gin.Default()
	// ルートにアクセスしたら、JSONで「{messge:ping}」を返す
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ping",
		})
	})

	// localhost:8888でアクセス
	r.Run(":8888")
}
