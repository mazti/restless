package service

import (
	"context"
	_ "github.com/jinzhu/gorm/dialects/mysql" //mysql dialects
)

type RecordService interface {
	List(ctx context.Context, base, table string, maxRecords, pageSize int32) ([]interface{}, error)
	Get(ctx context.Context, base, table string, recordID string) (interface{}, error)
	CreateRecords(ctx context.Context, base, table string, records []interface{}) ([]interface{}, error)
	UpdateRecords(ctx context.Context, base, table string, records []interface{}) ([]interface{}, error)
	Delete(ctx context.Context, base, table string, recordID string) error
}

func NewRecordService() RecordService {
	return recordService{}
}

type recordService struct {
}

func (h recordService) List(ctx context.Context, base, table string, maxRecords, pageSize int32) ([]interface{}, error) {
	panic("implement me")
}

func (h recordService) Get(ctx context.Context, base, table string, recordID string) (interface{}, error) {
	panic("implement me")
}

func (h recordService) CreateRecords(ctx context.Context, base, table string, records []interface{}) ([]interface{}, error) {
	panic("implement me")
}

func (h recordService) UpdateRecords(ctx context.Context, base, table string, records []interface{}) ([]interface{}, error) {
	panic("implement me")
}

func (h recordService) Delete(ctx context.Context, base, table string, recordID string) error {
	panic("implement me")
}
