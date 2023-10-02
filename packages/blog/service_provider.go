package blog

import (
	"github.com/goravel/framework/contracts/foundation"
)

const Binding = "blog"

var App foundation.Application

type ServiceProvider struct {
}

func (receiver *ServiceProvider) Register(app foundation.Application) {
	App = app

	app.Bind(Binding, func(app foundation.Application) (any, error) {
		return NewBlog(app.MakeConfig()), nil
	})
}

func (receiver *ServiceProvider) Boot(app foundation.Application) {
	app.Publishes("github.com/go-dyn/blog", map[string]string{
		"config/blog.go": app.ConfigPath("blog.go"),
	}, "blog-config")
	app.Publishes("github.com/go-dyn/blog", map[string]string{
		"migrations": app.DatabasePath("migrations"),
	}, "blog-migrations")
}
