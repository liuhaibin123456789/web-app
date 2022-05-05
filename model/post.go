package model

import "time"

type Post struct {
	Id          int64     `json:"id,string" form:"id" gorm:"primaryKey;autoIncrement"`
	PostId      int64     `json:"post_id,string" form:"post_id" gorm:"unique;not null"`
	UserId      int64     `json:"user_id,string" form:"user_id" gorm:"not null"`
	CommunityId int64     `json:"community_id,string" form:"community_id" gorm:"not null" binding:"required"`
	Status      int8      `json:"status" form:"status" gorm:"not null"` //默认未通过审核
	Title       string    `json:"title" form:"title" gorm:"type:varchar(128);not null" binding:"required"`
	Content     string    `json:"content" form:"content" gorm:"type:varchar(8192);not null" binding:"required"`
	CreateTime  time.Time `json:"create_time" form:"create_time" gorm:"not null"`
	UpdateTime  time.Time `json:"update_time" form:"update_time" gorm:""`
}

func (Post) TableName() string {
	return "post"
}
