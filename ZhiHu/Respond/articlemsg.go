/*
@Title : articlemsg
@Description :
@Author : 谭靖渝
@Update : 2021/5/9 10:06
*/
package Respond

import "ZhiHu/Model"

//文章序列化
type Article struct {
	Title    string `json:"title"`
	Context  string `json:"context"`
	CID      int `json:"cid;primary_key"`
	UserID   string `json:"user_id"`
	UserNickname string `json:"user_nickname"`
	UserAvatar string `json:"user_avatar"`
	Like     int   `json:"like"`
	ULike    int   `json:"u_like"`
	View     int    `json:"view"`
	Category string `json:"category"`
}

func buildArticle(article Model.Article) Article {
	user, _ := Model.GetUser(article.UserID)
	return Article{
		Title:    article.Title,
		Context:  article.Context,
		CID:      article.CID,
		UserID:   article.UserID,
		Like:     article.Like,
		ULike:    article.ULike,
		View: article.View,
		Category: article.Category,
		UserAvatar: user.Avatar,
		UserNickname: user.NickName,
	}
}

func BuildArticleResponse(article Model.Article) Response {
	return Response{
		Code: CodeArticleInfo,
		Data: buildArticle(article),
		Msg:  "文章信息",
	}
}
