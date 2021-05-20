/*
@Title : deliver
@Description :
@Author : 谭靖渝
@Update : 2021/5/13 15:55
*/
package Comment

import (
	"ZhiHu/Model"
	"ZhiHu/Respond"
	"ZhiHu/cache/RedisServe"
	"fmt"
)

type Deliver struct {
	CID  int `json:"c_id"`
	Content string `json:"content"`
	UserID string `json:"user_id"`
}


//因为要发表评论，用户一定要登录，这样用户的信息就在redis里面了
func (deliver *Deliver)DeliverComment()Respond.Response  {
	fmt.Println("进入了")
	fmt.Println("用户账号：",deliver.UserID,deliver.Content,deliver.CID)
	user := RedisServe.GetUser(deliver.UserID)
	comment := Model.Comment{
		CID: deliver.CID,
		Content: deliver.Content,
		UserID: deliver.UserID,
		UserNickname: user.NickName,
		UserAvatar:user.Avatar,
	}
	if err := Model.DB.Create(&comment).Error; err != nil {
		return Respond.DBError("评论失败", err)
	}
	return Respond.BuildCommentResponse(comment)
}
