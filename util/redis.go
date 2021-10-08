/*
 * @Description: Do not edit
 * @Date: 2021-09-27 16:05:40
 * @LastEditors: viletyy
 * @Author: viletyy
 * @LastEditTime: 2021-09-27 17:15:11
 * @FilePath: /redis-in-action/util/redis.go
 */
package util

import (
	"fmt"
	"go-inpro/redis-in-action/global"

	"github.com/go-redis/redis"
)

func init() {
	global.VoteRedis = InitRedis(global.Conf.Redis.Db.Vote)
}

func InitRedis(db int64) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", global.Conf.Redis.Host, global.Conf.Redis.Port),
		Password: global.Conf.Redis.Password,
		DB:       int(global.Conf.Redis.Db.Vote),
	})
	RedisSet(rdb)

	return rdb
}

func RedisSet(rdb *redis.Client) {
	_, pingErr := rdb.Ping().Result()
	if pingErr != nil {
		global.Logger.Error(fmt.Sprintf("Redis Connection Error: %v", pingErr))
	}
}
