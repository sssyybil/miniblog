package user

import (
	"context"
	"github.com/jinzhu/copier"
	"miniblog/internal/miniblog/store"
	"miniblog/internal/pkg/errno"
	"miniblog/internal/pkg/log"
	"miniblog/internal/pkg/model"
	v1 "miniblog/pkg/api/miniblog/v1"
	"regexp"
)

// UserBiz 定义了 user 模块在 biz 层所实现的方法
type UserBiz interface {
	Create(ctx context.Context, req *v1.CreateUserRequest) error
}

type UserBusiness struct {
	ds store.IStore
}

// 确保 UserBusiness 实现了 UserBiz 接口
var _ UserBiz = (*UserBusiness)(nil)

func New(ds store.IStore) *UserBusiness {
	return &UserBusiness{ds: ds}
}

func (b *UserBusiness) Create(ctx context.Context, req *v1.CreateUserRequest) error {
	var userModel model.UserM
	err := copier.Copy(&userModel, req)
	if err != nil {
		log.Errorw("copy CreateUserRequest to UserM fail", "err", err)
	}

	if err := b.ds.Users().Create(ctx, &userModel); err != nil {
		if match, _ := regexp.MatchString("Duplicate entry '.*' for key 'username'", err.Error()); match {
			return errno.ErrUserAlreadyExist
		}
	}
}
