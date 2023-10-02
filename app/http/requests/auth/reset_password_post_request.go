package auth

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type ResetPasswordPostRequest struct {
	Email string `form:"email" json:"email"`
	Token string `form:"token" json:"token"`
}

func (r *ResetPasswordPostRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *ResetPasswordPostRequest) Rules(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *ResetPasswordPostRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{
		"email": "required|email",
		"token": "required|max_len:255",
	}
}

func (r *ResetPasswordPostRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *ResetPasswordPostRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
