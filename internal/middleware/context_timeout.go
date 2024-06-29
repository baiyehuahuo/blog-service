package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"time"
)

func ContextTimeout(t time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), t)
		defer cancel()                         // 正常关闭计时器
		c.Request = c.Request.WithContext(ctx) // 将 timeout 上下文信息不断传递下去
		c.Next()
	}
}
