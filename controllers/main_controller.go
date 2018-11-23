package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"iris-base-app/utils"
)

func MainCtrl(app *mvc.Application) {
	// You can use normal middlewares at MVC apps of course.
	app.Router.Use(func(ctx iris.Context) {
		ctx.Application().Logger().Infof("Path: %s", ctx.Path())
		ctx.Next()
	})

	// Register dependencies which will be binding to the controller(s),
	// can be either a function which accepts an iris.Context and returns a single value (dynamic binding)
	// or a static struct value (service).
	app.Register(
		sessions.New(sessions.Config{}).Start,
		&utils.PrefixedLogger{Prefix: "DEV"},
	)

	// GET: http://localhost:8080/
	app.Handle(new(mainController))
}

type mainController struct {
	Logger  utils.LoggerService
	Session *sessions.Session
}

func (m *mainController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/", "Index")
}

func (m *mainController) AfterActivation(a mvc.AfterActivation) {
	if a.Singleton() {
		panic("mainController should be stateless, a request-scoped, we have a 'Session' which depends on the context.")
	}
}

func (m *mainController) Index(ctx iris.Context) {
	ctx.View("main/index.html")
}
