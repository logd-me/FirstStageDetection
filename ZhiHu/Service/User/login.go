/*
@Title : Login
@Description :
@Author : 谭靖渝
@Update : 2021/5/4 18:08
*/
package User

import (
	"ZhiHu/Model"
	"ZhiHu/Respond"
	"ZhiHu/cache/RedisServe"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//服务登录的结构体
type Login struct {
	UID      string `form:"uid" json:"uid" binding:"required,min=5,max=30"`
	PassWord string `form:"password" json:"password" binding:"required,min=8,max=40"`
}

// setSession 设置session
func (service *Login) setSession(c *gin.Context, user Model.User) {
	s := sessions.Default(c)
	s.Clear()
	s.Set("user_uid", user.UID)
	s.Set("data_type","redis")
	s.Save()
}

func (service *Login) ULogin(c *gin.Context) Respond.Response {
	var user Model.User
	if err := Model.DB.Where("uid = ?", service.UID).First(&user).Error; err != nil {
		return Respond.ParamError("账号或密码错误", nil)
	}

	if user.CheckPassword(service.PassWord) == false {
		return Respond.ParamError("账号或密码错误", nil)
	}

	// 设置session
	service.setSession(c, user)
	//将当前用户放入缓存
	RedisServe.AddUser(user)
	return Respond.BuildUserResponse(user)
}
