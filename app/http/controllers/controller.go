package controllers

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
	"github.com/goravel/framework/facades"
)

type Controller struct {
	//Dependent services
}

func NewController() *Controller {
	return &Controller{
		//Inject services
	}
}

func (c *Controller) Response(ctx http.Context, code int, data map[string]any) http.Response {
	if code != http.StatusOK && code != http.StatusCreated {
		c.Log("api error", http.Json{
			"code": code,
			"data": data,
		})
	}
	data["code"] = code
	data["key"] = http.StatusText(code)
	return ctx.Response().Json(code, http.Json{"data": data})
}

func (c *Controller) Success(ctx http.Context, data map[string]any) http.Response {
	return c.Response(ctx, http.StatusOK, data)
}

func (c *Controller) Error(ctx http.Context, data map[string]any) http.Response {
	return c.Response(ctx, http.StatusBadRequest, data)
}

func (c *Controller) Forbiden(ctx http.Context, data map[string]any) http.Response {
	return c.Response(ctx, http.StatusForbidden, data)
}

func (c *Controller) Unauthorized(ctx http.Context, data map[string]any) http.Response {
	return c.Response(ctx, http.StatusUnauthorized, data)
}

func (c *Controller) NotFound(ctx http.Context, data map[string]any) http.Response {
	return c.Response(ctx, http.StatusNotFound, data)
}

func (c *Controller) ValidateFail(ctx http.Context, data map[string]any) http.Response {
	return c.Response(ctx, http.StatusUnprocessableEntity, data)
}

func (c *Controller) Validate(ctx http.Context, rules map[string]string) (validator validation.Validator, err error) {

	validator, err = ctx.Request().Validate(rules)
	return
}

func (c *Controller) ValidateRequest(ctx http.Context, request http.FormRequest) (errors validation.Errors, err error) {

	errors, err = ctx.Request().ValidateRequest(request)
	return
}

func (c *Controller) Log(args ...any) {
	go facades.Log().Info(args)
}

func (c *Controller) TotalPage(total int64, limit int64) int64 {
	pages := total / limit

	if total%limit > 0 {
		pages++
	}

	return pages
}
