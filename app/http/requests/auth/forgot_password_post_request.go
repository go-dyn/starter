package auth

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type ForgotPasswordPostRequest struct {
	Email string `form:"email" json:"email"`
}

func (r *ForgotPasswordPostRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *ForgotPasswordPostRequest) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"email": "required|email",
	}
}

func (r *ForgotPasswordPostRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *ForgotPasswordPostRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *ForgotPasswordPostRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
