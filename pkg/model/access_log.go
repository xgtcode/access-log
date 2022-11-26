package model

import "gorm.io/gorm"

type AccessLog struct {
	gorm.Model
	User    string  `gorm:"column:user" json:"user"`
	Count	int		`gorm:"column:count" json:"count"`
}

