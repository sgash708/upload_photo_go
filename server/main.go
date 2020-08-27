package main

import (
	"main/handler"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// JSでAPIを叩く時のために処理している
	r.Use(cors.New(cors.Config{
		// localhost:8080を許容
		AllowOrigins: []string{"http://localhost:8080"},
		/*
			以下のメソッドが使える
			r.GET
			r.POST
			r.DELETE
			r.OPTIONS
		*/
		AllowMethods: []string{"GET", "POST", "DELETE", "OPTIONS"},
		// Headerはなんでも良い
		AllowHeaders: []string{"*"},
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
