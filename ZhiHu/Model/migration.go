/*
@Title : migration
@Description :
@Author : 谭靖渝
@Update : 2021/5/4 17:58
*/
package Model

func migration() {
	DB.AutoMigrate(&User{},&Article{},&Comment{})
}
