package main

import (
	"github.com/sgash708/upload_photo_go/handler"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// JSでAPIを叩く時のために処理している
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:        false,
		AllowOrigins:           []string{"http://localhost:8080"},
		AllowMethods:           []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowHeaders:           []string{"*"},
		AllowCredentials:       false,
		ExposeHeaders:          []string{},
		MaxAge:                 0,
		AllowWildcard:          false,
		AllowBrowserExtensions: false,
		AllowWebSockets:        false,
		AllowFiles:             false,
	}))

	// 静的ファイルの配信追加
	r.Use(static.Serve("/", static.LocalFile("./images", true)))

	// 選択
	r.GET("/images", handler.List)
	// 追加
	r.POST("/images", handler.Upload)
	// 削除
	r.DELETE("/images/:uuid", handler.Delete)

	// localhost:8888でアクセス
	r.Run(":8888")
}
