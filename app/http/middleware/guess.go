package middleware

import (
	"net/http"

	contractshttp "github.com/goravel/framework/contracts/http"
)

func Guess() contractshttp.Middleware {
	return func(ctx contractshttp.Context) {
		token := ctx.Request().Header("Authorization", "")
		if len(token) > 0 {
			ctx.Request().AbortWithStatus(http.StatusForbidden)
			return
		}

		ctx.Request().Next()
	}
}
