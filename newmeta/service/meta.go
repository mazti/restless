package service

import (
	"context"
	"github.com/speps/go-hashids"

	_ "github.com/jinzhu/gorm/dialects/mysql" //mysql dialects
	"github.com/mazti/restless/newmeta/dto"
	"github.com/mazti/restless/newmeta/repository"
	shared "github.com/mazti/restless/shared/dto"
)

type MetaService interface {
	Create(ctx context.Context, req dto.CreateMetaReq) (dto.MetaResp, error)
	Get(ctx context.Context, id string) (dto.MetaResp, error)
	List(ctx context.Context, offset int, limit int) (dto.ListMetaResp, error)
	Update(ctx context.Context, req dto.UpdateMetaReq) (dto.MetaResp, error)
	Delete(ctx context.Context, id string) error
}

func NewMetaService(repo repository.MetaRepository, hashID *hashids.HashID) MetaService {
	return metaService{
		repo:   repo,
		hashID: hashID,
	}
}

type metaService struct {
	repo   repository.MetaRepository
	hashID *hashids.HashID
}

func (h metaService) Create(ctx context.Context, req dto.CreateMetaReq) (resp dto.MetaResp, err error) {
	meta := req.ToMeta()
	result, err := h.repo.Create(ctx, meta)
	if err != nil {
		return resp, err
	}
	id, err := h.hashID.Encode([]int{result.ID})
	if err != nil {
		return resp, err
	}
	return dto.MetaResp{ID: id, Base: result.Base, Schema: result.Schema}, nil
}

func (h metaService) Get(ctx context.Context, id string) (resp dto.MetaResp, err error) {
	data, err := h.hashID.DecodeWithError(id)
	if err != nil {
		return resp, err
	}
	meta, err := h.repo.Get(ctx, data[0])
	if err != nil {
		return resp, err
	}
	return dto.MetaResp{ID: id, Base: meta.Base, Schema: meta.Schema}, nil
}

func (h metaService) List(ctx context.Context, offset int, limit int) (resp dto.ListMetaResp, err error) {
	items, err := h.repo.List(ctx, offset, limit)
	if err != nil {
		return resp, err
	}
	count := len(items)
	resp = dto.ListMetaResp{Results: make([]dto.MetaResp, count)}
	for i, item := range items {
		id, err := h.hashID.Encode([]int{item.ID})
		if err != nil {
			return resp, err
		}
		resp.Results[i] = dto.MetaResp{ID: id, Base: item.Base, Schema: item.Schema}
	}

	total, _ := h.repo.Count(ctx)
	resp.Metadata = shared.ListMetadata{
		Count:  count,
		Offset: offset,
		Limit:  limit,
		Total:  total,
	}
	return resp, nil
}

func (h metaService) Update(ctx context.Context, req dto.UpdateMetaReq) (resp dto.MetaResp, err error) {
	meta, err := req.ToMeta(h.hashID)
	if err != nil {
		return resp, err
	}
	updatedMeta, err := h.repo.Update(ctx, meta)
	if err != nil {
		return resp, err
	}
	return dto.MetaResp{ID: req.ID, Base: updatedMeta.Base, Schema: updatedMeta.Schema}, nil
}

func (h metaService) Delete(ctx context.Context, id string) error {
	data, err := h.hashID.DecodeWithError(id)
	if err != nil {
		return err
	}
	return h.repo.Delete(ctx, data[0])
}
