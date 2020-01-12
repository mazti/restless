package service

import (
	"context"
	"github.com/mazti/restless/base/repository/mocks"
	"testing"

	"github.com/mazti/restless/base/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TableTestSuite() (*mocks.BaseRepository, *mocks.MetaRepository, TableService, error) {
	baseRepo := &mocks.BaseRepository{}
	metaRepo := &mocks.MetaRepository{}
	tableService, err := NewTableService(baseRepo, metaRepo)
	return baseRepo, metaRepo, tableService, err
}

func TestCreateTableBaseSuccessfully(t *testing.T) {
	nameTable := "name-table"
	ctx := context.Background()

	baseRepo, _, tableService, err :=  TableTestSuite()
	assert.Nil(t, err)
	assert.NotNil(t, tableService)


	baseRepo.On("CreateTable", mock.Anything).Return( nil)

	req := dto.CreateTableReq{
		Name: nameTable,
	}

	createdTableErr := tableService.Create(ctx, req)

	assert.Nil(t, createdTableErr)
}
