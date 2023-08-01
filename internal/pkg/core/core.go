package core

import (
	"github.com/gin-gonic/gin"
	"miniblog/internal/pkg/errno"
	"net/http"
)

type ErrResponse struct {
	Code    string `json:"code"`    // 业务错误码
	Message string `json:"message"` // 可直接对外展示的错误信息
}

// WriteResponse 将错误或响应数据写入 HTTP 响应主体
func WriteResponse(c *gin.Context, err error, data any) {
	if err != nil {
		httpCode, code, message := errno.Decode(err)
		c.JSON(httpCode, ErrResponse{
			Code:    code,
			Message: message,
		})
		return
	}

	c.JSON(http.StatusOK, data)
}
