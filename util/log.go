/*
 * @Description: Do not edit
 * @Date: 2021-09-27 13:46:32
 * @LastEditors: viletyy
 * @Author: viletyy
 * @LastEditTime: 2021-09-27 15:41:09
 * @FilePath: /redis-in-action/util/log.go
 */
package util

import (
	"fmt"
	"go-inpro/redis-in-action/global"
	"io"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func init() {
	switch global.Conf.Env {
	case "development":
		InitLog(os.Stdout)
	case "production":
		hook := logToFileHook()
		InitLog(hook)
	default:
		InitLog(os.Stdout)
	}

	global.Logger.Debugf("Environment: %s", global.Conf.Env)
}

func logToFileHook() *lumberjack.Logger {
	return &lumberjack.Logger{
		Filename:   fmt.Sprintf("%s/%04d%02d%02d%02d%02d%02d.log", global.Conf.Log.Dir, time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), time.Now().Second()),
		MaxSize:    int(global.Conf.Log.MaxSize),
		MaxAge:     int(global.Conf.Log.MaxBackups),
		MaxBackups: int(global.Conf.Log.MaxAge),
		LocalTime:  false,
		Compress:   false,
	}
}

func InitLog(hook interface{}) {
	enConfig := zap.NewProductionEncoderConfig()
	enConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	level := zap.InfoLevel
	w := zapcore.AddSync(hook.(io.Writer))
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(enConfig),
		w,
		level,
	)

	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	global.Logger = logger.Sugar()
}
