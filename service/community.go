package service

import (
	"strconv"
	"web_app/dao/mysql"
	"web_app/model"
)

func GetCommunity() (rcs []model.ResCommunity, err error) {
	return mysql.SelectCommunity()
}

func CreateCommunity(community *model.Community) (err error) {
	return mysql.InsertCommunity(community)
}

func GetDetailCommunity(communityId string) (c *model.Community, err error) {
	c = new(model.Community)
	cId, err := strconv.ParseInt(communityId, 10, 64)
	if err != nil {
		return nil, err
	}
	return mysql.SelectDetailCommunity(cId)
}
