package golanggorm

import (
	"time"
	"gorm.io/gorm"
)

type Todo struct {
	ID int64 `gorm:"primary_key;column:id;autoIncrement"`
	UserID string `gorm:"column:user_id"`
	Title string `gorm:"column:title"`
	Description string `gorm:"column:description"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"deleted_at"`
}

func (t *Todo) TableName()string  {
	return "todos"
}