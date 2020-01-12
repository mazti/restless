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

func createTestSuite() (*mocks.BaseRepository, *mocks.MetaSchemaRepository, BaseService, error) {
	idService := NewIDService("123")
	baseRepo := &mocks.BaseRepository{}
	metaRepo := &mocks.MetaSchemaRepository{}
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

	metaResp := &ent.MetaSchema{
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

//function unit test update
func TestUpdateBaseSuccessfully(t *testing.T) {
	baseName := "test-update-base"
	id := "E4gVbBwDy013oLPOxlMp5r7k6WnR9Z"
	//id1 := "123"

	ctx := context.Background()

	_, metaRepo, baseService, err := createTestSuite()
	metaResp := &ent.MetaSchema{
		Base: baseName,
	}
	metaRepo.On("Update", ctx, mock.Anything).Return(metaResp, nil)

	assert.Nil(t, err)
	assert.NotNil(t, baseService)

	req := dto.UpdateBaseReq{
		ID: id,
	}

	resp, updateErr := baseService.Update(ctx, req)
	//
	assert.Nil(t, updateErr)
	assert.NotEmpty(t, resp)
}

//function unit test list
func TestListBaseSuccessfully(t *testing.T) {
	baseName := "test-list-base"
	offset := 4
	limit := 4
	ctx := context.Background()
	_, metaRepo, baseService, err := createTestSuite()

	assert.Nil(t, err)
	assert.NotNil(t, baseService)

	//var metaResp []*ent.Meta

	var metaResp []*ent.MetaSchema
	metaRespElement := &ent.MetaSchema{
		Base: baseName,
	}
	metaResp = append(metaResp, metaRespElement)
	mockCount := 0

	metaRepo.On("List", ctx, offset, limit).Return(metaResp, nil)
	metaRepo.On("Count", ctx).Return(mockCount, nil)

	resp, getErr := baseService.List(ctx, offset, limit)

	assert.Nil(t, getErr)
	assert.NotEmpty(t, resp)
}
func TestGetBaseSuccessfully(t *testing.T) {
	baseName := "test-get-base"

	id := "E4gVbBwDy013oLPOxlMp5r7k6WnR9Z"
	idDecode := 123
	ctx := context.Background()
	_, metaRepo, baseService, err := createTestSuite()
	assert.Nil(t, err)
	assert.NotNil(t, baseService)
	metaResp := &ent.MetaSchema{
		Base: baseName,
	}
	metaRepo.On("Get", ctx, idDecode).Return(metaResp, nil)

	resp, getErr := baseService.Get(ctx, id)

	assert.Nil(t, getErr)
	assert.NotEmpty(t, resp)
}

func TestDeleteBaseSuccessfully(t *testing.T) {
	id := 123
	idDecode := "E4gVbBwDy013oLPOxlMp5r7k6WnR9Z"

	baseName := "test-delete-base"

	ctx := context.Background()
	_, metaRepo, baseService, err := createTestSuite()

	metaResp := &ent.MetaSchema{
		Base: baseName,
	}
	assert.Nil(t, err)
	assert.NotNil(t, baseService)

	metaRepo.On("Delete", ctx, id).Return(metaResp, nil)

	getErr := baseService.Delete(ctx, idDecode)

	assert.Nil(t, getErr)
}
