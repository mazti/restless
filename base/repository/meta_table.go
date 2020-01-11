package repository

import (
	"context"
	"github.com/mazti/restless/base/ent"
	"github.com/mazti/restless/base/ent/metatable"
	"time"
)

type MetaTableRepository interface {
	Create(ctx context.Context, schemaId int, table ent.MetaTable) (*ent.MetaTable, error)
	Get(ctx context.Context, id int) (*ent.MetaTable, error)
	List(ctx context.Context, offset int, limit int) ([]*ent.MetaTable, error)
	Update(ctx context.Context, table ent.MetaTable) (*ent.MetaTable, error)
	Delete(ctx context.Context, id int) (*ent.MetaTable, error)
	Count(ctx context.Context) (int, error)
}

type tableRepository struct {
	client *ent.Client
}

func NewTableRepository(client *ent.Client) MetaTableRepository {
	return &tableRepository{
		client: client,
	}
}

func (repo *tableRepository) Create(ctx context.Context, schemaId int, table ent.MetaTable) (*ent.MetaTable, error) {
	return repo.client.MetaTable.
		Create().
		SetName(table.Name).
		SetSchemaID(schemaId).
		Save(ctx)
}

func (repo *tableRepository) Get(ctx context.Context, id int) (*ent.MetaTable, error) {
	m, err := repo.client.MetaTable.
		Query().
		Where(metatable.ID(id), metatable.DeletedAtIsNil()).
		First(ctx)
	return m, readError(err)
}

func (repo *tableRepository) List(ctx context.Context, offset int, limit int) ([]*ent.MetaTable, error) {
	tables, err := repo.client.MetaTable.
		Query().
		Where(metatable.DeletedAtIsNil()).
		Limit(limit).
		Offset(offset).
		All(ctx)
	return tables, readError(err)
}

func (repo *tableRepository) Update(ctx context.Context, m ent.MetaTable) (*ent.MetaTable, error) {
	exist, err := repo.client.MetaTable.
		Query().
		Where(metatable.ID(m.ID), metatable.DeletedAtIsNil()).
		Exist(ctx)
	err = readError(err)
	if err != nil || !exist {
		return nil, err
	}
	return repo.client.MetaTable.
		UpdateOneID(m.ID).
		SetName(m.Name).
		SetUpdatedAt(time.Now()).
		Save(ctx)
}

func (repo *tableRepository) Delete(ctx context.Context, id int) (*ent.MetaTable, error) {
	m, err := repo.client.MetaTable.
		UpdateOneID(id).
		SetDeletedAt(time.Now()).
		Save(ctx)
	return m, readError(err)
}

func (repo *tableRepository) Count(ctx context.Context) (int, error) {
	m, err := repo.client.MetaTable.Query().Count(ctx)
	return m, readError(err)
}
