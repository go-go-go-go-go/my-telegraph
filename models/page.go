package models

type Page struct {
	ID          any
	Path        string         `form:"path" json:"path"`
	Url         string         `form:"url" json:"url"`
	Title       string         `form:"title" json:"title"`
	Description string         `form:"description" json:"description"`
	AuthorName  string         `form:"author_name" json:"author_name"`
	AuthorUrl   string         `form:"author_url" json:"author_url"`
	ImageUrl    string         `form:"image_url" json:"image_url"`
	Content     map[string]any `form:"content" json:"content"`
	Views       int            `form:"views" json:"views"`
	CanEdit     bool           `form:"can_edit" json:"can_edit"`
}