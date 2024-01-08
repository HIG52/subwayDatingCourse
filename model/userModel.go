package model

import "gorm.io/gorm"

type Member struct {
	gorm.Model
	No       int `gorm:"primarykey"`
	Id       string
	Password string
	Name     string
}

// 테이블 명 명시
func (Member) TableName() string {
	return "member"
}
