package api

import (
	"github.com/gin-gonic/gin"
	"web_app/global"
	"web_app/model"
	"web_app/service"
	"web_app/tool"
)

// CreatePost
// @tags post
// @Summary 创建帖子
// @Description
// @Accept json
// @Produce json
// @Param post body model.ReqPost true "帖子json数据"
// @Success 200 {object} tool.ResJson{msg=string}
// @Failure 200 {object} tool.ResJson{msg=string}
// @Security CoreAPI
// @Router /post [post]
func CreatePost(c *gin.Context) {
	p := new(model.Post)
	err := c.ShouldBind(p)
	if err != nil {
		//日志记录真正的原本错误
		tool.SugaredWarn("func CreatePost(c *gin.Context)", err)
		//前端返回的是对前端有好的错误
		tool.ResponseError(c, global.CodeInvalidParam)
		return
	}
	p.UserId = c.GetInt64("user_id")
	err = service.CreatePost(p)
	if err != nil {
		tool.SugaredError("service.CreatePost(p)", err)
		tool.ResponseError(c, global.CodeServerBusy)
		return
	}
	tool.ResponseSuccess(c, nil)
}

func GetPost(c *gin.Context) {
	//默认一页查询10条数据
	page := c.Query("page")
	posts, err := service.GetPost(page)
	if err != nil {
		//不将错误暴露给前端,日志记录错误
		tool.SugaredError("service.GetPost(page):", err)
		//错误不返回前端
		tool.ResponseError(c, global.CodeInvalidPathParam)
		return
	}
	tool.ResponseSuccess(c, posts)
}

// GetPost2
// @tags post
// @Summary 获取帖子列表
// @Description 支持时间和投票分数排序和查找社区分类下的帖子
// @Produce json
// @Param page query string false "查询的页码"
// @Param size query string false "查询的单页数据"
// @Param order query string false "只有两个值：`time`表示结果按照时间排序返回；`score`表示结果按照分数排序返回 "
// @Param community_id query string false "为空表示默认不按照社区分类查询；不为空，将按照所给id对应社区分类的帖子返回"
// @Success 200 {object} tool.ResJson{msg=string,data=[]model.Post}
// @Failure 200 {object} tool.ResJson{msg=string}
// @Security CoreAPI
// @Router /post2 [get]
func GetPost2(c *gin.Context) {
	// ?page=&size=&order=
	page := c.DefaultQuery("page", "1")
	size := c.DefaultQuery("size", "10")
	order := c.DefaultQuery("order", global.OrderTime)
	communityId := c.Query("community_id")

	tool.SugaredDebugf("func GetPost2(c *gin.Context): order:%s,size:%s,page:%s,communityId:%s", order, size, page, communityId)
	posts, err := service.GetPost2(page, size, order, communityId)
	if err != nil {
		//不将错误暴露给前端,日志记录错误
		tool.SugaredError("func GetDetailPost(c *gin.Context): ", err)
		//错误不返回前端
		tool.ResponseError(c, global.CodeInvalidPathParam)
		return
	}
	tool.ResponseSuccess(c, posts)
}

// GetDetailPost
// @tags post
// @Summary 获取帖子详情
// @Produce json
// @Param post_id path string true "查询帖子的post_id"
// @Success 200 {object} tool.ResJson{msg=string,data=model.ResPost}
// @Failure 200 {object} tool.ResJson{msg=string}
// @Security CoreAPI
// @Router /post/{post_id} [get]
func GetDetailPost(c *gin.Context) {
	pId := c.Param("post_id")
	resPost, err := service.GetDetailPost(pId)
	if err != nil {
		//不将错误暴露给前端,日志记录错误
		tool.SugaredError("service.GetDetailPost(pId)", err)
		//错误不返回前端
		tool.ResponseError(c, global.CodeInvalidPathParam)
		return
	}
	tool.ResponseSuccess(c, resPost)
}

// VoteForPost
// @tags post
// @Summary 帖子投票
// @Description
// @Accept json
// @Produce json
// @Param post_id body string true "投票帖子post_id"
// @Param direction body string true "投票帖子，1表示投票赞成，-1表示反对，0表示不投票或取消投票" maxlength(1)
// @Success 200 {object} tool.ResJson{msg=string}
// @Failure 200 {object} tool.ResJson{msg=string}
// @Security CoreAPI
// @Router /vote [post]
func VoteForPost(c *gin.Context) {
	//获取并校验参数
	v := new(model.ReqVoteData)
	err := c.ShouldBind(v)
	if err != nil {
		//不将错误暴露给前端,日志记录错误
		tool.SugaredWarn("func VoteForPost(c *gin.Context)", err)
		//错误不返回前端
		tool.ResponseError(c, global.CodeInvalidPathParam)
		return
	}
	uId := c.GetInt64("user_id")
	err = service.VoteForPost(uId, *v)
	if err != nil {
		//不将错误暴露给前端,日志记录错误
		tool.SugaredError("service.VoteForPost(*v):", err)
		//错误不返回前端
		tool.ResponseError(c, global.CodeFailedVote)
		return
	}
	tool.ResponseSuccess(c, nil)
}
