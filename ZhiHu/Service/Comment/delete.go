/*
@Title : delete
@Description :
@Author : 谭靖渝
@Update : 2021/5/13 20:33
*/
package Comment

import (
	"ZhiHu/Model"
	"ZhiHu/Respond"
)

type Del struct {
	CCID    int `json:"cc_id" form:"cc_id"  gorm:"cc_id"`
	UserID string `json:"user_id" form:"user_id" gorm:"user_id"`
}

func (del *Del)DelCom()Respond.Response  {
	Model.DelComment(del.UserID,del.CCID)
	return Respond.Response{
		Code:  200,
		Msg:   "删除成功",
	}
}