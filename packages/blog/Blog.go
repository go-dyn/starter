package blog

import (
	"github.com/goravel/framework/contracts/config"
)

type Blog struct {
	config config.Config
}

func NewBlog(config config.Config) *Blog {
	return &Blog{config: config}
}
