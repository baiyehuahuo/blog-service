package model

type Article struct {
	*Model
	Title         string `gorm:"column:title"`           // 文章标题
	Desc          string `gorm:"column:desc"`            // 文章简述
	CoverImageUrl string `gorm:"column:cover_image_url"` // 封面图片地址
	Content       string `gorm:"column:content"`         // 文章内容
	State         uint8  `gorm:"column:state;default:0"` // 状态 0 启用， 1 禁用
}

func (Article) TableName() string {
	return "blog_article"
}
