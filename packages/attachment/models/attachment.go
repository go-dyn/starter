package models

type Attachment struct {
	BaseModel
	Name               string
	Path               string
	Size               uint
	ext                string
	AttachmentableID   uint
	AttachmentableType string
	AuthorableID       uint
	AuthorableType     string
	Type               uint8
}
