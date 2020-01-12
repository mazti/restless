package service

import (
	"context"
	"github.com/mazti/restless/base/dto"
	"github.com/mazti/restless/base/ent"
	"github.com/mazti/restless/base/repository/mocks"
	"github.com/stretchr/testify/mock"
	"testing"

	"github.com/stretchr/testify/assert"
)

func tableTestSuite() (*mocks.BaseRepository, *mocks.MetaTableRepository, TableService, error) {
	idService := NewIDService("123")
	baseRepo := &mocks.BaseRepository{}
	metaTableRepo := &mocks.MetaTableRepository{}
	tableService, err := NewTableService(baseRepo, metaTableRepo, idService)
	return baseRepo, metaTableRepo, tableService, err
}

func TestCreateTableSuccessfully(t *testing.T) {
	nameTable := "b7j5kmE0qvB4rlRNlndK8awgA16zVN"
	schemaId := 123
	schemaName := "E4gVbBwDy013oLPOxlMp5r7k6WnR9Z"
	ctx := context.Background()

	table := ent.MetaTable{
		Name: nameTable,
	}
	baseRepo, tableRepo, tableService, err := tableTestSuite()
	assert.Nil(t, err)
	assert.NotNil(t, tableService)

	metaResp := &ent.MetaTable{
		Name: nameTable,
	}

	tableRepo.On("Create", ctx, schemaId, table).Return(metaResp, nil)
	baseRepo.On("CreateTable", schemaName, nameTable, mock.Anything).Return(nil)

	req := dto.CreateTableReq{
		Schema: schemaName,
		Name:   nameTable,
	}

	resp, createdTableErr := tableService.Create(ctx, req)

	assert.Nil(t, createdTableErr)
	assert.NotEmpty(t, resp)
}
