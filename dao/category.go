package dao

import (
	"log"
	"ms-go-blog/models"
)

func GetPostPageByCategoryId(cId, page, pageSize int) ([]models.Post, error) {
	rows, err := DB.Query("select * from blog_post where category_id = ? limit ?,?", cId, (page - 1) * pageSize, pageSize)
	if err != nil {
		return nil, err
	}
	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err = rows.Scan(
			&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt,
		)
		if err != nil {
			log.Println("扫描数据出错：", err)
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func GetCategoryNameById(cId int) string {
	row := DB.QueryRow("select name from blog_category where  cid = ?", cId)
	if row.Err() != nil {
		log.Println(row.Err())
	}
	var categoryName string
	_ = row.Scan(&categoryName)
	return categoryName
}

func GetAllCategory() ([]models.Category, error) {
	rows, err := DB.Query("select * from blog_category")
	if err != nil {
		log.Fatal("GetAllCategory查询出错: ", err)
		return nil, err
	}
	var categorys []models.Category
	for rows.Next() {
		var category models.Category
		err = rows.Scan(&category.Cid, &category.Name, &category.CreateAt, &category.UpdateAt)
		if err != nil {
			log.Fatal("GetAllCategory取值出错: ", err)
			return nil, err
		}
		categorys = append(categorys, category)
	}
	return categorys, nil
}
