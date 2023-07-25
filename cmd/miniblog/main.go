package main

import (
	"log"
	"miniblog/internal/miniblog"
)

// 入口函数
func main() {
	command := miniblog.NewMiniBlogCommand()
	if err := command.Execute(); err != nil {
		log.Fatalf("miniblog 启动失败：%v", err)
	}
}
