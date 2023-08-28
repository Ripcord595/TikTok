package cache

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"tiktok/conf"
)

var ctx = context.Background()
var rdb *redis.Client

const (
	favor    = "favor"
	relation = "relation"
)

func init() {
	config := conf.NewConfig()

	rdb = redis.NewClient(
		&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", config.RDB.IP, config.RDB.Port),
			Password: "",
			DB:       config.RDB.Database,
		})
}

var (
	proxyIndexOperation ProxyIndexMap
)

type ProxyIndexMap struct {
}

func NewProxyIndexMap() *ProxyIndexMap {
	return &proxyIndexOperation
}

func (i *ProxyIndexMap) UpdateVideoFavorState(userId int64, videoId int64, state bool) {
	key := fmt.Sprintf("%s:%d", favor, userId)
	if state {
		rdb.SAdd(ctx, key, videoId)
		return
	}
	rdb.SRem(ctx, key, videoId)
}

func (i *ProxyIndexMap) GetVideoFavorState(userId int64, videoId int64) bool {
	key := fmt.Sprintf("%s:%d", favor, userId)
	ret := rdb.SIsMember(ctx, key, videoId)
	return ret.Val()
}

func (i *ProxyIndexMap) UpdateUserRelation(userId int64, followId int64, state bool) {
	key := fmt.Sprintf("%s:%d", relation, userId)
	if state {
		rdb.SAdd(ctx, key, followId)
		return
	}
	rdb.SRem(ctx, key, followId)
}

func (i *ProxyIndexMap) GetUserRelation(userId int64, followId int64) bool {
	key := fmt.Sprintf("%s:%d", relation, userId)
	ret := rdb.SIsMember(ctx, key, followId)
	return ret.Val()
}
