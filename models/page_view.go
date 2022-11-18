package models

type PageViewRequest struct {
	Year  int `form:"year,default=-1"`
	Month int `form:"month,default=-1"`
	Day   int `form:"day,default=-1"`
	Hour  int `form:"hour,default=-1"`
}
