package services

import (
	"alfath_lms/api/definitions"
	"alfath_lms/api/deps/pagination"
	"alfath_lms/api/models"

	"gorm.io/gorm"
)

type ChapterService struct {
	db        *gorm.DB
	paginator *pagination.Paginator
}

func (chapterSvc *ChapterService) Inject(db *gorm.DB, paginator *pagination.Paginator) {
	chapterSvc.db = db
	chapterSvc.paginator = paginator
}

func (chapterSvc *ChapterService) Create(chapter models.CourseChapter) (definitions.GenericCreationMessage, error) {
	result := chapterSvc.db.Create(&chapter)
	if result.Error != nil {
		return definitions.GenericCreationMessage{}, result.Error
	}
	return definitions.GenericCreationMessage{
		Status:     201,
		InstanceID: chapter.ID,
	}, nil
}

func (chapterSvc *ChapterService) Update(id int, chapter models.CourseChapter) (definitions.GenericAPIMessage, error) {
	var chapterTemp models.CourseChapter
	result := chapterSvc.db.Model(&chapterTemp).Where("id = ?", id).Updates(&chapter)
	if result.Error != nil {
		return definitions.GenericAPIMessage{}, result.Error
	}
	return definitions.GenericAPIMessage{
		Status:  200,
		Message: "Chapter is successfully updated",
	}, nil
}

func (chapterSvc *ChapterService) Delete(id int) (definitions.GenericAPIMessage, error) {
	//delete here means soft delete
	result := chapterSvc.db.Where("id = ?", id).Delete(&models.CourseChapter{})
	if result.Error != nil {
		return definitions.GenericAPIMessage{}, result.Error
	}
	return definitions.GenericAPIMessage{
		Status:  200,
		Message: "Chapter has been deleted successfully",
	}, nil
}

func (chapterSvc *ChapterService) Get(id int) (models.CourseChapter, error) {
	var chapter models.CourseChapter

	result := &chapter
	chapterSvc.db.First(result, "id = ?", id)

	return *result, nil

}
