package service

import (
	"github.com/PioneerDev/setflags/model"
	"github.com/PioneerDev/setflags/model/database"
)

//AssetService AssetService
type AssetService struct {
	BaseService
}

//NewAssetService new an Asset service
func NewAssetService() *AssetService {
	return &AssetService{BaseService: BaseService{db: database.ServerDb}}
}

//Get get the asset by id
func (service *AssetService) Get(id string) model.Asset {
	var asset model.Asset
	service.db.First(&asset).Where("id=?", id)
	return asset
}
