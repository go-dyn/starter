package attachment

import (
	"github.com/goravel/framework/contracts/config"
)

type Attachment struct {
	config config.Config
}

func NewAttachment(config config.Config) *Attachment {
	return &Attachment{config: config}
}
