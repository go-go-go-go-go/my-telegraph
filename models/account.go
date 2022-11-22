package models

type Account struct {
	ID          any
	ShortName   string `form:"short_name" json:"short_name" binding:"required"`
	AuthorName  string `form:"author_name" json:"author_name"`
	AuthorUrl   string `form:"author_url" json:"author_url"`
	AccessToken string `json:"access_token"`
	AuthUrl     string `json:"auth_url"`
	PageCount   string `json:"page_count"`
}
