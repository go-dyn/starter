package models

import (
	factories "github.com/go-dyn/database/factories"
	"github.com/goravel/framework/contracts/database/factory"
)

type Example struct {
	BaseModel
	Title   string `json:"Title"`
	Content string `json:"Content"`
	Active  uint8  `json:"Active"`
	Status  uint8  `json:"Status"`
}

func (u *Example) Factory() factory.Factory {
	return &factories.ExampleFactory{}
}
