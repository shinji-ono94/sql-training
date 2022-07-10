package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"main.go/sql"
	// "github.com/~/sql" //SQLについて記述したパッケージのパス
)

func main() {
	// Echoのインスタンス
	e := echo.New()

	// ミドルウェア類
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	//e.Use(interceptor.BasicAuth())

	// ルーティング
	e.GET("/sql/record/:id", sql.GetPost()) //プレースホルダでidをもらってくる
	e.GET("/sql/table", sql.GetPosts())

	// サーバー起動
	e.Start(":1323")
}
