package errno

var (
	// OK 代表请求成功
	OK = &Errno{
		HTTP:    200,
		Code:    "",
		Message: "",
	}

	// InternalServerError 表示所有未知的服务端错误
	InternalServerError = &Errno{
		HTTP:    500,
		Code:    "InternalError",
		Message: "Internal server error.",
	}

	ErrPageNotFound = &Errno{
		HTTP:    404,
		Code:    "ResourceNotFound.PageNotFound",
		Message: "Page not found.",
	}
)
