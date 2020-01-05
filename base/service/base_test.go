package service


import (
	"context"
	"github.com/mazti/restless/base/dto"
	"github.com/mazti/restless/base/ent"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/mazti/restless/base/repository"
	"github.com/mazti/restless/base/repository/mocks"
)



func CommonVaribleTest() (repository.BaseRepository,repository.MetaRepository, BaseService)  {
	baseRepo := &mocks.BaseRepository{}
	metaRepo := &mocks.MetaRepository{}
	baseService, _ := NewBaseService(baseRepo, metaRepo)
	return baseRepo,metaRepo,baseService
}

func TestCreateBaseSuccessfully(t *testing.T) {
	baseName := "tan-base"
	Init("abc")

	ctx := context.Background()

	baseRepo := &mocks.BaseRepository{}
	metaRepo := &mocks.MetaRepository{}
	//baseRepo,metaRepo,baseService := CommonVaribleTest()
	metaResp := &ent.Meta{
		Base: baseName,
	}
	//baseRepo
	metaRepo.On("Create", ctx, mock.Anything).Return(metaResp, nil)
	baseRepo.On("CreateSchema", mock.Anything ).Return(nil)

	baseService, err := NewBaseService(baseRepo, metaRepo)

	assert.Nil(t, err)
	assert.NotNil(t, baseService)



	req := dto.CreateBaseReq{
		Base: baseName,
	}

	resp, createdErr := baseService.Create(ctx, req)

	assert.Nil(t, createdErr)
	assert.Equal(t, resp.Base, baseName)
}
func TestGetBaseSuccessfully(t *testing.T) {
	Init("abc")

	//id := "1206"
	ctx := context.Background()
	metaRepo := &mocks.MetaRepository{}
	baseRepo := &mocks.BaseRepository{}

	//metaResp := &ent.Meta{
	//	ID: ,
	//}
	metaRepo.On("Get",ctx,mock.Anything).Return(nil)

	baseService, err := NewBaseService(baseRepo, metaRepo)

	assert.Nil(t, err)
	assert.NotNil(t, baseService)

	//req := dto.BaseResp{
	//	ID: "",
	//	Base :"",
	//	Schema:"",
	//	CreatedAt:1,
	//	UpdatedAt:1,
	//
	//}

	resp, getErr := baseService.Get(ctx, "1206")

	assert.Nil(t,getErr)
	assert.Equal(t,resp,dto.BaseResp{})
}