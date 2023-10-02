package facades

import (
	"log"
)

func Rpem() contracts.Blog {
	instance, err := blog.App.Make(blog.Binding)
	if err != nil {
		log.Println(err)
		return nil
	}

	return instance.(contracts.Rpem)
}
