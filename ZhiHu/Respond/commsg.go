/*
@Title : commsg
@Description :
@Author : 谭靖渝
@Update : 2021/5/4 18:36
*/
package Respond

// Response 基础序列化器
type Response struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data,omitempty"`
	Msg   string      `json:"msg"`
	Error string      `json:"error,omitempty"`
}

const (
	//CodeUserInfo 成功获取用户信息
	CodeUserInfo = 200
	//CodeArticleInfo 成功获取文章信息
	CodeArticleInfo = 200
	// CodeCheckLogin 未登录
	CodeCheckLogin = 401
	// CodeNoRightErr 未授权访问
	CodeNoRightErr = 403
	// CodeDBError 数据库操作失败
	CodeDBError = 50001
	// CodeEncryptError 加密失败
	CodeEncryptError = 50002
	//CodeParamErr 各种奇奇怪怪的参数错误
	CodeParamErr = 40001
	//CodeFailedLogin登录失败
	CodeFailedLogin = 4011
	//CodeFailedRegister 注册失败
	CodeFailedRegister = 500
)
// CheckLogin 检查登录
func CheckLogin() Response {
	return Response{
		Code: CodeCheckLogin,
		Msg:  "未登录",
	}
}

// Error 通用错误处理
func Error(errCode int, msg string, err error) Response {
	res := Response{
		Code: errCode,
		Msg:  msg,
	}
	if err != nil {
		res.Error = err.Error()
	}
	return res
}


// ParamError 各种参数错误
func ParamError(msg string, err error) Response {
	if msg == "" {
		msg = "参数错误"
	}
	return Error(CodeParamErr, msg, err)
}

// DBError 数据库操作失败
func DBError(msg string, err error) Response {
	if msg == "" {
		msg = "数据库操作失败"
	}
	return Error(CodeDBError, msg, err)
}


