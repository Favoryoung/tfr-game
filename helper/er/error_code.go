package er

// ErrorCode 错误码
type ErrorCode int

const (
	ApiErrorNotFoundError ErrorCode = 404
	Param                 ErrorCode = 10000
	User                  ErrorCode = 20000
	//Sql   sql异常
	Sql ErrorCode = 50001
	//SqlTransaction  sql事务异常
	SqlTransaction ErrorCode = 50002
	Unknown        ErrorCode = 99999
)

// DefaultErrorMsg 默认错误消息
var DefaultErrorMsg = map[ErrorCode]string{
	ApiErrorNotFoundError: "资源不存在",
	Param:                 "参数错误",
	User:                  "用户相关信息错误",
	Sql:                   "未知错误",
	SqlTransaction:        "未知错误",
	Unknown:               "未知异常",
}

// NeedLogCode 需要记录日志的异常
var NeedLogCode = []ErrorCode{
	Sql,
	SqlTransaction,
	Unknown,
}

func (code ErrorCode) Msg() string {
	if errorMsg, ok := DefaultErrorMsg[code]; ok {
		return errorMsg
	}

	return "系统错误"
}
