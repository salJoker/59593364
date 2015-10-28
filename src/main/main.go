package main

import(
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	"html/template"
	"io"
	"net/http"
	"fmt"
)

type Template struct{
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}) error {
	return t.templates.ExecuteTemplate(w,name,data)
}

func Hello(c *echo.Context) error {
	return c.Render(http.StatusOK,"hello","world")
}

func main(){
	echo_serve := echo.New()

	mws := []echo.Middleware{mw.Logger(),mw.Recover(),mw.Gzip()}
	//注册日志、故障恢复、响应Gzip压缩中间件
	echo_serve.Use(mws...)

//	echo_serve.Static("/","templates")
//	echo_serve.Static("/js","sources/scripts")
//	echo_serve.Static("/css","sources/css")
//
//	echo_serve.Index("templates/index.html")

	tmpl,err := template.ParseGlob("templates/*.html");
	if err != nil{
		fmt.Print(err.Error())
	}

	t := &Template{
		templates:template.Must(tmpl,err),
	}
	echo_serve.SetRenderer(t)
	echo_serve.Get("/hello",Hello)

	echo_serve.Run(":9060")
}