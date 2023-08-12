package errno

var (
	// ErrUserAlreadyExist 用户已经存在
	ErrUserAlreadyExist = &Errno{
		HTTP:    400,
		Code:    "FailedOperation.UserAlreadyExist",
		Message: "User already exist.",
	}
)
