package user

import (
	"miniblog/internal/miniblog/biz"
	"miniblog/internal/miniblog/store"
)

// UserController user 模块在 Controller 层的实现，用来处理用户模块的请求
type UserController struct {
	b biz.IBiz
}

func New(ds store.IStore) *UserController {
	return &UserController{biz.NewBiz(ds)}
}
