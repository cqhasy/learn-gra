package data

import (
	"gorm.io/gorm"
)

type User struct {
	UserName string `gorm:"column:username;unique"`
	Password string `gorm:"column:password"`
	gorm.Model
}
type Card struct {
	gorm.Model
	Title    string `gorm:"column:title"`
	Content  string `gorm:"column:content"`
	FromUser uint   `gorm:"column:from_user"`
	ToUser   uint   `gorm:"column:to_user"`
}
