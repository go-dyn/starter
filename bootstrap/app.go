package bootstrap

import (
	"github.com/goravel/framework/foundation"

	config "github.com/go-dyn/config"
)

func Boot() {
	app := foundation.NewApplication()

	//Bootstrap the application
	app.Boot()

	//Bootstrap the config.
	config.Boot()
}
