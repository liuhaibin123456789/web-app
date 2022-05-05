package tool

import (
	"github.com/go-redis/redis"
	"time"
)

const (
	KeyPrefix              = "bluebell:"
	KeyPostTimeZSet        = "post:time"   //帖子及发帖时间
	KeyPostScoreZSet       = "post:score"  //帖子及投票的分数
	KeyPostVotedZSetPrefix = "post:voted:" //记录用户及投票类型
)

var RDB *redis.Client

func Redis() (err error) {
	client := redis.NewClient(&redis.Options{
		Addr:     V.GetString("redis.host") + ":" + V.GetString("redis.port"),
		Password: V.GetString("redis.password"),
		DB:       V.GetInt("redis.db"),
	})
	RDB = client
	return nil
}
func Set(keyUserId string, value interface{}, duration time.Duration) (err error) {
	err = RDB.Set(keyUserId, value, duration).Err()
	return
}

func Get(keyUserId string) (value string, err error) {
	value, err = RDB.Get(keyUserId).Result()
	return
}

func GetKey(key string) string {
	return KeyPrefix + key
}

func SetPost(postId int64) error {
	//redis事务
	pipeline := RDB.TxPipeline()
	//帖子时间
	pipeline.ZAdd(GetKey(KeyPostTimeZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postId,
	})
	//帖子分数
	pipeline.ZAdd(GetKey(KeyPostScoreZSet), redis.Z{
		Score:  0,
		Member: postId,
	})
	_, err := pipeline.Exec()
	return err

}
