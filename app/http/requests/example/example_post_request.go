package auth

import (
	base "github.com/go-dyn/app/http/requests"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type ExamplePostRequest struct {
	base.BaseRequest
	Title   string `form:"title" json:"title"`
	Content string `form:"content" json:"content"`
}

func (r *ExamplePostRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *ExamplePostRequest) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"title":   "required|max_len:255",
		"content": "required|max_len:65535",
	}
}

func (r *ExamplePostRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *ExamplePostRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *ExamplePostRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
