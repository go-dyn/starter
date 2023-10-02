package attachment

import (
	"github.com/goravel/framework/contracts/foundation"
)

const Binding = "attachment"

var App foundation.Application

type ServiceProvider struct {
}

func (receiver *ServiceProvider) Register(app foundation.Application) {
	App = app

	app.Bind(Binding, func(app foundation.Application) (any, error) {
		return NewAttachment(app.MakeConfig()), nil
	})
}

func (receiver *ServiceProvider) Boot(app foundation.Application) {
	app.Publishes("github.com/go-dyn/attachment", map[string]string{
		"config/attachment.go": app.ConfigPath("attachment.go"),
	}, "attachment-config")
	app.Publishes("github.com/go-dyn/attachment", map[string]string{
		"migrations": app.DatabasePath("migrations"),
	}, "attachment-migrations")
}
