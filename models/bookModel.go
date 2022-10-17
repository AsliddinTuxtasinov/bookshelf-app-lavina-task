package models

import "gorm.io/gorm"

type BookModel struct {
	ISBN      string  `json:"isbn"`
	Title     string  `json:"title"`
	Author    string  `json:"author"`
	Published string  `json:"published"`
	Pages     float64 `json:"pages"`
}

type Book struct {
	gorm.Model
	BookModel
	Status string `json:"status" gorm:"default:new"`
}
