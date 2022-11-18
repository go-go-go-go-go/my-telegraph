package models

import "encoding/json"

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
	Content       []any  `form:"content" json:"content,omitempty"`
	Views         int    `form:"views" json:"views,omitempty"`
	CanEdit       bool   `json:"can_edit"`
}

type PageRequest struct {
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
}

func MakePage(req *PageRequest) (*Page, error) {
	page := &Page{
		AccessToken:   req.AccessToken,
		ReturnContent: req.ReturnContent,
		Path:          req.Path,
		Url:           req.Url,
		Title:         req.Title,
		Description:   req.Description,
		AuthorName:    req.AuthorName,
		AuthorUrl:     req.AuthorUrl,
		ImageUrl:      req.ImageUrl,
		Content:       []any{},
	}
	err := json.Unmarshal([]byte(req.Content), &page.Content)
	if err != nil {
		return nil, err
	} else {
		return page, nil
	}
}

type PageList struct {
	TotalCount int     `json:"total_count"`
	Pages      []*Page `json:"pages"`
	Offset     int     `json:"offset"`
	Limit      int     `json:"limit"`
	Count      int     `json:"count"`
}
