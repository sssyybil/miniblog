package biz

import (
	"miniblog/internal/miniblog/biz/user"
	"miniblog/internal/miniblog/store"
)

// IBiz 定义了 Biz 层需要实现的方法
type IBiz interface {
	Users() user.UserBiz
}

// Biz 是 IBiz 的一个具体实现.
type Biz struct {
	ds store.IStore
}

// 确保 Biz 实现了 IBiz 接口
var _ IBiz = (*Biz)(nil)

// NewBiz 创建一个 IBiz 类型的实例.
func NewBiz(ds store.IStore) *Biz {
	return &Biz{ds: ds}
}

// Users 返回一个实现了 UserBiz 接口的实例.
func (b *Biz) Users() user.UserBiz {
	return user.New(b.ds)
}
