package model

import "time"

//Community 社区分类
type Community struct {
	Id            int64     `json:"id,string"  gorm:"primaryKey;autoIncrement" form:"id"`
	CommunityId   int64     `json:"community_id,string" gorm:"not null;unique" form:"community_id"`
	CommunityName string    `json:"community_name" gorm:"type:varchar(128);not null;unique" form:"community_name" binding:"required"`
	Introduction  string    `json:"introduction" gorm:"type:varchar(256);not null;unique" form:"introduction"`
	CreateTime    time.Time `json:"create_time" gorm:"not null" form:"create_time"`
	UpdateTime    time.Time `json:"update_time" gorm:"not null" form:"update_time"`
}

func (Community) TableName() string {
	return "community"
}
