package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

type MySqlOptions struct {
	Host                  string
	Username              string
	Password              string
	Database              string
	MaxIdleConnections    int           // 空闲连接池的最大连接数
	MaxOpenConnections    int           // 数据库的最大打开连接数
	MaxConnectionLifeTime time.Duration // 连接可重用的最长时间
	LogLevel              int
}

// DSN (Data Source Name) 返回 DSN
func (o *MySqlOptions) DSN() string {
	return fmt.Sprintf(`%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=%t&loc=%s`,
		o.Username,
		o.Password,
		o.Host,
		o.Database,
		true,
		"Local")
}

func NewMySql(opts *MySqlOptions) (*gorm.DB, error) {
	// GORM log level, 1: silent, 2:error, 3:warn, 4:info
	logLevel := logger.Silent
	if opts.LogLevel != 0 {
		logLevel = logger.LogLevel(opts.LogLevel)
	}
	// 🍑根据自定义的日志等级初始化 database session 的过程还挺曲折的～
	db, err := gorm.Open(mysql.Open(opts.DSN()), &gorm.Config{Logger: logger.Default.LogMode(logLevel)})
	if err != nil {
		return nil, err
	}

	sqlDb, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDb.SetMaxOpenConns(opts.MaxOpenConnections)
	sqlDb.SetConnMaxLifetime(opts.MaxConnectionLifeTime)
	sqlDb.SetMaxIdleConns(opts.MaxIdleConnections)

	return db, nil
}
