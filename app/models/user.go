package models

import (
	attachment "github.com/go-dyn/attachment/models"
	factories "github.com/go-dyn/database/factories"
	"github.com/goravel/framework/contracts/database/factory"
)

type User struct {
	BaseModel
	Identifier string
	Name       string `json:"Name"`
	UserName   string `json:"UserName"`
	Password   string `json:"Password"`
	Email      string `json:"Email"`
	Active     uint8
	Status     uint8

	Avatar attachment.Attachment
}

func (u *User) IsActive() bool {
	return u.Active == 1
}

func (u *User) GetActive() uint8 {
	return u.Active
}

func (u *User) GetStatus() uint8 {
	return u.Status
}

func (u *User) Factory() factory.Factory {
	return &factories.UserFactory{}
}
