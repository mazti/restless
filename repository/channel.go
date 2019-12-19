package repository

import (
	"github.com/jinzhu/gorm"

	"github.com/tiennv147/restless/commons/errors"
	"github.com/tiennv147/restless/model"
)

type ChannelRepository interface {
	Create(resource model.Channel) (model.Channel, error)
	Get(id int64) (model.Channel, error)
	List(offset int64, limit int) ([]model.Channel, error)
	Update(resource model.Channel) error
	Delete(id int64) error
	Count() (int64, error)
}

type channelRepository struct {
	db *gorm.DB
}

func NewChannelRepository(db *gorm.DB) ChannelRepository {
	// Leave here for using when need
	//defer db.Close()
	//db.AutoMigrate(&model.Channel{}, &model.TVProgram{})
	return &channelRepository{
		db: db,
	}
}

func (repo *channelRepository) Create(resource model.Channel) (model.Channel, error) {
	if err := repo.db.Create(&resource).Error; err != nil {
		return resource, err
	}
	return resource, nil
}

func (repo *channelRepository) Update(resource model.Channel) error {
	return repo.db.Save(&resource).Error
}

func (repo *channelRepository) Get(id int64) (model.Channel, error) {
	resource := model.Channel{}
	r := repo.db.First(&resource, id)
	if r.RecordNotFound() {
		return resource, errors.ErrNotFound
	}
	return resource, r.Error
}

func (repo *channelRepository) List(offset int64, limit int) ([]model.Channel, error) {
	var resources []model.Channel
	err := repo.db.Offset(offset).Limit(limit).Find(&resources).Error
	return resources, err
}

func (repo *channelRepository) Delete(id int64) error {
	r := repo.db.Delete(&model.Channel{ID: id})
	if r.RowsAffected == 0 {
		return errors.ErrNotFound
	}
	return r.Error
}

func (repo *channelRepository) Count() (int64, error) {
	var count int64
	err := repo.db.Model(&model.Channel{}).Count(&count)
	return count, err.Error
}
