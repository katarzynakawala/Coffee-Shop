package main

import (
	"net/http"

	"katarzynakawala/github.com/coffee-shop/ui"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	dynamicMiddleware := alice.New(app.session.Enable, noSurf, app.authenticate)

	mux := pat.New()
	mux.Get("/", dynamicMiddleware.ThenFunc(app.home))
	mux.Get("/about", dynamicMiddleware.ThenFunc(app.about))
	mux.Get("/coffee/create", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.createCoffeeForm))
	mux.Post("/coffee/create", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.createCoffee))
	mux.Get("/coffee/:id", dynamicMiddleware.ThenFunc(app.displayCoffee))

	mux.Get("/user/signup", dynamicMiddleware.ThenFunc(app.signupUserForm))
	mux.Post("/user/signup", dynamicMiddleware.ThenFunc(app.signupUser))
	mux.Get("/user/login", dynamicMiddleware.ThenFunc(app.loginUserForm))
	mux.Post("/user/login", dynamicMiddleware.ThenFunc(app.loginUser))
	mux.Post("/user/logout", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.logoutUser))
	mux.Get("/user/profile", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.userProfile))
	mux.Get("/user/change-password", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.changePasswordForm))
	mux.Post("/user/change-password", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.changePassword))

	mux.Get("/ping", http.HandlerFunc(ping))

	fileServer := http.FileServer(http.FS(ui.Files))

	mux.Get("/static/", fileServer)

	return standardMiddleware.Then(mux)
}