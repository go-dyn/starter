package models

type Comment struct {
	BaseModel
	Content        string
	AuthorableID   uint
	AuthorableType string
}
