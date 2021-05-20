/*
@Title : redis
@Description :
@Author : 谭靖渝
@Update : 2021/5/11 17:06
*/
package cache

import (
	"github.com/garyburd/redigo/redis"
	"github.com/sirupsen/logrus"
)

var Conn redis.Conn

//存储用户信息用db0,点赞，差评，热度，放在db1
func RedisInit() {
	var err error
	Conn, err = redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		logrus.Panic("连接失败")
		return
	}
}
