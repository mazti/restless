package service

import (
	"context"
	"errors"

	_ "github.com/jinzhu/gorm/dialects/mysql" //mysql dialects
)

type RecordService interface {
	List(ctx context.Context, base, table string) ([]interface{}, error)
}

func NewRecordService() RecordService {
	return recordService{}
}

type recordService struct {
}

func (h recordService) List(ctx context.Context, base, table string) ([]interface{}, error) {

	return nil, errors.New("implement me")
}
