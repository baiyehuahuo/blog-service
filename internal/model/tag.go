package model

type Tag struct {
	*Model
	Name  string `gorm:"column:name"`            // 标签名
	State uint8  `gorm:"column:state;default:0"` // 状态 0 启用， 1 禁用
}

func (Tag) TableName() string {
	return "blob_tag"
}
