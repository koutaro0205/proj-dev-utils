package service

import (
	model "github.com/koutaro0205/backend-template/app/internal/models"
	"gorm.io/gorm"
)

func GetAllUsers(db *gorm.DB) ([]*model.User, error) {
	users := []*model.User{}
	db.Find(&users)

	return users, nil
}
