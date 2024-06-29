package dao

import "blog-service/internal/model"

// GetAuth
// @Summary 获取 JWT 验证字符串
// @Produce json
// @Param app_key formData string true "用户名"
// @Param app_secret formData string true "密码"
// @Success 200 {object} map[string]interface{} "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /auth [post]
func (d *Dao) GetAuth(appKey, appSecret string) (*model.Auth, error) {
	auth := model.Auth{AppKey: appKey, AppSecret: appSecret}
	return auth.Get(d.engine)
}
