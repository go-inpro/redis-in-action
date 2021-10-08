/*
 * @Description: Do not edit
 * @Date: 2021-09-27 16:18:42
 * @LastEditors: viletyy
 * @Author: viletyy
 * @LastEditTime: 2021-10-08 15:25:22
 * @FilePath: /redis-in-action/vote/article.go
 */
package vote

import (
	"go-inpro/redis-in-action/global"
	"go-inpro/redis-in-action/vote/model"
	"strings"

	"github.com/viletyy/yolk/convert"
)

const (
	ARTICLES_PER_PAGE = 25
)

func PostArticle(user *model.User, title, link string) {
	model.NewArticle(title, link, user.ToString())
}

func ArticleVote(user *model.User, article *model.Article) {
	user.VoteArticle(article)
}

func GetArticles(page int, order string) []*model.Article {
	start := (page - 1) * ARTICLES_PER_PAGE
	end := start + ARTICLES_PER_PAGE - 1

	ids := global.VoteRedis.ZRevRange(order, int64(start), int64(end)).Val()

	articles := make([]*model.Article, len(ids))
	for index, _ := range articles {
		dbArticle, _ := global.VoteRedis.HGetAll(ids[index]).Result()
		timeFloat, _ := convert.StrTo(dbArticle["time"]).Float64()
		voteInt, _ := convert.StrTo(dbArticle["votes"]).Int64()
		articles[index] = &model.Article{
			ID:     strings.Split(ids[index], ":")[1],
			Title:  dbArticle["title"],
			Link:   dbArticle["link"],
			Poster: dbArticle["poster"],
			Time:   timeFloat,
			Votes:  voteInt,
		}
	}
	return articles
}
