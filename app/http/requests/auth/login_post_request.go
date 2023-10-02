package auth

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type LoginPostRequest struct {
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
}

func (r *LoginPostRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *LoginPostRequest) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"email":    "required|email",
		"password": "required|max_len:255",
	}
}

func (r *LoginPostRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *LoginPostRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *LoginPostRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
