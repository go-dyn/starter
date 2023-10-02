package auth

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type RefreshTokenPostRequest struct {
	Name string `form:"name" json:"name"`
}

func (r *RefreshTokenPostRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *RefreshTokenPostRequest) Rules(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *RefreshTokenPostRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *RefreshTokenPostRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *RefreshTokenPostRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
