package app

import (
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"strings"
)

type ValidError struct {
	Key     string
	Message string
}

type ValidErrors []*ValidError

func (v *ValidError) Error() string {
	return v.Message
}

func (v ValidErrors) Error() string {
	return strings.Join(v.Errors(), ",")
}

func (v ValidErrors) Errors() []string {
	var errors []string
	for _, err := range v {
		errors = append(errors, err.Error())
	}
	return errors
}

// BindAndValid 针对入参校验做二次封装，主要是获取中间件的错误信息来进行翻译
func BindAndValid(c *gin.Context, v interface{}) (bool, ValidErrors) {
	var errs ValidErrors
	err := c.ShouldBind(v)
	if err == nil {
		return true, nil
	}
	v = c.Value("trans")
	trans, _ := v.(ut.Translator)
	verrs, ok := err.(validator.ValidationErrors)
	if !ok {
		return false, nil
	}
	for key, val := range verrs.Translate(trans) {
		errs = append(errs, &ValidError{Key: key, Message: val})
	}
	return false, errs
}
