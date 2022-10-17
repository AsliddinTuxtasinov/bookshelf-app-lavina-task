package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name   string `json:"name"`
	Key    string `json:"key" gorm:"unique"`
	Secret string `json:"secret"`
}
