package model

import "blog-service/pkg/app"

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
	return "blob_tag"
}
