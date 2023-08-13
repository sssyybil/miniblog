package model

import (
	"gorm.io/gorm"
	"miniblog/pkg/auth"
	"time"
)

// UserM 存储用户信息
// 结构体命名规范：表名首字母大写➕M（Model）
type UserM struct {
	ID        int64     `gorm:"column:id;primary_key"`
	Username  string    `gorm:"column:username;not null"`
	Password  string    `gorm:"column:password;not null"`
	Nickname  string    `gorm:"column:nickname"`
	Email     string    `gorm:"column:email"`
	Phone     string    `gorm:"column:phone"`
	CreatedAt time.Time `gorm:"column:createdAt"`
	UpdatedAt time.Time `gorm:"column:updatedAt"`
}

// TableName 指定映射的 MySQL 表名
func (u *UserM) TableName() string {
	return "user"
}

// BeforeCreate 加密明文密钥
func (u *UserM) BeforeCreate(db *gorm.DB) (err error) {
	if u.Password, err = auth.Encrypt(u.Password); err != nil {
		return err
	}
	return nil
}
