package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name   string `json:"name"`
	Key    string `json:"key"`
	Secret string `json:"secret"`
}
