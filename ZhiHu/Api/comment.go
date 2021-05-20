/*
@Title : comment
@Description :
@Author : 谭靖渝
@Update : 2021/5/13 20:08
*/
package Api

import (
	"ZhiHu/Model"
	"ZhiHu/Respond"
	"ZhiHu/Service/Comment"
	"github.com/gin-gonic/gin"
	"net/http"
)

//评论
func SaveComment(c *gin.Context)  {
	var deliver Comment.Deliver
	if err := c.ShouldBind(&deliver);err==nil{
		res :=deliver.DeliverComment()
		c.JSON(http.StatusOK,res)
	}else {
		c.JSON(http.StatusOK, Respond.ParamError("评论失败", err))
	}
}

//删除评论
func DelComment(c *gin.Context)  {
	var del Comment.Del
	if err := c.ShouldBind(&del);err==nil{
		res := del.DelCom()
		c.JSON(http.StatusOK,res)
	}else {
		c.JSON(http.StatusOK,Respond.ParamError("删除失败", err))
	}
}

//获取文章的全部评论
func GetComment(c *gin.Context)  {
	var com []Model.Comment
	cid, ok := c.Params.Get("id")
	if !ok {
		c.JSON(0, "无效的参数")
		return
	}
	com= Model.GetComment(cid)
	if com!=nil{
		var res []Respond.Response
		for i:=0;i< len(com);i++ {
			res = append(res, Respond.BuildCommentResponse(com[i]))
		}
		c.JSON(http.StatusOK,res)
	}else if len(com)==0 {
		c.JSON(http.StatusOK, Respond.Response{
			Code:  0,
			Msg:   "获取评论失败",
		})
	}
}
