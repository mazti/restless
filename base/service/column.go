package service

import (
	"context"
	"github.com/mazti/restless/base/dto"
	"github.com/mazti/restless/base/repository"
)

type ColumnService interface {
	Create(ctx context.Context, req dto.CreateColumnReq) (*dto.ColumnResp, error)
	Update(ctx context.Context, req dto.UpdateColumnReq) (*dto.ColumnResp, error)
	Delete(ctx context.Context, id string) error
}

func NewColumnService(baseRepo repository.BaseRepository,
	colRepo repository.MetaColumnRepository,
	tableRepo repository.MetaTableRepository,
	idService IDService) (ColumnService, error) {
	return &columnService{
		baseRepo:   baseRepo,
		colRepo:    colRepo,
		tableRepo:  tableRepo,
		idService:  idService,
	}, nil
}

type columnService struct {
	baseRepo   repository.BaseRepository
	colRepo    repository.MetaColumnRepository
	tableRepo  repository.MetaTableRepository
	idService  IDService
}

func (h columnService) Create(ctx context.Context, req dto.CreateColumnReq) (*dto.ColumnResp, error) {
	tableID, column, err := req.ToMeta(h.idService.DecodeID)
	if err != nil {
		return nil, err
	}
	table, err := h.tableRepo.Get(ctx, tableID)
	if err != nil {
		return nil, err
	}
	schema, err := table.QuerySchema().Only(ctx)
	if err != nil {
		return nil, err
	}
	schemaID, err := h.idService.EncodeID(schema.ID)
	if err != nil {
		return nil, err
	}
	resp, err := h.colRepo.Create(ctx, tableID, column)
	if err != nil {
		return nil, err
	}
	colID, err := h.idService.EncodeID(resp.ID)
	if err != nil {
		return nil, err
	}
	options, err := h.columnOptions(req.Type, req.Default, req.TypeOption)
	if err != nil {
		return nil, err
	}
	err = h.baseRepo.CreateColumn(schemaID, req.Table, colID, options)
	if err != nil {
		return nil, err
	}
	return dto.NewColumnResp(resp, h.idService.EncodeID)
}

func (h columnService) Update(ctx context.Context, req dto.UpdateColumnReq) (*dto.ColumnResp, error) {
	tableID, column, err := req.ToMeta(h.idService.DecodeID)
	if err != nil {
		return nil, err
	}
	table, err := h.tableRepo.Get(ctx, tableID)
	if err != nil {
		return nil, err
	}
	schema, err := table.QuerySchema().Only(ctx)
	if err != nil {
		return nil, err
	}
	schemaID, err := h.idService.EncodeID(schema.ID)
	if err != nil {
		return nil, err
	}
	options, err := h.columnOptions(req.Type, req.Default, req.TypeOption)
	if err != nil {
		return nil, err
	}
	err = h.baseRepo.CreateColumn(schemaID, req.Table, req.ID, options)
	if err != nil {
		return nil, err
	}
	updatedMeta, err := h.colRepo.Update(ctx, column)
	if err != nil {
		return nil, err
	}
	return dto.NewColumnResp(updatedMeta, h.idService.EncodeID)
}

func (h columnService) Delete(ctx context.Context, id string) error {
	colID, err := h.idService.DecodeID(id)
	if err != nil {
		return err
	}
	column, err := h.colRepo.Get(ctx, colID)
	if err != nil {
		return err
	}
	table, err := column.QueryTable().Only(ctx)
	if err != nil {
		return err
	}
	tableID, err := h.idService.EncodeID(table.ID)
	if err != nil {
		return err
	}
	schema, err := table.QuerySchema().Only(ctx)
	if err != nil {
		return err
	}
	schemaID, err := h.idService.EncodeID(schema.ID)
	if err != nil {
		return err
	}
	err = h.baseRepo.DeleteColumn(schemaID, tableID, id)
	if err != nil {
		return err
	}
	return nil
}

func (h columnService) columnOptions(colType string, defaultValue string, option dto.TypeOption) (string, error) {
	// TODO
	switch colType {
	case "number":
		if option.Format == "decimal" {

		} else if option.Format == "integer" {

		}
		return "INT", nil
	default:
		return "TEXT", nil

	}
	return "", nil
}
