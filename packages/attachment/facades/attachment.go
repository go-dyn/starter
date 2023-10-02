package facades

import (
	"log"
)

func Attachment() contracts.Attachment {
	instance, err := attachment.App.Make(attachment.Binding)
	if err != nil {
		log.Println(err)
		return nil
	}

	return instance.(contracts.Attachment)
}
