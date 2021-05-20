/*
@Title : Change
@Description :
@Author : 谭靖渝
@Update : 2021/5/8 19:49
*/
package User

import (
	"ZhiHu/Model"
	"ZhiHu/Respond"
)

//修改信息结构体
type Change struct {
	UID             string `form:"uid" json:"uid" binding:"required,min=5,max=30"`
	PassWord        string `form:"password" json:"password" binding:"required,min=8,max=40"`
	Avatar          string `form:"avatar" json:"avatar"`
	NickName        string `form:"nickname" json:"nickname"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm" binding:"required,min=8,max=40"`
}

//修改密码
func (change *Change) ChangePassword()(Respond.Response,bool) {
	var user Model.User
	if err := Model.DB.Where("uid = ?", change.UID).First(&user).Error;err!=nil{
		return Respond.Response{
			Code: 0,
			Msg: err.Error(),
		},false
	}
	if err := change.checkP(); err != nil {
		return *err,false
	}
	if err := user.SetPassword(change.PassWord); err != nil {
		return Respond.Error(
			Respond.CodeEncryptError,
			"密码加密失败",
			err,
		),false
	}
	if err := Model.DB.Save(&user).Error; err != nil {
		return Respond.DBError("修改密码失败", err),false
	}
	return Respond.BuildUserResponse(user),true
}
//修改昵称
func (change *Change)ChangeNickName()Respond.Response {
	var user Model.User
	if err := Model.DB.Where("uid = ?", change.UID).First(&user).Error; err != nil {
		return Respond.Response{
			Code: 0,
			Msg:  err.Error(),
		}
	}
	if err := change.checkN(); err != nil {
		return *err
	}
	user.NickName = change.NickName
	if err := Model.DB.Save(&user).Error; err != nil {
		return Respond.DBError("修改昵称失败", err)
	}
	return Respond.BuildUserResponse(user)
}
func (change *Change) ChangeAvatar()Respond.Response  {
	var user Model.User
	if err := Model.DB.Where("uid = ?", change.UID).First(&user).Error; err != nil {
		return Respond.Response{
			Code: 0,
			Msg:  err.Error(),
		}
	}
	user.Avatar = change.Avatar
	if err := Model.DB.Save(&user).Error; err != nil {
		return Respond.DBError("修改头像失败", err)
	}
	return Respond.BuildUserResponse(user)
}

//验证密码
func (change *Change) checkP() *Respond.Response {
	if change.PasswordConfirm != change.PassWord {
		return &Respond.Response{
			Code: 40001,
			Msg:  "两次输入的密码不相同",
		}
	}
	return nil
}
//验证昵称
func (change *Change) checkN() *Respond.Response {
	count := 0
	Model.DB.Model(&Model.User{}).Where("nickname = ?", change.NickName).Count(&count)
	if count > 0 {
		return &Respond.Response{
			Code: 40001,
			Msg:  "昵称被占用",
		}
	}
	return nil
}
