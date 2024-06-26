package model

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type Auth struct {
	*Model
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
}

func (Auth) TableName() string {
	return "blog_auth"
}

func (a Auth) Get(db *gorm.DB) (*Auth, error) {
	var auth Auth
	db = db.Where("app_key = ? AND app_secret = ? AND is_deleted = ?", a.AppKey, a.AppSecret, 0)
	err := db.First(&auth).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &auth, nil
}
