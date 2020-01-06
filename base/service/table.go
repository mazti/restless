package service

import (
	"context"
	"github.com/mazti/restless/base/dto"
	"github.com/mazti/restless/base/repository"
)

type TableService interface {
	Create(ctx context.Context, req dto.CreateTableReq) error
}

func NewTableService(baseRepo repository.BaseRepository) (TableService, error) {
	return &tableService{
		repo:     baseRepo,
	}, nil
}

type tableService struct {
	repo     repository.BaseRepository
}

func (h tableService) Create(ctx context.Context, req dto.CreateTableReq) error {
	//resp, err := h.metaClient.Create(ctx, &pb.CreateMetaRequest{Name: req.Name})
	//if err != nil {
	//	return err
	//}
	//if resp == nil {
	//	return errors.New("create base meta fail")
	//}
	err := h.repo.CreateTable(req)
	if err != nil {
		return err
	}
	return err
}
