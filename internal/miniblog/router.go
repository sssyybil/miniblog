package miniblog

import (
	"github.com/gin-gonic/gin"
	"miniblog/internal/miniblog/controller/v1/user"
	"miniblog/internal/miniblog/store"
	"miniblog/internal/pkg/core"
	"miniblog/internal/pkg/errno"
	"miniblog/internal/pkg/log"
)

func installRouters(engine *gin.Engine) error {
	// 注册 404 Handler
	engine.NoRoute(func(ctx *gin.Context) {
		core.WriteResponse(ctx, errno.ErrPageNotFound, nil)
	})

	// 注册 /health Handler
	engine.GET("/health", func(ctx *gin.Context) {
		log.C(ctx).Infow("Health function called")
		core.WriteResponse(ctx, nil, gin.H{"status": "OK"})
	})

	userController := user.New(store.DataStore)

	// 创建 v1 路由分组
	v1 := engine.Group("/v1")
	{
		// 创建 users 路由分组
		usersV1 := v1.Group("/users")
		{
			usersV1.POST("", userController.Create)
		}
	}
	return nil
}
