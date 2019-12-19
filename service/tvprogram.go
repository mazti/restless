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

type TVProgramService interface {
	Create(ctx context.Context, req dto.CreateTVProgramReq) (dto.TVProgramResp, error)
	Get(ctx context.Context, id int64) (dto.TVProgramResp, error)
	List(ctx context.Context, offset int64, limit int) (cdto.ListResp, error)
	Update(ctx context.Context, req dto.UpdateTVProgramReq) (dto.TVProgramResp, error)
	Delete(ctx context.Context, id int64) error
}

func NewTVProgramService(repo repository.TVProgramRepository, logger log.Logger) TVProgramService {
	return tvProgramService{
		repo:   repo,
		logger: logger,
	}
}

type tvProgramService struct {
	repo   repository.TVProgramRepository
	logger log.Logger
}

func getTVProgram(req dto.CreateTVProgramReq) model.TVProgram {
	return model.TVProgram{
		Name:      req.Name,
		Provider:  req.Provider,
		ChannelID: req.ChannelID,
	}
}

func getTVProgramUpdate(req dto.UpdateTVProgramReq) model.TVProgram {
	return model.TVProgram{
		ID:        req.ID,
		Name:      req.Name,
		Provider:  req.Provider,
		ChannelID: req.ChannelID,
	}
}

func getTVProgramResp(model model.TVProgram) dto.TVProgramResp {
	return dto.TVProgramResp{
		ID:        model.ID,
		Name:      model.Name,
		Provider:  model.Provider,
		ChannelID: model.ChannelID,
	}
}

func (h tvProgramService) Create(ctx context.Context, req dto.CreateTVProgramReq) (dto.TVProgramResp, error) {
	ret := dto.TVProgramResp{}
	tvProgram := getTVProgram(req)
	result, err := h.repo.Create(tvProgram)
	if err != nil {
		return ret, err
	}
	return getTVProgramResp(result), nil
}

func (h tvProgramService) Get(ctx context.Context, id int64) (dto.TVProgramResp, error) {
	ret := dto.TVProgramResp{}
	tvProgram, err := h.repo.Get(id)
	if err != nil {
		return ret, err
	}
	return getTVProgramResp(tvProgram), nil
}

func (h tvProgramService) List(ctx context.Context, offset int64, limit int) (cdto.ListResp, error) {
	ret := cdto.ListResp{}
	items, err := h.repo.List(offset, limit)
	if err != nil {
		return ret, err
	}
	count := len(items)
	ret.Results = make([]interface{}, count)
	for i, item := range items {
		ret.Results[i] = getTVProgramResp(item)
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

func (h tvProgramService) Update(ctx context.Context, req dto.UpdateTVProgramReq) (dto.TVProgramResp, error) {
	ret := dto.TVProgramResp{}
	tvProgram := getTVProgramUpdate(req)
	err := h.repo.Update(tvProgram)
	if err != nil {
		return ret, err
	}
	return getTVProgramResp(tvProgram), nil
}

func (h tvProgramService) Delete(ctx context.Context, id int64) error {
	return h.repo.Delete(id)
}
