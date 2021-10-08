/*
 * @Description: Do not edit
 * @Date: 2021-09-27 16:21:44
 * @LastEditors: viletyy
 * @Author: viletyy
 * @LastEditTime: 2021-10-08 15:24:18
 * @FilePath: /redis-in-action/vote/model/article.go
 */
package model

import (
	"fmt"
	"go-inpro/redis-in-action/global"
	"time"

	"github.com/go-redis/redis"
	"github.com/viletyy/yolk/convert"
)

const (
	ONE_WEEK_IN_SECONDS = 604800
	VOTE_SCORE          = 432
)

type Article struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Link   string  `json:"link"`
	Poster string  `json:"poster"`
	Time   float64 `json:"time"`
	Votes  int64   `json:"votes"`
}

func NewArticle(title, link, poster string) *Article {
	now := time.Now().Unix()
	pipe := global.VoteRedis.Pipeline()
	id, _ := global.VoteRedis.Incr("article:").Result()
	articleId := fmt.Sprintf("%d", id)
	pipe.SAdd("voted:article:"+articleId, poster)
	pipe.Expire("voted:article:"+articleId, ONE_WEEK_IN_SECONDS*time.Second)
	pipe.HMSet("article:"+articleId, map[string]interface{}{
		"title":  title,
		"link":   link,
		"poster": poster,
		"time":   float64(now),
		"votes":  1.0,
	})

	pipe.ZAdd("time:", redis.Z{Score: float64(now), Member: "article:" + articleId})
	pipe.ZAdd("score:", redis.Z{Score: float64(now) + VOTE_SCORE, Member: "article:" + articleId})
	_, err := pipe.Exec()
	if err != nil {
		return &Article{}
	}

	return &Article{
		ID:     articleId,
		Title:  title,
		Link:   link,
		Poster: poster,
		Time:   float64(now),
		Votes:  1,
	}
}

func GetArtcile(id string) *Article {
	article, _ := global.VoteRedis.HGetAll("article:" + id).Result()
	timeInt, _ := convert.StrTo(article["time"]).Float64()
	voteInt, _ := convert.StrTo(article["votes"]).Int64()
	return &Article{
		ID:     id,
		Title:  article["title"],
		Link:   article["link"],
		Poster: article["poster"],
		Time:   timeInt,
		Votes:  voteInt,
	}
}

func (article *Article) ToString() string {
	return "article:" + article.ID
}
