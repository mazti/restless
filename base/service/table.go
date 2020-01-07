package service

import (
	"context"
	"github.com/mazti/restless/base/dto"
	"github.com/mazti/restless/base/repository"
)

type TableService interface {
	Create(ctx context.Context, req dto.CreateTableReq) (*dto.TableResp, error)
}

func NewTableService(baseRepo repository.BaseRepository, tableRepo repository.MetaTableRepository, idService IDService) (TableService, error) {
	return &tableService{
		baseRepo:  baseRepo,
		tableRepo: tableRepo,
		idService: idService,
	}, nil
}

type tableService struct {
	baseRepo  repository.BaseRepository
	tableRepo repository.MetaTableRepository
	idService IDService
}

func (h tableService) Create(ctx context.Context, req dto.CreateTableReq) (*dto.TableResp, error) {
	schemaId, table, err := req.ToTable(h.idService.DecodeID)
	if err != nil {
		return nil, err
	}
	resp, err := h.tableRepo.Create(ctx, schemaId, table)
	if err != nil {
		return nil, err
	}
	tableName, _ := h.idService.EncodeID(resp.ID)
	err = h.baseRepo.CreateTable(req.Schema, tableName, req.Columns)
	if err != nil {
		return nil, err
	}
	return dto.NewTableResp(schemaId, table, h.idService.EncodeID)
}
