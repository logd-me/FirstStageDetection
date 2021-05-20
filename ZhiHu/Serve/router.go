/*
@Title : router
@Description :
@Author : 谭靖渝
@Update : 2021/5/5 16:21
*/
package Serve

import (
	"ZhiHu/middleware"
	"github.com/gin-gonic/gin"

	"ZhiHu/Api"
)

func NewRouter() *gin.Engine {
	//设置默认引擎
	r := gin.Default()
	r.Use(middleware.Session("TanJinYu"))
	r.Use(middleware.CurrentUser())
	v1 := r.Group("/v1")
	{
		v1.POST("/ping",Api.Ping)
		v1.POST("/user/register",Api.UserRegister)
		v1.POST("/user/login",Api.UserLogin)
		v1.GET("/hot",Api.HotArticle)
		v1.GET("Article/:id",Api.GetArticle)
		v1.GET("Article/:id/comment",Api.GetComment)
	}
	user := v1.Group("")
	user.Use(middleware.JudgeLogin())
	{
		user.GET("user/me", Api.UserGet)
		user.DELETE("user/out", Api.UserOut)
		user.DELETE("user/del/Article",Api.DelArticle)
		user.DELETE("user/del/Comment",Api.DelComment)
		user.POST("/user/Change/PassWord",Api.UserChangePassWord)
		user.POST("/user/Change/NickName",Api.UserChangeNickName)
		user.POST("/user/Change/Avatar",Api.UserChangeAvatar)
		user.POST("/user/Deliver",Api.Deliver)
		user.POST("/user/Article/like",Api.YesLike)
		user.POST("/user/Article/ulike",Api.NoLike)
		user.GET("/users/Article/:uid",Api.GetAll)
		user.POST("user/comment",Api.SaveComment)
	}
	return r
}