package mysql

import (
	"strings"
	"web_app/model"
)

func InsertPost(post *model.Post) (err error) {
	err = GDB.Model(&model.Post{}).Create(post).Error
	return
}

func SelectPost(page int) (posts []model.Post, err error) {
	posts = make([]model.Post, 0)
	err = GDB.Model(&model.Post{}).Order("id DESC").Offset((page - 1) * 10).Limit(10).Find(&posts).Error
	return
}

func SelectPost2(postIds []string) (posts []model.Post, err error) {
	posts = make([]model.Post, 0)
	post := model.Post{}
	rows, err := GDB.Raw("select * from post where post_id in ? order by find_in_set(post_id,?)", postIds, strings.Join(postIds, ",")).Rows()
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		if err = rows.Scan(
			&post.Id,
			&post.PostId,
			&post.UserId,
			&post.CommunityId,
			&post.Status,
			&post.Title,
			&post.Content,
			&post.CreateTime,
			&post.UpdateTime,
		); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return
}

func SelectDetailPost(postId int64) (post *model.Post, err error) {
	//查找post
	post = new(model.Post)
	err = GDB.Model(&model.Post{}).Where("post_id=?", postId).Find(post).Error
	return
}
