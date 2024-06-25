package v1

import (
	"blog-service/global"
	"blog-service/internal/service"
	"blog-service/pkg/app"
	"blog-service/pkg/errcode"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Tag struct {
}

func NewTag() Tag {
	return Tag{}
}

// Create
// @Summary 新增标签
// @Produce json
// @Param name body string true "标签名"
// @Param state body int false "状态"
// @Param created_by body string true "创建者"
// @Success 200 {object} model.TagSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags [post]
func (t Tag) Create(c *gin.Context) {
	var (
		param       = service.CreateTagRequest{}
		response    = app.NewResponse(c)
		valid, errs = app.BindAndValid(c, &param)
	)
	if !valid {
		p, q := c.PostForm("name"), c.PostForm("created_by")
		fmt.Println(p, q)
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c)
	err := svc.CreateTag(&param)
	if err != nil {
		global.Logger.Errorf("svc.CreateTag err: %v", err)
		response.ToErrorResponse(errcode.ErrorCreateTagFail)
		return
	}
	response.ToResponse(gin.H{})
	return
}

// Delete
// @Summary 删除标签
// @Produce json
// @Param id path int true "签名 id"
// @Success 200 {object} model.TagSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags/{id} [delete]
func (t Tag) Delete(c *gin.Context) {
	var (
		param       = service.DeleteTagRequest{}
		response    = app.NewResponse(c)
		valid, errs = app.BindAndValid(c, &param)
	)
	if !valid {
		global.Logger.Errorf("app.BindAndValid err: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c)
	err := svc.DeleteTag(&param)
	if err != nil {
		global.Logger.Errorf("svc.DeleteTag err: %v", err)
		response.ToErrorResponse(errcode.ErrorDeleteTagFail)
		return
	}
	response.ToResponse(gin.H{})
	return
}

// Update
// @Summary 更新标签
// @Produce json
// @Param id path int true "标签 id"
// @Param name body string false "标签名"
// @Param state body int false "状态"
// @Param modified_by body string true "修改者"
// @Success 200 {object} model.TagSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags [put]
func (t Tag) Update(c *gin.Context) {
	var (
		param       = service.UpdateTagRequest{}
		response    = app.NewResponse(c)
		valid, errs = app.BindAndValid(c, &param)
	)
	if !valid {
		global.Logger.Errorf("app.BindAndValid err: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c)
	err := svc.UpdateTag(&param)
	if err != nil {
		global.Logger.Errorf("svc.UpdateTag err: %v", err)
		response.ToErrorResponse(errcode.ErrorUpdateTagFail)
		return
	}
	response.ToResponse(gin.H{})
	return
}

func (t Tag) Get(c *gin.Context) {}

// List
// @Summary 获取多个标签
// @Produce json
// @Param name query string false "标签名"
// @Param state query int false "状态"
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} model.TagSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags [get]
func (t Tag) List(c *gin.Context) {
	var (
		param       = service.TagListRequest{}
		response    = app.NewResponse(c)
		valid, errs = app.BindAndValid(c, &param)
	)
	// 入参校验 绑定
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	// 获取标签总数
	var (
		svc            = service.New(c)
		paper          = app.Paper{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
		totalRows, err = svc.CountTag(&service.CountTagRequest{Name: param.Name, State: param.State})
	)
	if err != nil {
		global.Logger.Errorf("svc.CountTag err: %v", err)
		response.ToErrorResponse(errcode.ErrorCountTagFail)
		return
	}

	// 获取标签列表
	tags, err := svc.GetTagList(&param, &paper)
	if err != nil {
		global.Logger.Errorf("svc.GetTagList: %v", err)
		response.ToErrorResponse(errcode.ErrorGetTagListFail)
		return
	}

	// 序列化返回
	response.ToResponseList(tags, totalRows)
	return
}
