package models

import "time"

//内存对齐

type UserTable struct {
	Id         int64     `json:"id,string" gorm:"primaryKey;autoIncrement" form:"id"`
	UserId     int64     `json:"user_id,string" gorm:"type:bigint(20);not null;unique" form:"user_id"`
	UserName   string    `json:"user_name" gorm:"type:varchar(64);not null；unique" form:"user_name"`
	Password   string    `json:"password" gorm:"type:varchar(64);not null" form:"password"`
	Email      string    `json:"email" gorm:"type:varchar(11);not null" form:"email"`
	Gender     int8      `json:"gender" gorm:"type:tinyint(4);not null;default:0" form:"gender"`
	CreateTime time.Time `json:"create_time" gorm:"not null" form:"create_time"`
	UpdateTime time.Time `json:"update_time" gorm:"not null" form:"update_time"`
}

func (UserTable) TableName() string {
	return "user"
}

type PostTable struct {
	Id          int64     `json:"id,string" form:"id" gorm:"primaryKey;autoIncrement"`
	PostId      int64     `json:"post_id,string" form:"post_id" gorm:"unique;not null"`
	Title       string    `json:"title" form:"title" gorm:"type:varchar(128);not null" binding:"required"`
	Content     string    `json:"content" form:"content" gorm:"type:varchar(8192);not null" binding:"required"`
	AuthorId    int64     `json:"author_id,string" form:"user_id" gorm:"not null"`
	CommunityId int64     `json:"community_id,string" form:"community_id" gorm:"not null"`
	Status      int8      `json:"status" form:"status" gorm:"not null"` //默认未通过审核
	CreateTime  time.Time `json:"create_time" form:"create_time" gorm:"not null"`
	UpdateTime  time.Time `json:"update_time" form:"update_time" gorm:""`
}

func (PostTable) TableName() string {
	return "post"
}

type CommunityTable struct {
	Id            int64     `json:"id,string"  gorm:"primaryKey;autoIncrement" form:"id"`
	CommunityId   int64     `json:"community_id,string" gorm:"not null;unique" form:"community_id"`
	CommunityName string    `json:"community_name" gorm:"type:varchar(128);not null;unique" form:"community_name"`
	Introduction  string    `json:"introduction" gorm:"type:varchar(256);not null;unique" form:"introduction"`
	CreateTime    time.Time `json:"create_time" gorm:"not null" form:"create_time"`
	UpdateTime    time.Time `json:"update_time" gorm:"not null" form:"update_time"`
}

func (CommunityTable) TableName() string {
	return "community"
}
