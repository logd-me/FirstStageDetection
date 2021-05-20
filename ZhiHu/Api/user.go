/*
@Title : user
@Description :
@Author : 谭靖渝
@Update : 2021/5/5 14:53
*/
package Api

import (
	"ZhiHu/Respond"
	"ZhiHu/Service/User"
	"ZhiHu/cache/RedisServe"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

//用户登录
func UserLogin(c *gin.Context) {
	var s User.Login
	if err := c.ShouldBind(&s); err == nil {
		res := s.ULogin(c)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, Respond.Error(Respond.CodeFailedLogin, "登录失败", err))
	}
}

//用户注册
func UserRegister(c *gin.Context) {
	var s User.Register
	if err := c.ShouldBind(&s); err == nil {
		res := s.Register()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, Respond.Error(Respond.CodeFailedRegister, "注册失败", err))
	}
}

//获取用户当前信息
func UserGet(c *gin.Context) {
	if user, _ := c.Get("user"); user != nil {
		c.JSON(http.StatusOK, user)
	} else {
		res := Respond.Response{
			Msg: "获取用户信息失败",
		}
		c.JSON(http.StatusOK, res)
	}

}

//用户登出
func UserOut(c *gin.Context) {
	s := sessions.Default(c)
	RedisServe.DelUser(s.Get("user_uid"))
	s.Clear()
	s.Save()
	c.JSON(http.StatusOK, Respond.Response{
		Code: 0,
		Msg:  "登出成功",
	})
}

//用户修改密码
func UserChangePassWord(c *gin.Context) {
	var change User.Change
	if err := c.ShouldBind(&change); err == nil {
		res, ok := change.ChangePassword()
		if ok {
			UserOut(c)
			c.Redirect(http.StatusMovedPermanently, "127.0.0.1:9090/v1/user/login")
		}
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, Respond.ParamError("修改密码失败", err))
	}
}
//修改昵称
func UserChangeNickName(c *gin.Context) {
	var change User.Change
	if err := c.ShouldBind(&change); err == nil {
		res := change.ChangeNickName()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, Respond.ParamError("修改昵称失败", err))
	}
}
//修改头像
func UserChangeAvatar(c *gin.Context) {
	var change User.Change
	if err := c.ShouldBind(&change); err == nil {
		res := change.ChangeAvatar()
		c.JSON(http.StatusOK, res)
		fmt.Println(change)
	} else {
		c.JSON(http.StatusOK, Respond.ParamError("修改头像失败", err))
	}
}
