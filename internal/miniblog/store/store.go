package store

import (
	"gorm.io/gorm"
	"sync"
)

var (
	once      sync.Once
	DataStore *Datastore
)

// IStore 定义了 store 层所需要实现的方法
type IStore interface {
	Users() UserStore
}

// Datastore 是 IStore 的一个具体实现
type Datastore struct {
	db *gorm.DB
}

// 确保 Datastore 实现了 IStore 接口
var _ IStore = (*Datastore)(nil)

// NewStore 创建一个数据库实例
func NewStore(db *gorm.DB) *Datastore {
	// 确保 S 只被初始化一次
	once.Do(func() {
		DataStore = &Datastore{db: db}
	})

	return DataStore
}

func (ds *Datastore) Users() UserStore {
	return newUsers(ds.db)
}
