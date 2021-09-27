/*
 * @Description: Do not edit
 * @Date: 2021-09-27 11:42:29
 * @LastEditors: viletyy
 * @Author: viletyy
 * @LastEditTime: 2021-09-27 15:07:06
 * @FilePath: /redis-in-action/global/config.go
 */
package global

type Config struct {
	Env   string
	Redis RedisConfig
	Log   LogConfig
}

type RedisConfig struct {
	Host     string
	Port     int64
	Db       RedisDb
	Password string
}

type RedisDb struct {
	Vote int64
}

type LogConfig struct {
	Dir        string
	File       string
	MaxSize    int64
	MaxBackups int64
	MaxAge     int64
}
