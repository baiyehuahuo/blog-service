package v1

import (
	"github.com/gin-gonic/gin"
)

type Article struct{}

func NewArticle() Article {
	return Article{}
}

// Create
// @Summary 新增文章
// @Produce json
// @Param name body string true "文章名"
// @Param content body string true "文章内容"
// @Param state body int false "状态"
// @Param created_by body string true "创建者"
// @Success 200 {object} model.ArticleSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles [post]
func (a Article) Create(c *gin.Context) {}

// Delete
// @Summary 删除文章
// @Produce json
// @Param id body int true "文章id"
// @Success 200 {object} model.ArticleSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles/{id} [delete]
func (a Article) Delete(c *gin.Context) {}

// Update
// @Summary 更新文章
// @Produce json
// @Param id body int true "文章id"
// @Param name body string false "文章名"
// @Param content body string false "文章内容"
// @Param state body int false "状态"
// @Param modified_by body string true "修改者"
// @Success 200 {object} model.ArticleSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles/{id} [put]
func (a Article) Update(c *gin.Context) {}

func (a Article) Get(c *gin.Context) {}

// List
// @Summary 获取多个文章
// @Produce json
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} model.ArticleSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles [get]
func (a Article) List(c *gin.Context) {}
