/*
 * @Description: Do not edit
 * @Date: 2021-09-27 11:40:09
 * @LastEditors: viletyy
 * @Author: viletyy
 * @LastEditTime: 2021-09-27 15:07:11
 * @FilePath: /redis-in-action/util/config.go
 */
package util

import (
	"fmt"
	"go-inpro/redis-in-action/global"
	"os"

	"github.com/spf13/viper"
	"github.com/viletyy/yolk/convert"
)

func init() {
	var config = viper.New()
	config.SetConfigName("config")
	config.AddConfigPath("config")

	err := config.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}

	conf := &global.Config{}
	conf.Env = os.Getenv("ENV")
	if conf.Env == "" {
		conf.Env = config.GetString("env")
	}

	conf.Log.Dir = os.Getenv("LOG_DIR")
	if conf.Log.Dir == "" {
		conf.Log.Dir = config.GetString("log.dir")
	}

	conf.Log.File = os.Getenv("LOG_File")
	if conf.Log.File == "" {
		conf.Log.File = config.GetString("log.file")
	}

	conf.Log.MaxSize, _ = convert.StrTo(os.Getenv("LOG_MAXSIZE")).Int64()
	if conf.Log.MaxSize == 0 {
		conf.Log.MaxSize, _ = convert.StrTo(config.GetString("log.max_size")).Int64()
	}

	conf.Log.MaxBackups, _ = convert.StrTo(os.Getenv("LOG_MAXBACKUPS")).Int64()
	if conf.Log.MaxBackups == 0 {
		conf.Log.MaxBackups, _ = convert.StrTo(config.GetString("log.max_backups")).Int64()
	}

	conf.Log.MaxAge, _ = convert.StrTo(os.Getenv("LOG_MAXAGE")).Int64()
	if conf.Log.MaxAge == 0 {
		conf.Log.MaxAge, _ = convert.StrTo(config.GetString("log.max_age")).Int64()
	}

	conf.Redis.Host = os.Getenv("REDIS_HOST")
	if conf.Redis.Host == "" {
		conf.Redis.Host = config.GetString("redis.host")
	}

	conf.Redis.Port, _ = convert.StrTo(os.Getenv("REDIS_PORT")).Int64()
	if conf.Redis.Port == 0 {
		conf.Redis.Port, _ = convert.StrTo(config.GetString("redis.port")).Int64()
	}

	conf.Redis.Password = os.Getenv("REDIS_PASSWORD")
	if conf.Redis.Password == "" {
		conf.Redis.Password = config.GetString("redis.password")
	}

	conf.Redis.Db.Vote, _ = convert.StrTo(os.Getenv("REDIS_DB_VOTE")).Int64()
	if conf.Redis.Db.Vote == 0 {
		conf.Redis.Db.Vote, _ = convert.StrTo(config.GetString("redis_db.vote")).Int64()
	}

	global.Conf = conf
}
