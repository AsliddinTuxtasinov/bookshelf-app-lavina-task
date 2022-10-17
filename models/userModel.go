package models

type User struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	Name   string `json:"name"`
	Key    string `json:"key" gorm:"unique"`
	Secret string `json:"secret"`
}
