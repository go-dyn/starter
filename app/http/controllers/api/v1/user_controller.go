package v1

import (
	"github.com/go-dyn/app/services"
	"github.com/goravel/framework/contracts/http"
)

type UserController struct {
	//Dependent services
	BaseController
	service services.IUser
}

func NewUserController() *UserController {
	return &UserController{
		service: services.NewUserImpl(),
	}
}

func (c *UserController) List(ctx http.Context) http.Response {
	page := ctx.Request().InputInt("page", 1)
	limit := ctx.Request().InputInt("limit", 10)

	items, total, err := c.service.Paginate(page, limit)

	if err != nil {
		return c.Error(ctx, map[string]any{})
	}

	return c.Success(ctx, map[string]any{
		"total": total,
		"page":  page,
		"limit": limit,
		"items": items,
	})
}

func (c *UserController) Show(ctx http.Context) http.Response {
	Id := ctx.Request().RouteInt64("id")

	if Id <= 0 {
		return c.NotFound(ctx, map[string]any{})
	}

	item, err := c.service.Get(uint64(Id))

	if err != nil {
		return c.Error(ctx, map[string]any{})
	}

	return c.Success(ctx, map[string]any{
		"item": item,
	})
}
