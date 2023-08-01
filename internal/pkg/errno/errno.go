package errno

import "fmt"

// Errno 定义了 MiniBlog 使用的错误类型
type Errno struct {
	HTTP    int    // HTTP 状态码
	Code    string // 业务错误码
	Message string // 可直接暴露给用户的错误信息
}

// Error 实现了 error 接口中的 Error 方法
func (e *Errno) Error() string {
	return e.Message
}

// SetMessage 设置 Errno 错误类型中的 message
func (e *Errno) SetMessage(format string, args ...any) *Errno {
	e.Message = fmt.Sprintf(format, args...)
	return e
}

// Decode 尝试从 err 中解析中 HTTP 状态码、业务错误码和错误信息
func Decode(err error) (int, string, string) {
	if err == nil {
		return OK.HTTP, OK.Code, OK.Message
	}

	switch typed := err.(type) {
	case *Errno:
		return typed.HTTP, typed.Code, typed.Message
	default:
		return InternalServerError.HTTP, InternalServerError.Code, InternalServerError.Message
	}
}
