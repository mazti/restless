package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql" //mysql dialects
	"github.com/tiennv147/restless/base/dto"
)

type BaseRepository interface {
	CreateSchema(name string) error
	CreateTable(base dto.CreateTableReq) error
}

type baseRepository struct {
	db *sql.DB
}

func NewBaseRepository(db *sql.DB) BaseRepository {
	return &baseRepository{
		db: db,
	}
}

func (repo *baseRepository) CreateSchema(name string) error {
	_, err := repo.db.Exec(fmt.Sprintf("CREATE DATABASE `%s`", name))
	return err
}

func (repo *baseRepository) CreateTable(req dto.CreateTableReq) error {
	_, err := repo.db.Exec(req.ToQuery())
	return err
}
