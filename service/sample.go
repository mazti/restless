package service

import (
	"context"

	"github.com/go-kit/kit/log"
	_ "github.com/jinzhu/gorm/dialects/mysql" //mysql dialects

	cdto "github.com/tiennv147/mazti-commons/dto"
	"github.com/tiennv147/restless/dto"
	"github.com/tiennv147/restless/model"
	"github.com/tiennv147/restless/repository"
)

type SampleService interface {
	Create(ctx context.Context, req dto.CreateSampleReq) (dto.SampleResp, error)
	Get(ctx context.Context, id int) (dto.SampleResp, error)
	List(ctx context.Context, offset int, limit int) (cdto.ListResp, error)
	Update(ctx context.Context, req dto.UpdateSampleReq) (dto.SampleResp, error)
	Delete(ctx context.Context, id int) error
}

func NewSampleService(repo repository.SampleRepository, logger log.Logger) SampleService {
	return cameraService{
		repo:   repo,
		logger: logger,
	}
}

type cameraService struct {
	repo   repository.SampleRepository
	logger log.Logger
}

func getSample(req dto.CreateSampleReq) model.Sample {
	return model.Sample{
		Name:     req.Name,
		Provider: req.Provider,
	}
}

func getSampleUpdate(req dto.UpdateSampleReq) model.Sample {
	return model.Sample{
		ID:       req.ID,
		Name:     req.Name,
		Provider: req.Provider,
	}
}

func getSampleResp(model model.Sample) dto.SampleResp {
	return dto.SampleResp{
		ID:       model.ID,
		Name:     model.Name,
		Provider: model.Provider,
	}
}

func (h cameraService) Create(ctx context.Context, req dto.CreateSampleReq) (dto.SampleResp, error) {
	ret := dto.SampleResp{}
	camera := getSample(req)
	result, err := h.repo.Create(camera)
	if err != nil {
		return ret, err
	}
	return getSampleResp(result), nil
}

func (h cameraService) Get(ctx context.Context, id int) (dto.SampleResp, error) {
	ret := dto.SampleResp{}
	camera, err := h.repo.Get(id)
	if err != nil {
		return ret, err
	}
	return getSampleResp(camera), nil
}

func (h cameraService) List(ctx context.Context, offset int, limit int) (cdto.ListResp, error) {
	ret := cdto.ListResp{}
	items, err := h.repo.List(offset, limit)
	if err != nil {
		return ret, err
	}
	count := len(items)
	ret.Results = make([]interface{}, count)
	for i, item := range items {
		ret.Results[i] = getSampleResp(item)
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

func (h cameraService) Update(ctx context.Context, req dto.UpdateSampleReq) (dto.SampleResp, error) {
	ret := dto.SampleResp{}
	camera := getSampleUpdate(req)
	err := h.repo.Update(camera)
	if err != nil {
		return ret, err
	}
	return getSampleResp(camera), nil
}

func (h cameraService) Delete(ctx context.Context, id int) error {
	return h.repo.Delete(id)
}
