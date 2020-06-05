package service

import (
	"github.com/PioneerDev/setflags/model"
	"github.com/PioneerDev/setflags/model/database"
)

//EvidenceService evidence service
type EvidenceService struct {
	BaseService
}

//NewEvidenceService new evidence service
func NewEvidenceService() *EvidenceService {
	return &EvidenceService{BaseService: BaseService{db: database.ServerDb}}
}

//List list all evidence
func (service *EvidenceService) List(id string) []model.Evidence {
	var evidences []model.Evidence
	service.db.Find(&evidences).Where("flagId=?", id).Order("createdAt desc")
	return evidences
}

//Save save flags evidence
func (service *EvidenceService) Save(evidence *model.Evidence) bool {
	service.db.Save(evidence)
	return true
}
