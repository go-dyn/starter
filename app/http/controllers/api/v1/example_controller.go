package v1

import (
	requests "github.com/go-dyn/app/http/requests/example"
	"github.com/go-dyn/app/models"
	services "github.com/go-dyn/app/services"
	"github.com/goravel/framework/contracts/http"
)

type ExampleController struct {
	//Dependent services
	BaseController
	service services.IExample
}

func NewExampleController() *ExampleController {
	return &ExampleController{
		service: services.NewExampleImpl(),
	}
}

type ResultControllerIndex struct {
	list  []models.Example
	total int64
	err   error
}

func (c *ExampleController) List(ctx http.Context) http.Response {
	page := ctx.Request().InputInt("page", 1)
	limit := ctx.Request().InputInt("limit", 10)
	ch := make(chan ResultControllerIndex)

	go func() {
		defer close(ch)
		list, total, err := c.service.Paginate(page, limit)
		ch <- ResultControllerIndex{
			list, total, err,
		}
	}()

	result := <-ch

	if result.err != nil {
		return c.Error(ctx, map[string]any{
			"error": result.err,
		})
	}

	return c.Success(ctx, map[string]any{
		"items":        result.list,
		"current_page": page,
		"pages":        c.TotalPage(result.total, int64(limit)),
		"limit":        limit,
		"total":        result.total,
	})
}

func (c *ExampleController) Show(ctx http.Context) http.Response {
	id := ctx.Request().RouteInt64("id")
	item, err := c.service.Get(uint64(id))

	if err != nil {
		return c.Error(ctx, map[string]any{
			"error": err,
		})
	}

	return c.Success(ctx, map[string]any{
		"item": item,
	})
}

func (c *ExampleController) Store(ctx http.Context) http.Response {
	var post requests.ExamplePostRequest
	errors, err := c.ValidateRequest(ctx, &post)
	if err != nil {
		return c.ValidateFail(ctx, map[string]any{
			"msg":    "validation fail",
			"errors": errors,
		})
	}

	item, err := c.service.Create(models.Example{
		Title:   post.Title,
		Content: post.Content,
	})

	if err != nil {
		return c.Error(ctx, map[string]any{
			"error": err,
		})
	}

	return c.Success(ctx, map[string]any{
		"item": item,
	})
}

func (c *ExampleController) Update(ctx http.Context) http.Response {
	var post requests.ExamplePostRequest
	Id := ctx.Request().RouteInt64("id")

	if Id <= 0 {
		return c.NotFound(ctx, map[string]any{})
	}

	errors, err := c.ValidateRequest(ctx, &post)
	if err != nil {
		return c.ValidateFail(ctx, map[string]any{
			"msg":    "validation fail",
			"errors": errors,
		})
	}

	item, err := c.service.Get(uint64(Id))

	if err != nil {
		return c.NotFound(ctx, map[string]any{})
	}

	err = c.service.MergeTo(post, &item)

	if err != nil {
		return c.Error(ctx, map[string]any{
			"error": err,
		})
	}

	item, err = c.service.Update(item)

	if err != nil {
		return c.Error(ctx, map[string]any{
			"error": err,
		})
	}

	return c.Success(ctx, map[string]any{
		"item": item,
	})
}

func (c *ExampleController) Destroy(ctx http.Context) http.Response {
	Id := ctx.Request().RouteInt64("id")

	if Id <= 0 {
		return c.NotFound(ctx, map[string]any{})
	}

	item, err := c.service.Get(uint64(Id))

	if err != nil {
		return c.NotFound(ctx, map[string]any{})
	}

	item, _, err = c.service.Destroy(item)

	if err != nil {
		return c.Error(ctx, map[string]any{
			"error": err,
		})
	}

	return c.Success(ctx, map[string]any{
		"item": item,
	})
}
