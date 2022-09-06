package dao

import (
	"log"
	"ms-go-blog/models"
)

func CountGetAllPost() int {
	row := DB.QueryRow("select count(1) from blog_post")
	if row.Err() != nil {
		log.Println("查询post总数出现错误：", row.Err())
	}
	count := 0
	_ = row.Scan(&count)
	return count
}

func GetPostPage(page, pageSize int) ([]models.Post, error) {
	rows, err := DB.Query("select * from blog_post limit ?,?", (page - 1) * pageSize, pageSize)
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
