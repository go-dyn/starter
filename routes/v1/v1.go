package routes

import (
	"github.com/goravel/framework/contracts/route"
	"github.com/goravel/framework/facades"

	controllers "github.com/go-dyn/app/http/controllers/api/v1"
)

func V1() {
	facades.Route().Prefix("v1").Group(func(router route.Router) {

		//users
		userController := controllers.NewUserController()
		router.Get("users", userController.List)
		router.Prefix("users").Group(func(router route.Router) {
			router.Get("{id}", userController.Show)
		})

		// examples
		exampleController := controllers.NewExampleController()
		router.Get("examples", exampleController.List)
		router.Post("examples", exampleController.Store)
		router.Prefix("examples").Group(func(router route.Router) {
			router.Get("{id}", exampleController.Show)
			router.Put("{id}", exampleController.Update)
			router.Delete("{id}", exampleController.Destroy)
		})
	})
}
