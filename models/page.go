package models

type Page struct {
	Id            int    `json:"-"`
	AccountId     int    `json:"-"`
	AccessToken   string `form:"access_token" json:"-"`
	ReturnContent bool   `form:"return_content" json:"-"`
	Path          string `json:"path"`
	Url           string `json:"url"`
	Title         string `form:"title" json:"title"`
	Description   string `form:"description" json:"description,omitempty"`
	AuthorName    string `form:"author_name" json:"author_name,omitempty"`
	AuthorUrl     string `form:"author_url" json:"author_url,omitempty"`
	ImageUrl      string `form:"image_url" json:"image_url,omitempty"`
	Content       string `form:"content" json:"content,omitempty"`
	Views         int    `form:"views" json:"views,omitempty"`
	CanEdit       bool   `json:"can_edit"`
}

type PageList struct {
	TotalCount int     `json:"total_count"`
	Pages      []*Page `json:"pages"`
	Offset     int     `json:"offset"`
	Limit      int     `json:"limit"`
	Count      int     `json:"count"`
}
