package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID   string `gorm:"type:char(36);primaryKey"`
	Name string `json:"name"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = uuid.New().String()
	return
}
