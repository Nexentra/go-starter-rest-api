package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string `json:"user_name" gorm:"unique;not null" form:"username" binding:"required"`
	Email    string `json:"email" gorm:"unique;not null" form:"email" binding:"required"`
	Password string `json:"password" gorm:"not null" form:"password" binding:"required"`
}

type LoginUserBody struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type UserResponse struct {
    ID       uint   `json:"id"`
    UserName string `json:"user_name"`
    Email    string `json:"email"`
}