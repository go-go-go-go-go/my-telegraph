package models

type Account struct {
	Id          int    `json:"-"`
	ShortName   string `form:"short_name" json:"short_name,omitempty" binding:"required"`
	AuthorName  string `form:"author_name" json:"author_name,omitempty"`
	AuthorUrl   string `form:"author_url" json:"author_url,omitempty"`
	AccessToken string `form:"access_token" json:"access_token,omitempty"`
	AuthUrl     string `json:"auth_url,omitempty"`
	PageCount   int    `json:"page_count,omitempty"`
}
