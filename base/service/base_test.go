package service

import (
	"context"
	"github.com/mazti/restless/base/dto"
	"github.com/mazti/restless/base/ent"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/mazti/restless/base/repository/mocks"
)



func CommonVaribleTest() (*mocks.BaseRepository,*mocks.MetaRepository, BaseService, error)  {
	idService := NewIDService("salt")
	baseRepo := &mocks.BaseRepository{}
	metaRepo := &mocks.MetaRepository{}
	baseService, err := NewBaseService(baseRepo, metaRepo,idService)
	return baseRepo,metaRepo,baseService,err
}

// unit test for create Base
func TestCreateBaseSuccessfully(t *testing.T) {
	baseName := "test-create-base"

	ctx := context.Background()

	baseRepo,metaRepo,baseService,err := CommonVaribleTest()
	metaResp := &ent.Meta{
		Base: baseName,
	}
	//baseRepo
	metaRepo.On("Create", ctx, mock.Anything).Return(metaResp, nil)
	baseRepo.On("CreateSchema", mock.Anything ).Return(nil)

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

	id := "v6qMj1Wd9EXAx30lb3aZPoJbe2KOyG"
	ctx := context.Background()
	_,metaRepo,baseService,err := CommonVaribleTest()

	metaRepo.On("Get",ctx,id).Return(nil)

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

	resp, getErr := baseService.Get(ctx, id)

	assert.Nil(t,getErr)
	assert.Equal(t,resp,dto.BaseResp{})
}
func TestDeleteBaseSuccessfully(t *testing.T) {
	id := "v6qMj1Wd9EXAx30lb3aZPoJbe2KOyG"
	ctx := context.Background()
	_,metaRepo,baseService,err := CommonVaribleTest()

	metaRepo.On("Delete",ctx,id).Return(nil)

	assert.Nil(t, err)
	assert.NotNil(t, baseService)


	 getErr := baseService.Delete(ctx, id)

	assert.Nil(t,getErr)
}