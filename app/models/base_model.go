package models

import (
	"github.com/goravel/framework/database/orm"
)

type BaseModel struct {
	orm.Model
	orm.SoftDeletes
}
