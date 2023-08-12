package v1

// CreateUserRequest 定义了 `POST /v1/users` 接口的请求参数
type CreateUserRequest struct {
	Username string `json:"username" valid:"alphanum,required,stringlength(1|255)"` // valid tag 为自定义的参数校验
	Password string `json:"password" valid:"required,stringlength(6|18)"`
	Nickname string `json:"nickname" valid:"required,stringlength(1|255)"`
	Email    string `json:"email" valid:"required,email"`
	Phone    string `json:"phone" valid:"required,stringlength(11|11)"`
}
