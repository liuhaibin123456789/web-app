package mysql

import (
	"web_app/model"
)

func InsertCommunity(community *model.Community) (err error) {
	err = GDB.Model(&model.Community{}).Create(community).Error
	return
}

func SelectCommunity() (rcs []model.ResCommunity, err error) {
	rcs = make([]model.ResCommunity, 0)
	//指定Community的列名和ResCommunity有相同时，无需select指定，指定反而会冲突，
	//因此采用下面这种方案，查询哪几个字段，就把Community里的字段抽象出来形成新的结构体
	err = GDB.Model(&model.Community{}).Limit(10).Find(&rcs).Error
	return
}

func SelectDetailCommunity(communityId int64) (c *model.Community, err error) {
	c = new(model.Community)
	err = GDB.Model(&model.Community{}).Where("community_id=?", communityId).Find(c).Error
	return
}

func SelectCommunityIsExist(communityId int64) (isExist bool) {
	if err := GDB.Model(&model.Community{}).Where("community_id=?", communityId).Find(&model.Community{}).Error; err != nil {
		isExist = true
	}
	return
}
