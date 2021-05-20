/*
@Title : delete
@Description :
@Author : 谭靖渝
@Update : 2021/5/12 14:55
*/
package Article

import (
	"ZhiHu/Model"
	"ZhiHu/Respond"
	"ZhiHu/cache/RedisServe"
)

type Del struct {
	CID    int `json:"c_id" form:"c_id"  gorm:"c_id"`
	UserID string `json:"user_id" form:"user_id" gorm:"user_id"`
}

func (del *Del)DelArt()Respond.Response{
	Model.DelArticle(del.UserID,del.CID)
	RedisServe.RDelArticle(del.CID)
	return Respond.Response{
		Code:  200,
		Msg:   "删除成功",
	}
}