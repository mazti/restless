package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql" //mysql dialects
	"github.com/pkg/errors"
	"github.com/tiennv147/mazti-commons/config"
	"github.com/tiennv147/restless/base/dto"
)

type BaseRepository interface {
	Create(base dto.CreateBaseReq) error
}

type baseRepository struct {
	db *sql.DB
}

func NewBaseRepository(dbCfg *config.Database) (BaseRepository, error) {
	db, err := sql.Open("mysql", dbCfg.URL)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to open db URL")
	}
	db.SetMaxOpenConns(dbCfg.MaxActive)
	db.SetMaxIdleConns(dbCfg.MaxIdle)

	return &baseRepository{
		db: db,
	}, nil
}

func (repo *baseRepository) Create(base dto.CreateBaseReq) error {
	_, err := repo.db.Exec(fmt.Sprintf("CREATE DATABASE `%s`", base.Name))
	return err
}
