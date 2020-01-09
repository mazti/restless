package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql" //mysql dialects
	"github.com/mazti/restless/base/dto"
	"github.com/mazti/restless/base/ent"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
)

type BaseRepository interface {
	CreateSchema(name string) error
	CreateTable(schema string, table string, columns []dto.Column) error
	CreateColumn(schema string, table string, columnName string, options string) error
	UpdateColumn(schema string, table string, columnName string, options string) error
	DeleteColumn(schema string, table string, columnName string) error
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

func (repo *baseRepository) CreateTable(schema string, table string, columns []dto.Column) error {
	_, err := repo.db.Exec(queryCreatTable(schema, table, columns))
	return err
}

func (repo *baseRepository) CreateColumn(schema string, table string, columnName string, options string) error {
	_, err := repo.db.Exec(queryCreatColumn(schema, table, columnName, options))
	return err
}

func (repo *baseRepository) UpdateColumn(schema string, table string, columnName string, options string) error {
	_, err := repo.db.Exec(queryUpdateColumn(schema, table, columnName, options))
	return err
}
func (repo *baseRepository) DeleteColumn(schema string, table string, columnName string) error {
	_, err := repo.db.Exec(queryDeleteColumn(schema, table, columnName))
	return err
}

func readError(err error) error {
	if err == nil {
		return err
	} else if ent.IsNotFound(err) {
		return status.Error(codes.NotFound, err.Error())
	}
	return err
}

func queryCreatTable(schema string, table string, columns []dto.Column) string {
	var query []string
	query = append(query, "CREATE TABLE")
	query = append(query, fmt.Sprintf("`%s`.`%s`", schema, table))
	query = append(query, "(")

	var cols []string
	for _, c := range columns {
		cols = append(cols, c.ToQuery())
	}
	query = append(query, strings.Join(cols, ","))
	query = append(query, ")")

	return strings.Join(query, " ")
}

func queryCreatColumn(schema string, table string, columnName string, options string) string {
	return columnQuery("ADD", schema, table, columnName, options)
}

func queryUpdateColumn(schema string, table string, columnName string, options string) string {
	return columnQuery("MODIFY", schema, table, columnName, options)
}

func queryDeleteColumn(schema string, table string, columnName string) string {
	return columnQuery("DROP", schema, table, columnName, "")
}

func columnQuery(action string, schema string, table string, columnName string, options string) string {
	var query []string
	query = append(query, "ALTER TABLE")
	query = append(query, fmt.Sprintf("`%s`.`%s`", schema, table))
	query = append(query, action)
	query = append(query, columnName)
	query = append(query, options)
	return strings.Join(query, " ")
}
