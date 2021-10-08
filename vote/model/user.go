/*
 * @Description: Do not edit
 * @Date: 2021-09-27 16:21:09
 * @LastEditors: viletyy
 * @Author: viletyy
 * @LastEditTime: 2021-10-08 10:30:27
 * @FilePath: /redis-in-action/vote/model/user.go
 */
package model

import (
	"errors"
	"go-inpro/redis-in-action/global"
	"time"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func NewUser(id, name string) *User {
	global.VoteRedis.HSet("user:"+id, "name", name)

	return &User{ID: id, Name: name}
}

func (user *User) VoteArticle(article *Article) error {
	cutoff := time.Now().Unix() - ONE_WEEK_IN_SECONDS

	zscoreResult, err := global.VoteRedis.ZScore("time:", article.ToString()).Result()
	if err != nil {
		global.Logger.Errorf("ZScore err: %v", err)
		return err
	}
	if zscoreResult < float64(cutoff) {
		return errors.New("Article is not permit vote:")
	}

	pipe := global.VoteRedis.Pipeline()
	pipe.SAdd("voted:"+article.ToString(), user.ToString())
	pipe.ZIncrBy("score:", VOTE_SCORE, article.ToString())
	pipe.HIncrBy(article.ToString(), "votes", 1)
	_, err = pipe.Exec()

	return nil
}

func GetUser(id string) *User {
	dbUser, _ := global.VoteRedis.HGetAll("user:" + id).Result()
	return &User{
		ID:   id,
		Name: dbUser["name"],
	}
}

func (user *User) ToString() string {
	return "user:" + user.ID
}
