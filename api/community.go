package api

import (
	"github.com/gin-gonic/gin"
	"web_app/global"
	"web_app/model"
	"web_app/service"
	"web_app/tool"
)

// GetCommunity
// @tags community
// @Summary 获取社区分类信息
// @Description 获取前十个标签的id及名字，没有十个则返回所有
// @Produce json
// @Success 200 {object} tool.ResJson{data=[]model.ResCommunity}
// @Failure 200 {object} tool.ResJson{msg=string}
// @Security CoreAPI
// @Router /community [get]
func GetCommunity(c *gin.Context) {
	communities, err := service.GetCommunity()
	if err != nil {
		//不将错误暴露给前端,日志记录错误
		tool.SugaredError("service.GetCommunity()", err)
		//错误不返回前端
		tool.ResponseError(c, global.CodeServerBusy)
		return
	}
	tool.ResponseSuccess(c, communities)
}

// CreateCommunity
// @tags community
// @Summary 创建社区分类
// @Description
// @Accept json
// @Produce json
// @Param community_name body string true "分类名" minlength(1)  maxlength(128)
// @Param introduction body string false "介绍" maxlength(256)
// @Success 200 {object} tool.ResJson{msg=string}
// @Failure 200 {object} tool.ResJson{msg=string}
// @Security CoreAPI
// @Router /community [post]
func CreateCommunity(c *gin.Context) {
	community := new(model.Community)
	err := c.ShouldBind(community)
	if err != nil {
		//不将错误暴露给前端,日志记录错误
		tool.SugaredWarn("CreateCommunity", err)
		//错误不返回前端
		tool.ResponseError(c, global.CodeServerBusy)
		return
	}
	err = service.CreateCommunity(community)
	if err != nil {
		//不将错误暴露给前端,日志记录错误
		tool.SugaredError("service.CreateCommunity(community)", err)
		//错误不返回前端
		tool.ResponseError(c, global.CodeServerBusy)
		return
	}
	tool.ResponseSuccess(c, nil)

}

// GetDetailCommunity
// @tags community
// @Summary 获取社区分类详细信息
// @Description 获取单个分类信息数据
// @Produce json
// @Param community_id path string true "查询帖子的community_id"
// @Success 200 {object} tool.ResJson{data=model.Community}
// @Failure 200 {object} tool.ResJson{msg=string}
// @Security CoreAPI
// @Router /community/{community_id} [get]
func GetDetailCommunity(c *gin.Context) {
	cId := c.Param("community_id")
	community, err := service.GetDetailCommunity(cId)
	if err != nil {
		//不将错误暴露给前端,日志记录错误
		tool.SugaredError("service.GetDetailCommunity(cId)", err)
		//错误不返回前端
		tool.ResponseError(c, global.CodeInvalidPathParam)
		return
	}
	tool.ResponseSuccess(c, community)
}
