package errno

var (
	OK = &JsonResult{Code: 0, ErrorMsg: "OK"}

	ErrNotAuthorized = &JsonResult{Code: 1000, ErrorMsg: "访问未授权"}
	ErrInvalidParam  = &JsonResult{Code: 1001, ErrorMsg: "参数格式有误"}
	ErrSystemError   = &JsonResult{Code: 1002, ErrorMsg: "系统错误"}
	ErrAuthTimeout   = &JsonResult{Code: 1003, ErrorMsg: "登录超时"}
	ErrAuthTokenErr  = &JsonResult{Code: 1004, ErrorMsg: "Token错误"}
	ErrUserErr       = &JsonResult{Code: 1005, ErrorMsg: "用户更新错误"}
)
