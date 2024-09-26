package models

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Email     string    `json:"email" gorm:"unique"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserInfo struct {
	ID      uint   `gorm:"primaryKey"` // 主键
	UserID  uint   `gorm:"not null"`   // 外键，关联到 User 表
	Address string `gorm:"not null"`   // 地址
	Phone   string
}
