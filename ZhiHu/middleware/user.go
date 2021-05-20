/*
@Title : user
@Description :
@Author : 谭靖渝
@Update : 2021/5/5 15:43
*/
package middleware

import (
	"ZhiHu/Model"
	"ZhiHu/Respond"
	"ZhiHu/cache/RedisServe"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)
//登录判断
func JudgeLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		if user,_ := c.Get("user");user!=nil{
			c.Next()
			return
		}
		c.JSON(http.StatusOK,Respond.CheckLogin())
		c.Abort()
	}
}
//当前用户
func CurrentUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		if judge:=session.Get("data_type"); judge=="redis" {
			user := RedisServe.GetUser(session.Get("user_uid"))
			c.Set("user", &user)
			c.Next()
			return
		}
		if uid := session.Get("user_uid"); uid != nil {
			user, err := Model.GetUser(uid)
			if err == nil {
				c.Set("user", &user)
			}
			c.Next()
			return
		}
	}
}
