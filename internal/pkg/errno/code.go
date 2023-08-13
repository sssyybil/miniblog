package errno

var (
	// OK 请求成功
	OK = &Errno{
		HTTP:    200,
		Code:    "",
		Message: "",
	}

	// InternalServerError 所有未知的服务端错误
	InternalServerError = &Errno{
		HTTP:    500,
		Code:    "InternalError",
		Message: "Internal server error.",
	}

	// ErrPageNotFound 路由不匹配错误
	ErrPageNotFound = &Errno{
		HTTP:    404,
		Code:    "ResourceNotFound.PageNotFound",
		Message: "Page not found.",
	}

	// ErrBind 参数绑定错误
	ErrBind = &Errno{
		HTTP:    400,
		Code:    "InvalidParameter.BindError",
		Message: "Error occurred while binding the request body to the struct.",
	}

	// ErrInvalidParam 表示所有验证失败的错误
	ErrInvalidParam = &Errno{
		HTTP:    400,
		Code:    "InvalidParameter",
		Message: "Parameter verification failed.",
	}
)
