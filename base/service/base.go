package service

import (
	"context"
	"errors"
	mazti "github.com/tiennv147/mazti-commons/dto"
	"github.com/tiennv147/restless/base/config"
	"github.com/tiennv147/restless/base/dto"
	"github.com/tiennv147/restless/base/repository"
	pb "github.com/tiennv147/restless/meta/pb/meta"
	"google.golang.org/grpc"
)

type BaseService interface {
	Create(ctx context.Context, req dto.CreateBaseReq) (dto.BaseResp, error)
	Get(ctx context.Context, id string) (dto.BaseResp, error)
	List(ctx context.Context, offset int, limit int) (dto.ListBaseResp, error)
	Update(ctx context.Context, req dto.UpdateBaseReq) (dto.BaseResp, error)
	Delete(ctx context.Context, id string) error
}

func NewBaseService(repo repository.BaseRepository, conf config.Config) (BaseService, error) {
	conn, err := grpc.Dial(conf.Meta, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &baseService{
		repo:       repo,
		metaClient: pb.NewMetaClient(conn),
	}, nil
}

type baseService struct {
	repo       repository.BaseRepository
	metaClient pb.MetaClient
}

func (h baseService) Create(ctx context.Context, req dto.CreateBaseReq) (base dto.BaseResp, err error) {
	resp, err := h.metaClient.Create(ctx, &pb.CreateMetaRequest{Name: req.Name})
	if err != nil {
		return base, err
	}
	if resp == nil {
		return base, errors.New("create base meta fail")
	}
	err = h.repo.Create(dto.CreateBaseReq{Name:resp.Id})
	if err != nil {
		return base, err
	}
	baseResp := dto.BaseResp{ID: resp.Id, Name: resp.Name}
	return baseResp, err
}

func (h baseService) Get(ctx context.Context, id string) (dto.BaseResp, error) {
	ret := dto.BaseResp{}
	meta, err := h.metaClient.Get(ctx, &pb.GetMetaRequest{Id: id})
	if err != nil {
		return ret, err
	}
	return dto.BaseResp{
		ID:   meta.Id,
		Name: meta.Name,
	}, nil
}

func (h baseService) List(ctx context.Context, offset int, limit int) (dto.ListBaseResp, error) {
	ret := dto.ListBaseResp{}
	reply, err := h.metaClient.List(ctx, &pb.ListMetaRequest{Offset: int32(offset), Limit: int32(limit)})
	if err != nil {
		return ret, err
	}
	ret.Results = make([]dto.BaseResp, reply.Metadata.Count)
	for i, item := range reply.Metas {
		ret.Results[i] = dto.BaseResp{
			ID:   item.Id,
			Name: item.Name,
		}
	}

	ret.Metadata = mazti.ListMetadata{
		Count:  int(reply.Metadata.Count),
		Offset: int(reply.Metadata.Offset),
		Limit:  int(reply.Metadata.Limit),
		Total:  reply.Metadata.Total,
	}
	return ret, nil
}

func (h baseService) Update(ctx context.Context, req dto.UpdateBaseReq) (dto.BaseResp, error) {
	ret := dto.BaseResp{}
	meta, err := h.metaClient.Update(ctx, &pb.UpdateMetaRequest{Id: req.ID, Name: req.Name})
	if err != nil {
		return ret, err
	}
	return dto.BaseResp{
		ID:   meta.Id,
		Name: meta.Name,
	}, nil
}

func (h baseService) Delete(ctx context.Context, id string) error {
	_, err := h.metaClient.Delete(ctx, &pb.DeleteMetaRequest{Id: id})
	return err
}
