package store

import (
	"context"
	"gorm.io/gorm"
	"miniblog/internal/pkg/model"
)

type UserStore interface {
	Create(ctx context.Context, user *model.UserM) error
}

type users struct {
	db *gorm.DB
}

// 确保 users 实现了 UserStore 接口
var _ UserStore = (*users)(nil)

func newUsers(db *gorm.DB) *users {
	return &users{db: db}
}

// Create 插入一条 User 记录
func (u *users) Create(ctx context.Context, user *model.UserM) error {
	return u.db.Create(&user).Error
}
