/*
@Title : timeChange
@Description :
@Author : 谭靖渝
@Update : 2021/5/6 17:30
*/
package RedisServe

import (
	"time"
)

var store []string
var i = 0

func GetViewToSql() {
	ticker := time.Tick(2*60*time.Second)
	for  range ticker {
		store = []string{}
		i = 0
		Transfer(i,store)
	}
}