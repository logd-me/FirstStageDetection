/*
@Title : register
@Description :
@Author : 谭靖渝
@Update : 2021/5/5 11:46
*/
package User

import (
	"ZhiHu/Model"
	"ZhiHu/Respond"
)

const length = 8

//服务注册结构体
type Register struct {
	UID        string `form:"uid" json:"uid"`
	NickName        string `form:"nickname" json:"nickname" binding:"required,min=2,max=30"`
	PassWord        string `form:"password" json:"password" binding:"required,min=8,max=40,omitempty"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm" binding:"required,min=8,max=40,omitempty"`
}

// check 验证注册输入的东西是否有效
func (service *Register) check() *Respond.Response {
	if service.PasswordConfirm != service.PassWord {
		return &Respond.Response{
			Code: 40001,
			Msg:  "两次输入的密码不相同",
		}
	}

	count := 0
	Model.DB.Model(&Model.User{}).Where("nickname = ?", service.NickName).Count(&count)
	if count > 0 {
		return &Respond.Response{
			Code: 40001,
			Msg:  "昵称被占用",
		}
	}
	return nil
}

// Register 用户注册
func (service *Register) Register() Respond.Response {
	user := Model.User{
		UID: Model.Random(length),
		NickName: service.NickName,
		Status:   Model.Active,
	}

	// 表单验证
	if err := service.check(); err != nil {
		return *err
	}

	// 加密密码
	if err := user.SetPassword(service.PassWord); err != nil {
		return Respond.Error(
			Respond.CodeEncryptError,
			"密码加密失败",
			err,
		)
	}

	// 创建用户
	if err := Model.DB.Create(&user).Error; err != nil {
		return Respond.DBError("注册失败", err)
	}

	return Respond.BuildUserResponse(user)
}
