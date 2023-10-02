package routes

import (
	"github.com/goravel/framework/contracts/route"
	"github.com/goravel/framework/facades"

	controllers "github.com/go-dyn/app/http/controllers/api"
	wsControllers "github.com/go-dyn/app/http/controllers/ws"
	middlewares "github.com/go-dyn/app/http/middleware"
)

func Api() {

	facades.Route().Prefix("oauth").Group(func(router route.Router) {
		authController := controllers.NewAuthController()
		router.Middleware(middlewares.Authenticated()).Get("me", authController.Show)
		router.Middleware(middlewares.Guess()).Post("login", authController.Login)
		router.Middleware(middlewares.Guess()).Post("register", authController.Register)
		router.Middleware(middlewares.Guess()).Post("password/forgot", authController.ForgotPassword)
		router.Middleware(middlewares.Guess()).Post("password/reset", authController.ResetPassword)
		router.Middleware(middlewares.Authenticated()).Post("refresh_token", authController.RefreshToken)
	})

	// Websocket
	websocketController := wsControllers.NewWebsocketController()
	facades.Route().Get("/ws", websocketController.Server)
}
