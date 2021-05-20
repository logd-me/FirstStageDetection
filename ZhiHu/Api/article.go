/*
@Title : article
@Description :
@Author : 谭靖渝
@Update : 2021/5/9 11:09
*/
package Api

import (
	"ZhiHu/Model"
	"ZhiHu/Respond"
	"ZhiHu/Service/Article"
	"ZhiHu/cache/RedisServe"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

//发布文章
func Deliver(c *gin.Context) {
	var deliver Article.Deliver
	if err := c.ShouldBind(&deliver); err == nil {
		res := deliver.DeliverArticle()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, Respond.ParamError("发布失败", err))
	}
}

//点赞
func YesLike(c *gin.Context) {
	var think Article.Think
	if err := c.ShouldBind(&think); err == nil {
		res := think.Thinks(1)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, Respond.ParamError("点赞失败", err))
	}
}

//差评
func NoLike(c *gin.Context) {
	var think Article.Think
	if err := c.ShouldBind(&think); err == nil {
		res := think.Thinks(0)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, Respond.ParamError("差评失败", err))
	}
}

//获取文章信息
func GetArticle(c *gin.Context) {
	cid, ok := c.Params.Get("id")
	if !ok {
		c.JSON(0, "无效的参数")
		return
	}
	session := sessions.Default(c)
	if session.Get("art_from")=="redis"&&RedisServe.JudgeAE(cid)==1 {
		res := RedisServe.RGetArticle(cid)
		fmt.Println(res.CID)
		res.ViewAdd()
		c.JSON(http.StatusOK, Respond.BuildArticleResponse(res))
	}else {
		if res, err := Model.GetArticle(cid); err == nil {
			res.ViewAdd()
			RedisServe.AddArticle(res)
			session.Set("art_from","redis")
			session.Save()
			c.JSON(http.StatusOK, Respond.BuildArticleResponse(res))
		} else {
			c.JSON(http.StatusOK, Respond.ParamError("获取信息失败", err))
		}
	}
}

//hot推荐
func HotArticle(c *gin.Context)  {
	c.JSON(http.StatusOK,Article.GetHot())
}

//获取某位用户的全部文章信息
func GetAll(c *gin.Context)  {
	var art []Model.Article
	uid, ok := c.Params.Get("uid")
	if !ok {
		c.JSON(0, "无效的参数")
		return
	}
	art= Model.GetAllArticle(uid)
	if art!=nil{
		var res []Respond.Response
		for i:=0;i< len(art);i++ {
			res = append(res, Respond.BuildArticleResponse(art[i]))
		}
		c.JSON(http.StatusOK,res)
	}else if len(art)==0 {
		c.JSON(http.StatusOK, Respond.Response{
			Code:  0,
			Msg:   "获取信息失败",
		})
	}
}

func DelArticle(c *gin.Context)  {
	var del Article.Del
	if err := c.ShouldBind(&del);err==nil{
		res := del.DelArt()
		c.JSON(http.StatusOK,res)
	}else {
		c.JSON(http.StatusOK,Respond.ParamError("删除失败", err))
	}
}

