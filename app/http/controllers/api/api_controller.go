package api

import (
	base "github.com/go-dyn/app/http/controllers"
)

type ApiController struct {
	//Dependent services
	base.Controller
}

func NewApiController() *ApiController {
	return &ApiController{
		//Inject services
	}
}
