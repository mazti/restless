package service

import (
	"context"

	"github.com/go-kit/kit/log"

	"github.com/tiennv147/restless/record/dto"
	"github.com/tiennv147/restless/record/repository"
)

type RecordService interface {
	Select(ctx context.Context, req dto.SelectRecordsReq) error
}

func NewRecordService(repo repository.RecordRepository, logger log.Logger) RecordService {
	return recordService{
		repo:   repo,
		logger: logger,
	}
}

type recordService struct {
	repo   repository.RecordRepository
	logger log.Logger
}

func (h recordService) Select(ctx context.Context, req dto.SelectRecordsReq) error {
	return nil
}
