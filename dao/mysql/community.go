package mysql

import (
	"bluebell/models"
	"gorm.io/gorm"

	"go.uber.org/zap"
)

func GetCommunityList() (communityList []*models.Community, err error) {
	//sqlStr := "select community_id, community_name from community"
	//if err := db.Select(&communityList, sqlStr); err != nil {
	//	if err == sql.ErrNoRows {
	//		zap.L().Warn("there is no community in db")
	//		err = nil
	//	}
	//}
	//return
	communityList = make([]*models.Community, 0, 10)
	err = db.Model(&models.CommunityTable{}).Select("community_id", "community_name").Limit(10).Find(communityList).Error
	if err == gorm.ErrRecordNotFound {
		zap.L().Warn("there is no community in mysql")
		err = nil
	}
	return
}

// GetCommunityDetailByID 根据ID查询社区详情
func GetCommunityDetailByID(id int64) (community *models.CommunityDetail, err error) {
	community = new(models.CommunityDetail)
	//sqlStr := `select
	//		community_id, community_name, introduction, create_time
	//		from community
	//		where community_id = ?
	//`
	//if err := db.Get(community, sqlStr, id); err != nil {
	//	if err == sql.ErrNoRows {
	//		err = ErrorInvalidID
	//	}
	//}
	//return community, err
	err = db.Model(&models.CommunityTable{}).Where("community_id=?", id).Find(community).Error
	return
}
