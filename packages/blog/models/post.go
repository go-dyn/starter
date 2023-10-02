package models

import (
	attachment "github.com/go-dyn/attachment/models"
)

type Post struct {
	BaseModel
	Content        string
	AuthorableID   uint
	AuthorableType string
	Attachments    []*attachment.Attachment `gorm:"polymorphic:Attachmentsable"`
	Comments       []*Comment               `gorm:"polymorphic:Commentable"`
}
