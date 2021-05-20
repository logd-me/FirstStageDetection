/*
@Title : commentmsg
@Description :
@Author : 谭靖渝
@Update : 2021/5/13 15:38
*/
package Respond

import "ZhiHu/Model"

//评论序列化

type Comment struct {
	CCID int `json:"cc_id"`
	CID  int `json:"c_id"`
	Content string `json:"content"`
	UserID string `json:"user_id"`
	UserNickname string `json:"user_nickname"`
	UserAvatar string `json:"user_avatar"`
}

func buildComment(comment Model.Comment) Comment {
	return Comment{
		CCID :comment.CCID,
		CID:      comment.CID,
		Content: comment.Content,
		UserID:   comment.UserID,
		UserAvatar: comment.UserAvatar,
		UserNickname: comment.UserNickname,
	}
}

func BuildCommentResponse(comment Model.Comment) Response {
	return Response{
		Code: CodeArticleInfo,
		Data: buildComment(comment),
		Msg:  "文章信息",
	}
}
