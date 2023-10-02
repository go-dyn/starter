package auth

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type ExamplePostRequest struct {
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
	Name     string `form:"name" json:"name"`
}

func (r *ExamplePostRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *ExamplePostRequest) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"email":    "required|email",
		"password": "required|max_len:255",
		"name":     "required|max_len:255",
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
