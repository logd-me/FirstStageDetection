/*
@Title : init
@Description :
@Author : 谭靖渝
@Update : 2021/5/4 17:55
*/
package Model

import (
	"ZhiHu/cache"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func Init() (err error) {
	dsn := "root:root@(127.0.0.1:3306)/ZhiHu?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	DB.LogMode(true)
	// Error
	if err != nil {
	return err
	}
	//自动转移模板
	migration()

	//设置外键
	if err=DB.Model(&Article{}).AddForeignKey("user_id","users(uid)","CASCADE","CASCADE").Error;err!=nil{
		return err
	}
	if err=DB.Model(&Comment{}).AddForeignKey("c_id","articles(c_id)","CASCADE","CASCADE").Error;err!=nil{
		return err
	}
	//初始化redis
	cache.RedisInit()
	return nil
}
