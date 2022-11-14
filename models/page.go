package models

type Page struct {
	Id            int
	AccountId     int    `form:"account_id" json:"account_id"`
	AccessToken   string `form:"access_token" json:"access_token"`
	ReturnContent bool   `form:"return_content" json:"return_content"`
	Path          string `form:"path" json:"path"`
	Url           string `form:"url" json:"url"`
	Title         string `form:"title" json:"title"`
	Description   string `form:"description" json:"description"`
	AuthorName    string `form:"author_name" json:"author_name"`
	AuthorUrl     string `form:"author_url" json:"author_url"`
	ImageUrl      string `form:"image_url" json:"image_url"`
	Content       string `form:"content" json:"content"`
	Views         int    `form:"views" json:"views"`
	CanEdit       bool   `form:"can_edit" json:"can_edit"`
}

type PageList struct {
	TotalCount int     `json:"total_count"`
	Pages      []*Page `json:"pages"`
	Offset     int     `json:"offset"`
	Limit      int     `json:"limit"`
	Count      int     `json:"count"`
}
