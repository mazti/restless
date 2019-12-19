package repository

import (
	"github.com/jinzhu/gorm"

	"github.com/tiennv147/restless/commons/errors"
	"github.com/tiennv147/restless/model"
)

type TVProgramRepository interface {
	Create(resource model.TVProgram) (model.TVProgram, error)
	Get(id int64) (model.TVProgram, error)
	List(offset int64, limit int) ([]model.TVProgram, error)
	Update(resource model.TVProgram) error
	Delete(id int64) error
	Count() (int64, error)
}

type tvProgramRepository struct {
	db *gorm.DB
}

func NewTVProgramRepository(db *gorm.DB) TVProgramRepository {
	return &tvProgramRepository{
		db: db,
	}
}

func (repo *tvProgramRepository) Create(resource model.TVProgram) (model.TVProgram, error) {
	if err := repo.db.Create(&resource).Error; err != nil {
		return resource, err
	}
	return resource, nil
}

func (repo *tvProgramRepository) Update(resource model.TVProgram) error {
	return repo.db.Save(&resource).Error
}

func (repo *tvProgramRepository) Get(id int64) (model.TVProgram, error) {
	resource := model.TVProgram{}
	r := repo.db.First(&resource, id)
	if r.RecordNotFound() {
		return resource, errors.ErrNotFound
	}
	return resource, r.Error
}

func (repo *tvProgramRepository) List(offset int64, limit int) ([]model.TVProgram, error) {
	var resources []model.TVProgram
	err := repo.db.Offset(offset).Limit(limit).Find(&resources).Error
	return resources, err
}

func (repo *tvProgramRepository) Delete(id int64) error {
	r := repo.db.Delete(&model.TVProgram{ID: id})
	if r.RowsAffected == 0 {
		return errors.ErrNotFound
	}
	return r.Error
}

func (repo *tvProgramRepository) Count() (int64, error) {
	var count int64
	err := repo.db.Model(&model.TVProgram{}).Count(&count)
	return count, err.Error
}
