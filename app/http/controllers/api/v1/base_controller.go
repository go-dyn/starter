package v1

import (
	api "github.com/go-dyn/app/http/controllers/api"
)

type BaseController struct {
	//Dependent services
	api.ApiController
}

func NewBaseController() *BaseController {
	return &BaseController{
		//Inject services
	}
}
