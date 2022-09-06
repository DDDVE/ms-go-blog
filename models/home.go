package models

import "ms-go-blog/config"

type HomeResponse struct {
	config.Viewer
	Categorys []Category
	Posts []PostMore
	Total int
	Page int
	Pages []int
	PageEnd bool
}
