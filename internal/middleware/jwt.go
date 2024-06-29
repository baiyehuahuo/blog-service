package middleware

import (
	"blog-service/pkg/app"
	"blog-service/pkg/errcode"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			token string
			eCode = errcode.Success
		)
		// 检测是否有输出的一种方式
		if s, exist := c.GetQuery("token"); exist {
			token = s
		} else {
			token = c.GetHeader("token")
		}

		// 仅仅检测 token 是否存在以及解析后是否过期
		if token == "" {
			eCode = errcode.InvalidParams
		} else {
			_, err := app.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					eCode = errcode.UnauthorizedTokenTimeout
				default:
					eCode = errcode.UnauthorizedTokenError
				}
			}
		}

		if !errors.Is(eCode, errcode.Success) {
			response := app.NewResponse(c)
			response.ToErrorResponse(eCode)
			c.Abort()
			return
		}

		c.Next()
	}
}
