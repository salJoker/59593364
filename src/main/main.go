package main

import(
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
)

func main(){
	echo_serve := echo.New()

	mws := []echo.Middleware{mw.Logger(),mw.Recover(),mw.Gzip()}
	//注册日志、故障恢复、响应Gzip压缩中间件
	echo_serve.Use(mws...)

	echo_serve.Static("/","templates")
	echo_serve.Static("/js","sources/scripts")
	echo_serve.Static("/css","sources/css")

	echo_serve.Index("templates/index.html")

	echo_serve.Run(":9060")
}