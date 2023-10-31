package services

import (
	"alfath_lms/api/definitions"
	"alfath_lms/api/deps/pagination"
	"alfath_lms/api/models"

	"gorm.io/gorm"
)

type MaterialService struct {
	db        *gorm.DB
	paginator *pagination.Paginator
}

func (materialSvc *MaterialService) Inject(db *gorm.DB, paginator *pagination.Paginator) {
	materialSvc.db = db
	materialSvc.paginator = paginator
}

func (materialSvc *MaterialService) Create(material models.ChapterMaterial) (definitions.GenericCreationMessage, error) {
	result := materialSvc.db.Create(&material)
	if result.Error != nil {
		return definitions.GenericCreationMessage{}, result.Error
	}
	return definitions.GenericCreationMessage{
		Status:     201,
		InstanceID: material.ID,
	}, nil
}

func (materialSvc *MaterialService) Update(id int, material models.ChapterMaterial) (definitions.GenericAPIMessage, error) {
	var materialTemp models.ChapterMaterial
	result := materialSvc.db.Model(&materialTemp).Where("id = ?", id).Updates(&material)
	if result.Error != nil {
		return definitions.GenericAPIMessage{}, result.Error
	}
	return definitions.GenericAPIMessage{
		Status:  200,
		Message: "material is successfully updated",
	}, nil
}

func (materialSvc *MaterialService) Delete(id int) (definitions.GenericAPIMessage, error) {
	//delete here means soft delete
	result := materialSvc.db.Where("id = ?", id).Delete(&models.ChapterMaterial{})
	if result.Error != nil {
		return definitions.GenericAPIMessage{}, result.Error
	}
	return definitions.GenericAPIMessage{
		Status:  200,
		Message: "Material has been deleted successfully",
	}, nil
}

func (materialSvc *MaterialService) Get(id int) (models.ChapterMaterial, error) {
	var material models.ChapterMaterial

	result := &material
	materialSvc.db.First(result, "id = ?", id)

	return *result, nil

}
