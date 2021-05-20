/*
@Title : usermsg
@Description :
@Author : 谭靖渝
@Update : 2021/5/5 9:49
*/
package Respond

import (
	"ZhiHu/Model"
)

// User 用户序列化器
type User struct {
	UID string `json:"uid"`
	NickName       string `json:"nickname"`
	Status         string `json:"status"`
	Avatar         string `json:"avatar"`
}

// BuildUser 序列化用户
func buildUser(user Model.User) User {
	return User{
		UID:      user.UID,
		NickName: user.NickName,
		Status:   user.Status,
		Avatar:   user.Avatar,
	}
}

// BuildUserResponse 序列化用户响应
func BuildUserResponse(user Model.User) Response {
	return Response{
		Code: CodeUserInfo,
		Data: buildUser(user),
		Msg: "用户信息",
	}
}

