package users

import (
	model "eirc.app/internal/v1/structure"
)

type Table struct {
	// 使用者編號
	UserID string `gorm:"primaryKey;;column:user_id;type:serial;not null;" json:"user_id,omitempty"`
	// 使用者名稱
	UserName string `gorm:"column:username;type:VARCHAR;not null;" json:"name,omitempty"`
	//使用者密碼
	Password string `gorm:"column:password;type:VARCHAR;not null;" json:"uniform_number,omitempty"`
	// 使用者電子郵件
	Email string `gorm:"column:email;type:bool;VARCHAR;not null;" json:"email,omitempty"`
}

// Base struct is corresponding to table structure file
type Base struct {
	// 使用者編號
	UserID string `json:"user_id,omitempty"`
	// 使用者名稱
	UserName string `json:"username,omitempty"`
	// 使用者密碼
	Password string `json:"password,omitempty"`
	// 使用者電子郵件
	Email string `json:"email,omitempty"`
}

// Single return structure file
type Single struct {
	// 使用者編號
	UserID string `json:"user_id,omitempty"`
	// 使用者名稱
	UserName string `json:"username,omitempty"`
	// 使用者密碼
	Password string `json:"password,omitempty"`
	// 使用者電子郵件
	Email string `json:"email,omitempty"`
}

// Created struct is used to create
type Created struct {
	// 使用者編號
	UserName string `json:"username,omitempty" binding:"required" validate:"required"`
	// 使用者密碼
	Password string `json:"password,omitempty" binding:"required" validate:"required"`
	// 使用者電子郵件
	Email string `json:"email,omitempty" binding:"required" validate:"required"`
}

// Updated struct is used to update
type Updated struct {
	// 使用者編號
	UserID string `json:"user_id,omitempty" swaggerignore:"true"`
	// 使用者名稱
	UserName string `json:"username,omitempty"`
	// 使用者密碼
	Password string `json:"password,omitempty"`
	// 使用者電子郵件
	Email *string `json:"email,omitempty" `
}

// Field is structure file for search
type Field struct {
	// 使用者編號
	UserID string `json:"user_id,omitempty"  binding:"omitempty,serial" swaggerignore:"true"`
	// 使用者名稱
	UserName *string `json:"username,omitempty" form:"UserName"`
	// 使用者密碼
	Password *string `json:"password,omitempty" form:"Password"`
	// 使用者電子郵件
	Email *string `json:"email,omitempty"  form:"email"`
}

// Fields is the searched structure file (including pagination)
type Fields struct {
	Field
	model.InPage
}

// List is multiple return structure files
type List struct {
	Companies []*struct {
		// 使用者編號
		UserID string `json:"user_id,omitempty"`
		// 使用者名稱
		UserName string `json:"username,omitempty"`
		// 使用者密碼
		Password string `json:"password,omitempty"`
		// 使用者電子郵件
		Email *string `json:"email,omitempty" `
	} `json:"users"`
	model.OutPage
}

// TableName sets the insert table Username for this struct type
func (c *Table) TableName() string {
	return "users"
}
