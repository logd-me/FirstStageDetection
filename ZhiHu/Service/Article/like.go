/*
@Title : Like
@Description :
@Author : 谭靖渝
@Update : 2021/5/9 10:03
*/
package Article

import (
	"ZhiHu/Model"
	"ZhiHu/Respond"
	"ZhiHu/cache"
	"ZhiHu/cache/RedisServe"
)

//点赞或者差评
//服务点赞的结构体
type Think struct {
	CID    int    `json:"c_id" form:"c_id"  gorm:"c_id"`
	UserID string `json:"user_id" form:"user_id" gorm:"user_id"`
	Like   int    `json:"like" form:"like" gorm:"like"`
	ULike  int    `json:"u_like"  form:"u_like" gorm:"u_like"`
}

//像点赞啊，热度之类的，这些数据要修改不用每次都来连接数据库，先放在redis里面，等过了一段时间再来添加上去
func (think *Think) Thinks(judge int) Respond.Response {
	cache.Conn.Do("SELECT", 1)
	var article Model.Article
	switch judge {
	case 1:
		if RedisServe.JudgeAE(think.CID) == 0 {
			if err := Model.DB.Model(&article).Where(" c_id = ?", think.CID).First(&article).Update("like", article.Like+1).Error; err != nil {
				RedisServe.AddArticle(article)
				return Respond.Response{
					Code: 0,
					Msg:  err.Error(),
				}
			}
		} else {
			res := RedisServe.RGetArticle(think.CID)
			cache.Conn.Do("HMSet", think.CID, "like", res.Like+1)
		}
	case 0:
		if RedisServe.JudgeAE(think.CID) == 0 {
			if err := Model.DB.Model(&article).Where(" c_id = ?", think.CID).First(&article).Update("u_like", article.ULike+1).Error; err != nil {
				RedisServe.AddArticle(article)
				return Respond.Response{
					Code: 0,
					Msg:  err.Error(),
				}
			}
		} else {
			res := RedisServe.RGetArticle(think.CID)
			cache.Conn.Do("HMSet", think.CID, "u_like", res.ULike+1)
		}

	}
	return Respond.Response{
		Code: 200,
		Msg:  "修改成功",
	}

}
