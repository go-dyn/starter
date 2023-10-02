package api

import (
	requests "github.com/go-dyn/app/http/requests/auth"
	models "github.com/go-dyn/app/models"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

type AuthController struct {
	//Dependent services
	ApiController
}

func NewAuthController() *AuthController {
	return &AuthController{
		//Inject services
	}
}

func (c *AuthController) Show(ctx http.Context) http.Response {
	var user models.User
	err := facades.Auth().User(ctx, &user)
	if err != nil {
		return c.Unauthorized(ctx, map[string]any{
			"msg": "Unauthorized",
		})
	}

	return c.Success(ctx, map[string]any{
		"data": user,
	})
}

func (c *AuthController) Login(ctx http.Context) http.Response {
	var post requests.LoginPostRequest
	errors, err := c.ValidateRequest(ctx, &post)
	if err != nil {
		return c.ValidateFail(ctx, map[string]any{
			"msg":    "validation fail",
			"errors": errors,
		})
	}

	var user models.User
	err = facades.Orm().Query().Where("identifier", post.Email).First(&user)

	if err != nil {
		return c.NotFound(ctx, map[string]any{
			"msg": "NotFound",
		})
	}

	compare := facades.Hash().Check(post.Password, user.Password)

	if !compare {
		return c.Error(ctx, map[string]any{
			"msg": "Invalid email or password",
		})
	}

	token, err := facades.Auth().Login(ctx, &user)

	if err != nil {
		return c.Error(ctx, map[string]any{
			"msg": "Invalid email or password",
		})
	}

	return c.Success(ctx, map[string]any{
		"token":         token,
		"refresh_token": token,
		"user":          user,
	})
}

func (c *AuthController) Register(ctx http.Context) http.Response {
	var post requests.ExamplePostRequest
	errors, err := c.ValidateRequest(ctx, &post)
	if err != nil {
		return c.ValidateFail(ctx, map[string]any{
			"msg":    "validation fail",
			"errors": errors,
		})
	}

	var user models.User
	passwd, err := facades.Hash().Make(post.Password)

	if err != nil {
		return c.Error(ctx, map[string]any{
			"msg":    "Create fail",
			"params": post,
		})
	}

	err = facades.Orm().Query().Where("identifier", post.Email).FirstOrCreate(&user, models.User{Name: post.Name, Identifier: post.Email, Email: post.Email, UserName: post.Email, Password: passwd})

	if err != nil {
		return c.Error(ctx, map[string]any{
			"msg":    "Create fail",
			"params": post,
		})
	}

	token, err := facades.Auth().Login(ctx, &user)

	if err != nil {
		return c.Error(ctx, map[string]any{
			"msg": "Invalid email or password",
		})
	}

	return c.Success(ctx, map[string]any{
		"token":         token,
		"refresh_token": token,
		"user":          user,
	})
}

func (c *AuthController) ForgotPassword(ctx http.Context) http.Response {
	var post requests.ForgotPasswordPostRequest
	errors, err := c.ValidateRequest(ctx, &post)
	if err != nil {
		return c.ValidateFail(ctx, map[string]any{
			"msg":    "validation fail",
			"errors": errors,
		})
	}

	return c.Success(ctx, map[string]any{
		"Hello": "ForgotPassword",
	})
}

func (c *AuthController) ResetPassword(ctx http.Context) http.Response {
	var post requests.ResetPasswordPostRequest
	errors, err := c.ValidateRequest(ctx, &post)
	if err != nil {
		return c.ValidateFail(ctx, map[string]any{
			"msg":    "validation fail",
			"errors": errors,
		})
	}
	return c.Success(ctx, map[string]any{
		"Hello": "ResetPassword",
	})
}

func (c *AuthController) RefreshToken(ctx http.Context) http.Response {
	token, err := facades.Auth().Refresh(ctx)
	if err != nil {
		return c.Forbiden(ctx, map[string]any{
			"msg": "Forbiden",
		})
	}

	return c.Success(ctx, map[string]any{
		"token":         token,
		"refresh_token": token,
	})

}
