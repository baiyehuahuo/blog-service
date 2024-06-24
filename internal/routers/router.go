package routers

import (
	_ "blog-service/docs"
	v1 "blog-service/internal/routers/api/v1"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	// 注册一个针对 swagger 的路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	var (
		tag     v1.Tag
		article v1.Article
	)
	if apiv1 := r.Group("/api/v1"); apiv1 != nil {
		// 标签管理
		apiv1.POST("/tags", tag.Create)       // 新增标签
		apiv1.DELETE("/tags/:id", tag.Delete) // 删除指定标签
		apiv1.PUT("/tags/:id", tag.Update)    // 更新指定标签
		apiv1.GET("/tags", tag.List)          // 获取标签列表

		// 文章管理
		apiv1.POST("/articles", article.Create)       // 新增文章
		apiv1.DELETE("/articles/:id", article.Delete) // 删除指定文章
		apiv1.PUT("/articles/:id", article.Update)    // 更新指定文章
		apiv1.GET("/articles/:id", article.Get)       // 获取指定文章
		apiv1.GET("/articles", article.List)          // 获取文章列表
	}
	return r
}
