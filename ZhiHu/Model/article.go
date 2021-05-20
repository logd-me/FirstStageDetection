/*
@Title : comment
@Description :
@Author : 谭靖渝
@Update : 2021/5/8 16:03
*/
package Model

import (
	"ZhiHu/cache"
	"fmt"
	"time"
)

//文章模型
type Article struct {
	CreatedAt time.Time `gorm:"comment:'创建时间';type:datetime;"`
	UpdatedAt time.Time `gorm:"comment:'修改时间';type:datetime;"`
	Title     string    `gorm:"title"`
	Context   string    `gorm:"context;type:TEXT"`
	CID       int      `gorm:"c_id;primary_key;AUTO_INCREMENT"`
	UserID    string    `gorm:"user_id"`
	Like      int      `gorm:"like"`
	ULike     int      `gorm:"u_like"`
	View      int      `gorm:"view"`
	Category  string    `gorm:"category"`
}

//获取一篇文章
func GetArticle( CID interface{}) (Article, error) {
	var article Article
	result := DB.Model(&Article{}).Where(" c_id = ?",CID).First(&article)
	return article, result.Error
}

//增加热度
//func (article *Article) ViewAdd() {
//	DB.Model(&Article{}).Where("c_id=?", article.CID).Update("view", 99)
//}

func (article *Article) ViewAdd() {
	cache.Conn.Do("SELECT",1)
	cache.Conn.Do("HMSet",article.CID,"view",article.View+1)
}

//获取某位用户的全部文章
func GetAllArticle(UID interface{})[]Article  {
	var allArticle []Article
	result:=DB.Model(&Article{}).Where("user_id = ?",UID).Find(&allArticle)
	fmt.Println(result.Error)
	return allArticle
}

//删除用户的文章

func DelArticle(UID interface{},CID interface{})error  {
	err:=DB.Where("user_id = ? AND c_id = ?",UID,CID).Unscoped().Delete(&Article{}).Error
	return err
}