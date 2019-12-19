package service

import (
	"context"

	"github.com/go-kit/kit/log"
	_ "github.com/jinzhu/gorm/dialects/mysql" //mysql dialects

	cdto "github.com/tiennv147/restless/commons/dto"
	"github.com/tiennv147/restless/dto"
	"github.com/tiennv147/restless/model"
	"github.com/tiennv147/restless/repository"
)

type ChannelService interface {
	Create(ctx context.Context, req dto.CreateChannelReq) (dto.ChannelResp, error)
	Get(ctx context.Context, id int64) (dto.ChannelResp, error)
	List(ctx context.Context, offset int64, limit int) (cdto.ListResp, error)
	Update(ctx context.Context, req dto.UpdateChannelReq) (dto.ChannelResp, error)
	Delete(ctx context.Context, id int64) error
}

func NewChannelService(repo repository.ChannelRepository, logger log.Logger) ChannelService {
	return cameraService{
		repo:   repo,
		logger: logger,
	}
}

type cameraService struct {
	repo   repository.ChannelRepository
	logger log.Logger
}

func getChannel(req dto.CreateChannelReq) model.Channel {
	return model.Channel{
		Name:     req.Name,
		Provider: req.Provider,
	}
}

func getChannelUpdate(req dto.UpdateChannelReq) model.Channel {
	return model.Channel{
		ID:       req.ID,
		Name:     req.Name,
		Provider: req.Provider,
	}
}

func getChannelResp(model model.Channel) dto.ChannelResp {
	return dto.ChannelResp{
		ID:       model.ID,
		Name:     model.Name,
		Provider: model.Provider,
	}
}

func (h cameraService) Create(ctx context.Context, req dto.CreateChannelReq) (dto.ChannelResp, error) {
	ret := dto.ChannelResp{}
	camera := getChannel(req)
	result, err := h.repo.Create(camera)
	if err != nil {
		return ret, err
	}
	return getChannelResp(result), nil
}

func (h cameraService) Get(ctx context.Context, id int64) (dto.ChannelResp, error) {
	ret := dto.ChannelResp{}
	camera, err := h.repo.Get(id)
	if err != nil {
		return ret, err
	}
	return getChannelResp(camera), nil
}

func (h cameraService) List(ctx context.Context, offset int64, limit int) (cdto.ListResp, error) {
	ret := cdto.ListResp{}
	items, err := h.repo.List(offset, limit)
	if err != nil {
		return ret, err
	}
	count := len(items)
	ret.Results = make([]interface{}, count)
	for i, item := range items {
		ret.Results[i] = getChannelResp(item)
	}

	total, _ := h.repo.Count()

	ret.Metadata = cdto.ListMetadata{
		Count:  count,
		Offset: offset,
		Limit:  limit,
		Total:  total,
	}
	return ret, nil
}

func (h cameraService) Update(ctx context.Context, req dto.UpdateChannelReq) (dto.ChannelResp, error) {
	ret := dto.ChannelResp{}
	camera := getChannelUpdate(req)
	err := h.repo.Update(camera)
	if err != nil {
		return ret, err
	}
	return getChannelResp(camera), nil
}

func (h cameraService) Delete(ctx context.Context, id int64) error {
	return h.repo.Delete(id)
}
