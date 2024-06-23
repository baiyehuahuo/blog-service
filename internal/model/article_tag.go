package model

type ArticleTag struct {
	*Model
	ArticleID uint32 `gorm:"column:article_id;NOT NULL"` // 文章 ID
	TagID     uint32 `gorm:"column:tag_id;NOT NULL"`     // 标签 ID
}
