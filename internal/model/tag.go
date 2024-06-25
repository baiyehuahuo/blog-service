package model

import (
	"blog-service/pkg/app"
	"github.com/jinzhu/gorm"
)

type Tag struct {
	*Model
	Name  string `gorm:"column:name"`            // 标签名
	State uint8  `gorm:"column:state;default:0"` // 状态 0 启用， 1 禁用
}

type TagSwagger struct {
	List  []*Tag     `json:"list"`
	Paper *app.Paper `json:"paper"`
}

func (Tag) TableName() string {
	return "blog_tag"
}

func (t Tag) Count(db *gorm.DB) (int, error) {
	var count int
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	db = db.Where("state = ?", t.State)
	if err := db.Model(&t).Where("is_deleted = ?", 0).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (t Tag) List(db *gorm.DB, pageOffset, pageSize int) ([]*Tag, error) {
	var tags []*Tag
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		// Offset 偏移量 用于指定返回记录前要跳过的记录数
		// Limit 限制检索的记录数
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	db = db.Where("state = ?", t.State)
	if err = db.Where("is_deleted = ?", 0).Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil
}

func (t Tag) Create(db *gorm.DB) error {
	return db.Create(&t).Error
}

func (t Tag) Update(db *gorm.DB, values interface{}) error {
	return db.Model(&t).Where("id = ? AND is_deleted = ?", t.ID, 0).Updates(values).Error
}

func (t Tag) Delete(db *gorm.DB) error {
	return db.Model(&t).Where("id = ? AND is_deleted = ?", t.ID, 0).Delete(t).Error
}
