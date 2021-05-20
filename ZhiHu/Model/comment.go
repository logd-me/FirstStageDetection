/*
@Title : comment
@Description :
@Author : 谭靖渝
@Update : 2021/5/13 15:27
*/
package Model

import (
	"fmt"
	"time"
)

type Comment struct {
	CreatedAt time.Time `gorm:"comment:'创建时间';type:datetime;"`
	UpdatedAt time.Time `gorm:"comment:'修改时间';type:datetime;"`
	CCID int `gorm:"cc_id;primary_key;AUTO_INCREMENT"`
	CID  int `gorm:"c_id"`
	Content string `gorm:"content"`
	UserID string `gorm:"user_id"`
	UserNickname string `gorm:"user_nickname"`
	UserAvatar string `gorm:"user_avatar"`
}

//获取一个文章的全部评论


//删除用户的一条评论
func DelComment(UID interface{},CCID interface{}) error {
	err:=DB.Where("user_id = ? AND cc_id = ?",UID,CCID).Unscoped().Delete(&Comment{}).Error
	return err
}

//获取一篇文章的全部评论
func GetComment(CID interface{})[]Comment  {
	var allComment []Comment
	result:=DB.Model(&Comment{}).Where("c_id = ?",CID).Find(&allComment)
	fmt.Println(result.Error)
	return allComment
}

