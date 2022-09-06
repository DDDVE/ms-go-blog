package service

import (
	"html/template"
	"ms-go-blog/config"
	"ms-go-blog/dao"
	"ms-go-blog/models"
)

func GetPostsByCategoryId(cId, page, pageSize int) (*models.CategoryResponse, error) {
	var categorys, err = dao.GetAllCategory()
	if err != nil {
		return nil, err
	}
	posts, err := dao.GetPostPageByCategoryId(cId, page, pageSize)
	var postMores []models.PostMore
	for _, post := range posts {
		categoryName := dao.GetCategoryNameById(post.CategoryId)
		userName := dao.GetUserNameById(post.UserId)
		content := []rune(post.Content)
		if len(content) > 100 {
			content = content[ : 100]
		}
		postMore := models.PostMore{
			Pid: post.Pid,
			Title: post.Title,
			Slug: post.Slug,
			Content: template.HTML(content),
			CategoryId: post.CategoryId,
			CategoryName: categoryName,
			UserName: userName,
			ViewCount: post.ViewCount,
			Type: post.Type,
			CreateAt: models.DateDay(post.CreateAt),
			UpdateAt: models.DateDay(post.UpdateAt),
		}
		postMores = append(postMores, postMore)
	}
	total := dao.CountGetAllPostByCategoryId(cId)
	pagesCount := (total - 1) / pageSize + 1
	pages := make([]int, pagesCount)
	for i := 0; i < pagesCount; i++ {
		pages[i] = i + 1
	}
	var hr = &models.HomeResponse{
		config.Cfg.Viewer,
		categorys,
		postMores,
		total,
		page,
		pages,
		page != pagesCount,
	}
	categoryName := dao.GetCategoryNameById(cId)
	categoryResponse := &models.CategoryResponse{
		hr,
		categoryName,
	}
	return categoryResponse, nil
}
