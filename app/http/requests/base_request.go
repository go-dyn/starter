package auth

import (
	"encoding/json"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type BaseRequest struct {
}

func (r *BaseRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *BaseRequest) Rules(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *BaseRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *BaseRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *BaseRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}

func (r *BaseRequest) ToMap() (map[string]any, error) {
	var out map[string]interface{}
	inrec, _ := json.Marshal(r)
	err := json.Unmarshal(inrec, &out)
	return out, err
}
