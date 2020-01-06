package repository

import (
	"context"
	"github.com/mazti/restless/base/ent"
	"github.com/mazti/restless/base/ent/metaschema"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type MetaSchemaRepository interface {
	Create(ctx context.Context, meta ent.MetaSchema) (*ent.MetaSchema, error)
	Get(ctx context.Context, id int) (*ent.MetaSchema, error)
	List(ctx context.Context, offset int, limit int) ([]*ent.MetaSchema, error)
	Update(ctx context.Context, meta ent.MetaSchema) (*ent.MetaSchema, error)
	Delete(ctx context.Context, id int) (*ent.MetaSchema, error)
	Count(ctx context.Context) (int, error)
}

type metaRepository struct {
	client *ent.Client
}

func NewMetaRepository(client *ent.Client) MetaSchemaRepository {
	return &metaRepository{
		client: client,
	}
}

func (repo *metaRepository) Create(ctx context.Context, meta ent.MetaSchema) (*ent.MetaSchema, error) {
	return repo.client.MetaSchema.
		Create().
		SetBase(meta.Base).
		Save(ctx)
}

func readError(err error) error {
	if err == nil {
		return err
	} else if ent.IsNotFound(err) {
		return status.Error(codes.NotFound, err.Error())
	}
	return err
}

func (repo *metaRepository) Get(ctx context.Context, id int) (*ent.MetaSchema, error) {
	m, err := repo.client.MetaSchema.
		Query().
		Where(metaschema.ID(id), metaschema.DeletedAtIsNil()).
		First(ctx)
	return m, readError(err)
}

func (repo *metaRepository) List(ctx context.Context, offset int, limit int) ([]*ent.MetaSchema, error) {
	metas, err := repo.client.MetaSchema.
		Query().
		Where(metaschema.DeletedAtIsNil()).
		Limit(limit).
		Offset(offset).
		All(ctx)
	return metas, readError(err)
}

func (repo *metaRepository) Update(ctx context.Context, m ent.MetaSchema) (*ent.MetaSchema, error) {
	exist, err := repo.client.MetaSchema.
		Query().
		Where(metaschema.ID(m.ID), metaschema.DeletedAtIsNil()).
		Exist(ctx)
	err = readError(err)
	if err != nil || !exist {
		return nil, err
	}
	return repo.client.MetaSchema.
		UpdateOneID(m.ID).
		SetBase(m.Base).
		SetUpdatedAt(time.Now()).
		Save(ctx)
}

func (repo *metaRepository) Delete(ctx context.Context, id int) (*ent.MetaSchema, error) {
	m, err := repo.client.MetaSchema.
		UpdateOneID(id).
		SetDeletedAt(time.Now()).
		Save(ctx)
	return m, readError(err)
}

func (repo *metaRepository) Count(ctx context.Context) (int, error) {
	m, err := repo.client.MetaSchema.Query().Count(ctx)
	return m, readError(err)
}
