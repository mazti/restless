package repository

import (
	"database/sql"
	_ "github.com/jinzhu/gorm/dialects/mysql" //mysql dialects
	"github.com/pkg/errors"
	"github.com/tiennv147/mazti-commons/config"
	"github.com/tiennv147/restless/record/dto"
)

type RecordRepository interface {
	Select(schema dto.SelectRecordsReq) error
}

type recordRepository struct {
	db *sql.DB
}

func NewRecordRepository(dbCfg *config.Database) (RecordRepository, error) {
	db, err := sql.Open("mysql", dbCfg.URL)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to open db URL")
	}
	db.SetMaxOpenConns(dbCfg.MaxActive)
	db.SetMaxIdleConns(dbCfg.MaxIdle)

	return &schemaRepository{
		db: db,
	}, nil
}

func (repo *schemaRepository) Select(schema dto.SelectRecordsReq) error {
	return nil
}
