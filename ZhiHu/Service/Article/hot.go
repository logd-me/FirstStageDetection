/*
@Title : hot
@Description :
@Author : 谭靖渝
@Update : 2021/5/10 15:28
*/
package Article

import (
	"ZhiHu/Model"
	"ZhiHu/Respond"
)


func GetHot()[]Respond.Response  {
	var  hotArticle []Model.Article
	var res []Respond.Response
	if err:=Model.DB.Model(&Model.Article{}).Order("view ASC").Limit(10).Find(&hotArticle).Error;err==nil{
		for i:=0;i< len(hotArticle);i++ {
			res = append(res, Respond.BuildArticleResponse(hotArticle[i]))
		}
	}else {
		res = append(res, Respond.Response{
			Code:  0,
			Msg:   "获取失败",
		})
	}
	return res
}
