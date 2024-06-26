package v1

import (
	"blog-service/internal/service"
	"blog-service/pkg/app"
	"blog-service/pkg/convert"
	"blog-service/pkg/errcode"
	"blog-service/pkg/upload"
	"github.com/gin-gonic/gin"
)

type Upload struct{}

func NewUpload() Upload {
	return Upload{}
}

// UploadFile
// @Summary 文件上传
// @Produce json
// @Param file formData file true "上传文件" swagger:extension
// @Param type formData int true "文件类型"
// @Success 200 {object} map[string]interface{} "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /upload/file [post]
func (u Upload) UploadFile(c *gin.Context) {
	response := app.NewResponse(c)
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		response.ToErrorResponse(errcode.ErrorUploadFileFail.WithDetails(err.Error()))
		return
	}

	fileType := convert.StrTo(c.PostForm("type")).MustInt()
	if fileHeader == nil || fileType <= 0 {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}

	svc := service.New(c.Request.Context())
	fileInfo, err := svc.UploadFile(upload.FileType(fileType), file, fileHeader)
	if err != nil {
		response.ToErrorResponse(errcode.ErrorUploadFileFail.WithDetails(err.Error()))
		return
	}
	response.ToResponse(gin.H{
		"file_access_url": fileInfo.AccessURL,
	})
}
