package repository

import (
	"github.com/jinzhu/gorm"

	"github.com/mazti/restless/meta/model"
	"github.com/tiennv147/mazti-commons/errors"
)

type MetaRepository interface {
	Create(meta model.Meta) (model.Meta, error)
	Get(id string) (model.Meta, error)
	List(offset int, limit int) ([]model.Meta, error)
	Update(meta model.Meta) error
	Delete(id string) error
	Count() (int, error)
}

type metaRepository struct {
	db *gorm.DB
}

func NewMetaRepository(db *gorm.DB) MetaRepository {
	// Leave here for using when need
	//defer db.Close()
	//db.AutoMigrate(&model.Meta{})
	return &metaRepository{
		db: db,
	}
}

func (repo *metaRepository) Create(meta model.Meta) (model.Meta, error) {
	if err := repo.db.Create(&meta).Error; err != nil {
		return meta, err
	}
	return meta, nil
}

func (repo *metaRepository) Update(meta model.Meta) error {
	return repo.db.Save(&meta).Error
}

func (repo *metaRepository) Get(id string) (model.Meta, error) {
	meta := model.Meta{}
	r := repo.db.Where("id = ?", id).First(&meta)
	if r.RecordNotFound() {
		return meta, errors.ErrNotFound
	}
	return meta, r.Error
}

func (repo *metaRepository) List(offset int, limit int) ([]model.Meta, error) {
	var metas []model.Meta
	err := repo.db.Offset(offset).Limit(limit).Find(&metas).Error
	return metas, err
}

func (repo *metaRepository) Delete(id string) error {
	r := repo.db.Delete(&model.Meta{ID: id})
	if r.RowsAffected == 0 {
		return errors.ErrNotFound
	}
	return r.Error
}

func (repo *metaRepository) Count() (int, error) {
	var count int
	err := repo.db.Model(&model.Meta{}).Count(&count)
	return count, err.Error
}
