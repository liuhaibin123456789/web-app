package mysql

import (
	"bluebell/models"
	"strings"
)

// CreatePost 创建帖子
func CreatePost(p *models.PostTable) (err error) {
	//sqlStr := `insert into post(
	//post_id, title, content, author_id, community_id)
	//values (?, ?, ?, ?, ?)
	//`
	//_, err = db.Exec(sqlStr, p.ID, p.Title, p.Content, p.AuthorID, p.CommunityID)
	//return
	return db.Model(&models.Post{}).Create(p).Error
}

// GetPostById 根据id查询单个贴子数据
func GetPostById(pid int64) (post *models.PostTable, err error) {
	post = new(models.PostTable)
	//sqlStr := `select
	//post_id, title, content, author_id, community_id, create_time
	//from post
	//where post_id = ?
	//`
	//err = db.Get(post, sqlStr, pid)
	//return
	err = db.Model(&models.PostTable{}).Where("post_id=?", pid).Find(post).Error
	return
}

// GetPostList 查询帖子列表函数
func GetPostList(page, size int64) (posts []*models.PostTable, err error) {
	posts = make([]*models.PostTable, 0) // 不要写成make([]*models.Post, 2)
	//sqlStr := `select
	//post_id, title, content, author_id, community_id, create_time
	//from post
	//ORDER BY create_time
	//DESC
	//limit ?,?
	//`
	//posts = make([]*models.Post, 0, 2) // 不要写成make([]*models.Post, 2)
	//err = db.Select(&posts, sqlStr, (page-1)*size, size)
	//return
	err = db.Model(&models.PostTable{}).Order("create_time DESC").Offset(int((page - 1) * size)).Limit(int(size)).Find(posts).Error
	return posts, err
}

// GetPostListByIDs 根据给定的id列表查询帖子数据
func GetPostListByIDs(ids []string) (posts []*models.PostTable, err error) {
	//sqlStr := `select post_id, title, content, author_id, community_id, create_time
	//from post
	//where post_id in (?)
	//order by FIND_IN_SET(post_id, ?)
	//`
	//query, args, err := sqlx.In(sqlStr, ids, strings.Join(ids, ","))
	//if err != nil {
	//	return nil, err
	//}
	//query = db.Rebind(query)
	//err = db.Select(&postList, query, args...)
	//return
	posts = make([]*models.PostTable, 0)
	post := new(models.PostTable)
	rows, err := db.Raw("select * from post where post_id in ? order by find_in_set(post_id,?)", ids, strings.Join(ids, ",")).Rows()
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		if err = rows.Scan(
			&post.Id,
			&post.PostId,
			&post.AuthorId,
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
