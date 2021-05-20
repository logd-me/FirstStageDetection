/*
@Title : deliver
@Description :
@Author : 谭靖渝
@Update : 2021/5/9 10:30
*/
package Article

import (
	"ZhiHu/Model"
	"ZhiHu/Respond"
)

//发表文章服务结构体
type Deliver struct {
	Title string `json:"title" form:"title" gorm:"title"`
	Context string `json:"context" form:"context" gorm:"context"`
	CID int `json:"c_id" form:"c_id" gorm:"c_id"`
	UserID string `json:"user_id" form:"user_id" gorm:"user_id"`
	Category string `json:"category" form:"category" gorm:"category"`
}

func (deliver *Deliver)DeliverArticle()Respond.Response {
	article := Model.Article{
		Title:deliver.Title,
		Context: deliver.Context,
		UserID:deliver.UserID,
		Category:deliver.Category,
	}
	if err := Model.DB.Create(&article).Error; err != nil {
		return Respond.DBError("发布失败", err)
	}
	return Respond.BuildArticleResponse(article)
}
