package helper

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Success 响应数据
func Success(c *gin.Context, content interface{}, message ...string) {
	var msg string
	if len(message) > 0 {
		msg = message[0]
	} else {
		msg = "ok"
	}

	c.JSON(http.StatusOK, outData(0, msg, content))
}

// Msg 响应消息
func Msg(c *gin.Context, message ...string) {
	Success(c, nil, message...)
}

// Fail 响应错误消息
func Fail(c *gin.Context, errcode int, errmsg string) {
	c.JSON(http.StatusOK, outData(errcode, errmsg))
}

func outData(errcode int, errmsg string, contents ...interface{}) gin.H {
	var content interface{}
	if len(contents) > 0 {
		content = contents[0]
	} else {
		content = nil
	}

	return gin.H{"errcode": errcode, "errmsg": errmsg, "content": content}
}
