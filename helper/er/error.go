package er

import (
	"github.com/gin-gonic/gin"
	"tfr-game/helper"
)

// ApiError api错误
type ApiError struct {
	code    ErrorCode
	message string
	LogMsg  string
}

// Msg 获取错误消息
func (ae ApiError) Error() string {
	if "" != ae.message {
		return ae.message
	}

	return ae.code.Msg()
}

// NeedLog 是否需要记录日志
func (ae *ApiError) NeedLog() bool {
	for _, code := range NeedLogCode {
		if code == ae.code {
			return true
		}
	}

	return false
}

// NeedLog 是否需要记录日志
func (ae *ApiError) HandleWithGin(c *gin.Context) bool {
	helper.Fail(c, int(ae.code), ae.Error())
	return ae.NeedLog()
}

// Panic 抛出一个panic
// errorCode 错误码
// message[0] 非必填，前端错误消息，缺省时 会根据 errorCode
// message[1] 非必填，用于后端排查的错误信息
func Panic(errorCode ErrorCode, message ...string) {
	ae := New(errorCode, message...)

	panic(ae)
}

// New 实例化一个Api错误
// errorCode 错误码
// message[0] 非必填，前端错误消息，缺省时 会根据 errorCode
// message[1] 非必填，用于后端排查的错误信息
func New(errorCode ErrorCode, message ...string) *ApiError {
	ae := &ApiError{message: "", code: errorCode}

	if len(message) > 0 {
		ae.message = message[0]
	}
	if len(message) > 1 {
		ae.LogMsg = message[1]
	}

	return ae
}

// Out 响应错误消息
func Out(c *gin.Context, errcode ErrorCode, errmsg ...string) {
	var errorMsg string

	if len(errmsg) > 0 {
		errorMsg = errmsg[0]
	} else {
		errorMsg = errcode.Msg()
	}

	helper.Fail(c, int(errcode), errorMsg)
}
