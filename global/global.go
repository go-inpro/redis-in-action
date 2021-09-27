/*
 * @Description: Do not edit
 * @Date: 2021-09-27 13:53:11
 * @LastEditors: viletyy
 * @Author: viletyy
 * @LastEditTime: 2021-09-27 16:10:25
 * @FilePath: /redis-in-action/global/global.go
 */
package global

import (
	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

var (
	Conf      *Config
	Logger    *zap.SugaredLogger
	VoteRedis *redis.Client
)
