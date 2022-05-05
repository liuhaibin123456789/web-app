package service

import (
	"errors"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"math"
	"strconv"
	"time"
	"web_app/dao/mysql"
	"web_app/global"
	"web_app/model"
	"web_app/tool"
)

const (
	oneWeekInSeconds = 7 * 24 * 3600 //一周的秒数
	scorePerVote     = 500           //每票分数
)

var (
	ErrorPostTitleLong    = errors.New("帖子标题太长")
	ErrorPostContentLong  = errors.New("帖子内容太长")
	ErrorPostNotBelong    = errors.New("帖子归属不明")
	ErrorPostVoteExpired  = errors.New("投票时间已过")
	ErrorPostVoteInvalid  = errors.New("投票数据错误")
	ErrorPostInvalidParam = errors.New("请求参数错误")
)

func CreatePost(post *model.Post) (err error) {
	//业务校验
	if post.UserId == 0 || post.CommunityId == 0 {
		return ErrorPostNotBelong
	}
	if len(post.Title) > 128 {
		return ErrorPostTitleLong
	}
	if len(post.Content) > 8192 {
		return ErrorPostContentLong
	}
	post.CreateTime = time.Now()
	post.PostId = tool.GetId()

	err = mysql.InsertPost(post)
	if err != nil {
		return err
	}

	return tool.SetPost(post.PostId)
}

func GetPost(page string) (posts []model.Post, err error) {
	p, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		return
	}
	return mysql.SelectPost(int(p))
}

func GetPost2(page, size, order, communityId string) (posts []model.Post, err error) {
	if order == global.OrderTime || order == global.OrderScore {
		p, err1 := strconv.ParseInt(page, 10, 64)
		if err1 != nil {
			err = err1
			return
		}
		s, err1 := strconv.ParseInt(size, 10, 64)
		if err1 != nil {
			err = err1
			return
		}
		tool.Debug("GetPost2", zap.String("order", order))
		//获取根据排序规则从redis里获取post_id列表
		key := tool.GetKey(tool.KeyPostTimeZSet)
		if order == global.OrderScore {
			key = tool.GetKey(tool.KeyPostScoreZSet)
		}
		start := (p - 1) * s
		end := start + s - 1
		ids, err1 := tool.RDB.ZRevRange(key, start, end).Result()
		if err1 != nil {
			err = err1
			return
		}
		tool.Debug("redis", zap.Any("ids", ids))
		//根据排好序的post_id列表获取帖子
		return mysql.SelectPost2(ids)
	}
	return nil, ErrorPostInvalidParam
}

func GetDetailPost(postId string) (resPost *model.ResPost, err error) {
	resPost = new(model.ResPost)
	pId, err := strconv.ParseInt(postId, 10, 64)
	if err != nil {
		return nil, err
	}
	//帖子数据
	post, err := mysql.SelectDetailPost(pId)
	if err != nil {
		return nil, err
	}
	//帖子分类信息
	community, err := mysql.SelectDetailCommunity(post.CommunityId)
	if err != nil {
		return nil, err
	}

	userName, err := mysql.SelectUserName(post.UserId)
	if err != nil {
		return nil, err
	}
	resPost.UserName = userName
	resPost.Post = post
	resPost.Community = community
	//帖子分数
	resPost.VoteScore, err = tool.RDB.ZScore(tool.KeyPostScoreZSet, strconv.FormatInt(resPost.PostId, 10)).Result()
	return
}

func VoteForPost(userId int64, data model.ReqVoteData) (err error) {
	if data.Direction == 1 || data.Direction == -1 || data.Direction == 0 {
		//判断投票限制:大于7天的帖子不能在投票
		postTime := tool.RDB.ZScore(tool.GetKey(tool.KeyPostTimeZSet), strconv.FormatInt(data.PostId, 10)).Val()
		if float64(time.Now().Unix())-postTime > oneWeekInSeconds {
			err = ErrorPostVoteExpired
			return
		}
		//更新帖子分数
		ov := tool.RDB.ZScore(tool.GetKey(tool.KeyPostVotedZSetPrefix+strconv.FormatInt(data.PostId, 10)), strconv.FormatInt(userId, 10)).Val()
		var dir float64
		if ov < data.Direction {
			dir = 1
		} else {
			dir = -1
		}
		diff := math.Abs(data.Direction - ov)

		txPipeline := tool.RDB.TxPipeline()

		err = txPipeline.ZIncrBy(tool.GetKey(tool.KeyPostScoreZSet), scorePerVote*diff*dir, strconv.FormatInt(data.PostId, 10)).Err()
		if err != nil {
			return
		}
		if data.Direction != 0 {
			//记录用户已为该帖子投票
			_, err = txPipeline.ZAdd(tool.GetKey(tool.KeyPostVotedZSetPrefix+strconv.FormatInt(userId, 10)), redis.Z{
				Score:  data.Direction,
				Member: userId,
			}).Result()
		}
		_, err = txPipeline.Exec()
		tool.Debug("VoteForPost", zap.Int64("user_id", userId), zap.Int64("post_id", data.PostId), zap.Float64("direction", data.Direction))
		return
	}
	return ErrorPostVoteInvalid
}
