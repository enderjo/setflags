package service

import "github.com/jinzhu/gorm"

//BaseService base service
type BaseService struct {
	db *gorm.DB
}
