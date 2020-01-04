package service

import (
	"context"

	_ "github.com/jinzhu/gorm/dialects/mysql" //mysql dialects
	"github.com/mazti/restless/meta/dto"
	"github.com/mazti/restless/meta/model"
	"github.com/mazti/restless/meta/repository"
	shared "github.com/mazti/restless/shared/dto"
)

type MetaService interface {
	Create(ctx context.Context, req dto.CreateMetaReq) (dto.MetaResp, error)
	Get(ctx context.Context, id string) (dto.MetaResp, error)
	List(ctx context.Context, offset int, limit int) (dto.ListMetaResp, error)
	Update(ctx context.Context, req dto.UpdateMetaReq) (dto.MetaResp, error)
	Delete(ctx context.Context, id string) error
}

func NewMetaService(repo repository.MetaRepository) MetaService {
	return metaService{
		repo: repo,
	}
}

type metaService struct {
	repo repository.MetaRepository
}

func getMeta(req dto.CreateMetaReq) model.Meta {
	return model.Meta{
		Name: req.Name,
	}
}

func getMetaUpdate(req dto.UpdateMetaReq) model.Meta {
	return model.Meta{
		ID:   req.ID,
		Name: req.Name,
	}
}

func getMetaResp(model model.Meta) dto.MetaResp {
	return dto.MetaResp{
		ID:   model.ID,
		Name: model.Name,
	}
}

func (h metaService) Create(ctx context.Context, req dto.CreateMetaReq) (dto.MetaResp, error) {
	ret := dto.MetaResp{}
	meta := getMeta(req)
	result, err := h.repo.Create(meta)
	if err != nil {
		return ret, err
	}
	return getMetaResp(result), nil
}

func (h metaService) Get(ctx context.Context, id string) (dto.MetaResp, error) {
	ret := dto.MetaResp{}
	meta, err := h.repo.Get(id)
	if err != nil {
		return ret, err
	}
	return getMetaResp(meta), nil
}

func (h metaService) List(ctx context.Context, offset int, limit int) (dto.ListMetaResp, error) {
	ret := dto.ListMetaResp{}
	items, err := h.repo.List(offset, limit)
	if err != nil {
		return ret, err
	}
	count := len(items)
	ret.Results = make([]dto.MetaResp, count)
	for i, item := range items {
		ret.Results[i] = getMetaResp(item)
	}

	total, _ := h.repo.Count()

	ret.Metadata = shared.ListMetadata{
		Count:  count,
		Offset: offset,
		Limit:  limit,
		Total:  total,
	}
	return ret, nil
}

func (h metaService) Update(ctx context.Context, req dto.UpdateMetaReq) (dto.MetaResp, error) {
	ret := dto.MetaResp{}
	meta := getMetaUpdate(req)
	err := h.repo.Update(meta)
	if err != nil {
		return ret, err
	}
	return getMetaResp(meta), nil
}

func (h metaService) Delete(ctx context.Context, id string) error {
	return h.repo.Delete(id)
}
