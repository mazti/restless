package repository

import (
	"context"
	"github.com/mazti/restless/base/ent"
	"github.com/mazti/restless/base/ent/metacolumn"
	"time"
)

type MetaColumnRepository interface {
	Create(ctx context.Context, tableId int, column *ent.MetaColumn) (*ent.MetaColumn, error)
	Get(ctx context.Context, id int) (*ent.MetaColumn, error)
	List(ctx context.Context, offset int, limit int) ([]*ent.MetaColumn, error)
	Update(ctx context.Context, column *ent.MetaColumn) (*ent.MetaColumn, error)
	Delete(ctx context.Context, id int) (*ent.MetaColumn, error)
	Count(ctx context.Context) (int, error)
}

type columnRepository struct {
	client *ent.Client
}

func NewColumnRepository(client *ent.Client) MetaColumnRepository {
	return &columnRepository{
		client: client,
	}
}

func (repo *columnRepository) Create(ctx context.Context, tableId int, column *ent.MetaColumn) (*ent.MetaColumn, error) {
	return repo.client.MetaColumn.
		Create().
		SetName(column.Name).
		SetTableID(tableId).
		Save(ctx)
}

func (repo *columnRepository) Get(ctx context.Context, id int) (*ent.MetaColumn, error) {
	m, err := repo.client.MetaColumn.
		Query().
		Where(metacolumn.ID(id), metacolumn.DeletedAtIsNil()).
		First(ctx)
	return m, readError(err)
}

func (repo *columnRepository) List(ctx context.Context, offset int, limit int) ([]*ent.MetaColumn, error) {
	columns, err := repo.client.MetaColumn.
		Query().
		Where(metacolumn.DeletedAtIsNil()).
		Limit(limit).
		Offset(offset).
		All(ctx)
	return columns, readError(err)
}

func (repo *columnRepository) Update(ctx context.Context, m *ent.MetaColumn) (*ent.MetaColumn, error) {
	exist, err := repo.client.MetaColumn.
		Query().
		Where(metacolumn.ID(m.ID), metacolumn.DeletedAtIsNil()).
		Exist(ctx)
	err = readError(err)
	if err != nil || !exist {
		return nil, err
	}
	return repo.client.MetaColumn.
		UpdateOneID(m.ID).
		SetName(m.Name).
		SetType(m.Type).
		SetDefault(m.Default).
		SetTypeOption(m.TypeOption).
		SetUpdatedAt(time.Now()).
		Save(ctx)
}

func (repo *columnRepository) Delete(ctx context.Context, id int) (*ent.MetaColumn, error) {
	m, err := repo.client.MetaColumn.
		UpdateOneID(id).
		SetDeletedAt(time.Now()).
		Save(ctx)
	return m, readError(err)
}

func (repo *columnRepository) Count(ctx context.Context) (int, error) {
	m, err := repo.client.MetaColumn.Query().Count(ctx)
	return m, readError(err)
}
