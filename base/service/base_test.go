package service

import (
	"context"
	"testing"

	"github.com/mazti/restless/base/dto"
	"github.com/mazti/restless/base/ent"
	"github.com/mazti/restless/base/repository/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func createTestSuite() (*mocks.BaseRepository, *mocks.MetaRepository, BaseService, error) {
	idService := NewIDService("salt")
	baseRepo := &mocks.BaseRepository{}
	metaRepo := &mocks.MetaRepository{}
	baseService, err := NewBaseService(baseRepo, metaRepo, idService)
	return baseRepo, metaRepo, baseService, err
}

// unit test for create Base
func TestCreateBaseSuccessfully(t *testing.T) {
	baseName := "test-create-base"

	ctx := context.Background()

	baseRepo, metaRepo, baseService, err := createTestSuite()
	assert.Nil(t, err)
	assert.NotNil(t, baseService)

	metaResp := &ent.Meta{
		Base: baseName,
	}

	metaRepo.On("Create", ctx, mock.Anything).Return(metaResp, nil)
	baseRepo.On("CreateSchema", mock.Anything).Return(nil)

	req := dto.CreateBaseReq{
		Base: baseName,
	}

	resp, createdErr := baseService.Create(ctx, req)

	assert.Nil(t, createdErr)
	assert.Equal(t, resp.Base, baseName)
}
// function have error , i will try fix this
func TestUpdateBaseSuccessfully(t *testing.T) {
	baseName := "test-update-base"
	//id := "v6qMj1Wd9EXAx30lb3aZPoJbe2KOyG"
	ctx := context.Background()

	_,metaRepo,baseService,err := createTestSuite()
	metaResp := &ent.Meta{
		Base: baseName,
	}
	metaRepo.On("Update", ctx,mock.Anything).Return( metaResp,nil)

	assert.Nil(t, err)
	assert.NotNil(t, baseService)

	req := dto.UpdateBaseReq{
		Base: baseName,
	}

	resp, updateErr := baseService.Update(ctx,req)

	assert.Nil(t, updateErr)
	assert.Equal(t, resp.Base, baseName)
}
// function have error , i will try fix this
func TestGetBaseSuccessfully(t *testing.T) {

	id := "v6qMj1Wd9EXAx30lb3aZPoJbe2KOyG"
	ctx := context.Background()
	_, metaRepo, baseService, err := createTestSuite()
	assert.Nil(t, err)
	assert.NotNil(t, baseService)

	metaRepo.On("Get", ctx, id).Return(nil)

	resp, getErr := baseService.Get(ctx, id)

	assert.Nil(t, getErr)
	assert.Equal(t, resp, dto.BaseResp{})
}
//function have error , i will try fix this in next time
func TestDeleteBaseSuccessfully(t *testing.T) {
	id := "v6qMj1Wd9EXAx30lb3aZPoJbe2KOyG"
	ctx := context.Background()
	_, metaRepo, baseService, err := createTestSuite()
	assert.Nil(t, err)
	assert.NotNil(t, baseService)

	metaRepo.On("Delete", ctx, id).Return(nil)

	getErr := baseService.Delete(ctx, id)

	assert.Nil(t, getErr)
}
