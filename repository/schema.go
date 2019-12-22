package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql" //mysql dialects
	"github.com/pkg/errors"
	"github.com/tiennv147/mazti-commons/config"
	"github.com/tiennv147/restless/dto"
)

type SchemaRepository interface {
	Create(schema dto.CreateSchemaReq) error
}

type schemaRepository struct {
	db *sql.DB
}

func NewSchemaRepository(dbCfg *config.Database) (SchemaRepository, error) {
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

func (repo *schemaRepository) Create(schema dto.CreateSchemaReq) error {
	_, err := repo.db.Exec(fmt.Sprintf("CREATE DATABASE %s", schema.Name))
	return err
}
