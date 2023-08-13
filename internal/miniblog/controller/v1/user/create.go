package user

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"miniblog/internal/pkg/core"
	"miniblog/internal/pkg/errno"
	"miniblog/internal/pkg/log"
	v1 "miniblog/pkg/api/miniblog/v1"
)

// Create 创建一个新的用户
func (ctrl *UserController) Create(ctx *gin.Context) {
	log.C(ctx).Infow("Create user function called")

	var req v1.CreateUserRequest
	// 将上下文中携带的请求参数解析到 CreateUserRequest 结构体中
	if err := ctx.ShouldBindJSON(&req); err != nil {
		core.WriteResponse(ctx, errno.ErrBind, nil)
		return
	}

	// 参数校验。govalidator 包能够根据结构体中的 valid tag 进行校验
	if _, err := govalidator.ValidateStruct(req); err != nil {
		core.WriteResponse(ctx, errno.ErrInvalidParam.SetMessage(err.Error()), nil)
		return
	}

	if err := ctrl.b.Users().Create(ctx, &req); err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	core.WriteResponse(ctx, nil, nil)
}
