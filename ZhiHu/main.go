/*
@Title : main
@Description :
@Author : 谭靖渝
@Update : 2021/5/5 9:45
*/
package main

import (
	"ZhiHu/Model"
	"ZhiHu/Serve"
	"ZhiHu/cache"
	"ZhiHu/cache/RedisServe"
)

func main() {
	//初始化数据库
	Model.Init()
	r:=Serve.NewRouter()
	go RedisServe.GetViewToSql()
	//运行
	r.Run("127.0.0.1:9090")
	defer cache.Conn.Close()
}