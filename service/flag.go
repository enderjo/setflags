package service

import (
	"github.com/PioneerDev/setflags/model"
	"github.com/PioneerDev/setflags/model/database"
)

//FlagService flag service
type FlagService struct {
	BaseService
}

//NewFlagService new flag service
func NewFlagService() *FlagService {
	return &FlagService{BaseService: BaseService{db: database.ServerDb}}
}

//List flag list
func (service *FlagService) List() []model.Flag {
	var flags []model.Flag
	service.db.Find(&flags).Order("createdAt desc")
	return flags
}

//Save save
func (service *FlagService) Save(flag *model.Flag) bool {
	service.db.Save(flag)
	return true
}

//Update update
func (service *FlagService) Update(flag *model.Flag) bool {
	service.db.Update(flag)
	return true
}

//MyList get my flag list
func (service *FlagService) MyList(id string) []model.Flag {
	var flags []model.Flag
	service.db.Where("payerId=?", id).Find(&flags).Order("createdAt desc")
	return flags
}
