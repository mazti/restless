package repository

import (
	"github.com/jinzhu/gorm"

	"github.com/tiennv147/mazti-commons/errors"
	"github.com/tiennv147/restless/model"
)

type SampleRepository interface {
	Create(resource model.Sample) (model.Sample, error)
	Get(id int64) (model.Sample, error)
	List(offset int64, limit int64) ([]model.Sample, error)
	Update(resource model.Sample) error
	Delete(id int64) error
	Count() (int64, error)
}

type channelRepository struct {
	db *gorm.DB
}

func NewSampleRepository(db *gorm.DB) SampleRepository {
	// Leave here for using when need
	//defer db.Close()
	//db.AutoMigrate(&model.Sample{}, &model.TVProgram{})
	return &channelRepository{
		db: db,
	}
}

func (repo *channelRepository) Create(resource model.Sample) (model.Sample, error) {
	if err := repo.db.Create(&resource).Error; err != nil {
		return resource, err
	}
	return resource, nil
}

func (repo *channelRepository) Update(resource model.Sample) error {
	return repo.db.Save(&resource).Error
}

func (repo *channelRepository) Get(id int64) (model.Sample, error) {
	resource := model.Sample{}
	r := repo.db.First(&resource, id)
	if r.RecordNotFound() {
		return resource, errors.ErrNotFound
	}
	return resource, r.Error
}

func (repo *channelRepository) List(offset int64, limit int64) ([]model.Sample, error) {
	var resources []model.Sample
	err := repo.db.Offset(offset).Limit(limit).Find(&resources).Error
	return resources, err
}

func (repo *channelRepository) Delete(id int64) error {
	r := repo.db.Delete(&model.Sample{ID: id})
	if r.RowsAffected == 0 {
		return errors.ErrNotFound
	}
	return r.Error
}

func (repo *channelRepository) Count() (int64, error) {
	var count int64
	err := repo.db.Model(&model.Sample{}).Count(&count)
	return count, err.Error
}
