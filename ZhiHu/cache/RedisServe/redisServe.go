/*
@Title : adduser
@Description :
@Author : 谭靖渝
@Update : 2021/5/11 18:59
*/
package RedisServe

import (
	"ZhiHu/Model"
	"ZhiHu/cache"
	"github.com/garyburd/redigo/redis"
	"github.com/sirupsen/logrus"
	"strconv"
	"time"
)

func AddUser(user Model.User) {
	cache.Conn.Do("SELECT", 0)
	_, err := cache.Conn.Do("HMSet", user.UID, "userid", user.UID, "nickname", user.NickName,
		"status", user.Status, "avatar", user.Avatar, "CreatedAt", user.CreatedAt, "UpdatedAt", user.UpdatedAt, "PassWord", user.PassWord)
	if err != nil {
		logrus.Info("添加用户失败")
		return
	}
}

func GetUser(UID interface{}) Model.User {
	cache.Conn.Do("SELECT", 0)
	r2, err := redis.Values(cache.Conn.Do("HMGet", UID, "CreatedAt", "UpdatedAt", "userid", "PassWord", "nickname", "status", "avatar"))
	if err != nil {
		logrus.Info("获取用户失败")
	}
	timeFormat := "2006-01-02 15:04:05 +0800 CST"
	CreatedAt, _ := time.ParseInLocation(timeFormat, string(r2[0].([]uint8)), time.Local)
	UpdatedAt, _ := time.ParseInLocation(timeFormat, string(r2[1].([]uint8)), time.Local)
	return Model.User{
		CreatedAt: CreatedAt,
		UpdatedAt: UpdatedAt,
		UID:       string(r2[2].([]uint8)),
		PassWord:  string(r2[3].([]uint8)),
		NickName:  string(r2[4].([]uint8)),
		Status:    string(r2[5].([]uint8)),
		Avatar:    string(r2[6].([]uint8)),
	}
}

func DelUser(UID interface{}) {
	var delStr = []string{"CreatedAt", "UpdatedAt", "userid", "PassWord", "nickname", "status", "avatar"}
	cache.Conn.Do("SELECT", 0)
	cache.Conn.Do("hdel", redis.Args{}.Add(UID).AddFlat(delStr)...)
}

func AddArticle(article Model.Article) {
	cache.Conn.Do("SELECT", 1)
	cache.Conn.Do("HMSet", article.CID, "c_id", article.CID, "title", article.Title,
		"context", article.Context, "user_id", article.UserID, "like", article.Like, "u_like", article.ULike,
		"view", article.View, "category", article.Category, "createdAt", article.CreatedAt)
}

func RGetArticle(CID interface{}) Model.Article {
	cache.Conn.Do("SELECT", 1)
	r2, err := redis.Values(cache.Conn.Do("HMGet", CID, "createdAt", "c_id", "title", "context", "user_id", "like", "u_like", "view", "category"))
	if err != nil {
		logrus.Info("获取文章失败")
	}
	timeFormat := "2006-01-02 15:04:05 +0800 CST"
	cid, _ := strconv.Atoi(string(r2[1].([]uint8)))
	create, _ := time.ParseInLocation(timeFormat, string(r2[0].([]uint8)), time.Local)
	like, _ := strconv.Atoi(string(r2[5].([]uint8)))
	ulike, _ := strconv.Atoi(string(r2[6].([]uint8)))
	view, _ := strconv.Atoi(string(r2[7].([]uint8)))
	Category := string(r2[8].([]uint8))
	title := string(r2[2].([]uint8))
	UserID := string(r2[4].([]uint8))
	context := string(r2[3].([]uint8))

	return Model.Article{
		CreatedAt: create,
		Title:     title,
		Context:   context,
		CID:       cid,
		UserID:    UserID,
		Like:      like,
		ULike:     ulike,
		View:      view,
		Category:  Category,
	}
}

func JudgeAE(CID interface{}) int64 {
	cache.Conn.Do("SELECT", 1)
	reply, _ := cache.Conn.Do("exists", CID)
	return reply.(int64)
}

func RDelArticle(CID interface{})  {
	cache.Conn.Do("SELECT", 1)
	cache.Conn.Do("del",CID)
}
