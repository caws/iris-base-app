package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"os"
	"go-clock/controllers"
)

func main() {
	app := iris.Default()
	app.Logger().SetLevel("debug")

	app.StaticWeb("/css", "./assets/css")
	app.StaticWeb("/js", "./assets/js")
	app.StaticWeb("/fonts", "./assets/fonts")
	app.StaticWeb("/images", "./assets/images")

	app.RegisterView(iris.HTML("./views", ".html"))

	mvc.Configure(app.Party("/"), controllers.MainCtrl)

	port := ":" + os.Getenv("PORT")

	if port == ":" {
		port = ":8080"
	}

	app.Run(iris.Addr(port))
}
