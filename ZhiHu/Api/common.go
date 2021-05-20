/*
@Title : common
@Description :
@Author : 谭靖渝
@Update : 2021/5/5 16:19
*/
package Api

import (
	"ZhiHu/Respond"
	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context)  {
	c.JSON(200, Respond.Response{
		Msg:  "Pong",
	})
}
