package controllers

import (
	"github.com/kataras/iris/sessions"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris"
	"html/template"
	"go-clock/model"
	"go-clock/utils"
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
	b.Handle("GET", "/enroll_user", "EnrollUser")
	b.Handle("GET", "/register_user_fingerprint/{user_id}", "RegisterUserFingerprint")
	b.Handle("GET", "/clock_in_or_out", "ClockInOrOut")
	b.Handle("GET", "/websocket_test", "WebSocketTest")
}

func (m *mainController) AfterActivation(a mvc.AfterActivation) {
	if a.Singleton() {
		panic("mainController should be stateless, a request-scoped, we have a 'Session' which depends on the context.")
	}
}

func (m *mainController) Index(ctx iris.Context) {
	ctx.View("main/index.html")
}

func (m *mainController) EnrollUser(ctx iris.Context) {
	usersInGroups, _ := model.Usuario{}.GetUsersFromLicencaInGroups()
	ctx.ViewData("usersInGroups", usersInGroups)
	ctx.View("main/enroll_user.html")
}

func (m *mainController) RegisterUserFingerprint(ctx iris.Context) {
	user, _ := model.Usuario{}.ReturnUserById(ctx.Params().Get("user_id"))
	ctx.ViewData("user", user)
	userPhoto := template.URL(user.Foto)
	ctx.ViewData("userPhoto", userPhoto)
	ctx.View("main/register_user_fingerprint.html")
}

func (m *mainController) ClockInOrOut(ctx iris.Context) {
	ctx.View("main/clock_in_or_out.html")
}

func (m *mainController) WebSocketTest(ctx iris.Context) {
	ctx.View("main/index.html")
}
