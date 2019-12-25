package service

import (
	"context"

	"github.com/go-kit/kit/log"

	"github.com/tiennv147/restless/record/dto"
	"github.com/tiennv147/restless/record/repository"
)

type SchemaService interface {
	Create(ctx context.Context, req dto.CreateSchemaReq) error
}

func NewSchemaService(repo repository.SchemaRepository, logger log.Logger) SchemaService {
	return schemaService{
		repo:   repo,
		logger: logger,
	}
}

type schemaService struct {
	repo   repository.SchemaRepository
	logger log.Logger
}

func (h schemaService) Create(ctx context.Context, req dto.CreateSchemaReq) error {
	return h.repo.Create(req)
}
