package repository

import (
	"context"
	"github.com/mazti/restless/base/ent"
	"github.com/mazti/restless/base/ent/meta"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type MetaRepository interface {
	Create(ctx context.Context, meta ent.Meta) (*ent.Meta, error)
	Get(ctx context.Context, id int) (*ent.Meta, error)
	List(ctx context.Context, offset int, limit int) ([]*ent.Meta, error)
	Update(ctx context.Context, meta ent.Meta) (*ent.Meta, error)
	Delete(ctx context.Context, id int) (*ent.Meta, error)
	Count(ctx context.Context) (int, error)
}

type metaRepository struct {
	client *ent.Client
}

func NewMetaRepository(client *ent.Client) MetaRepository {
	return &metaRepository{
		client: client,
	}
}

func (repo *metaRepository) Create(ctx context.Context, meta ent.Meta) (*ent.Meta, error) {
	return repo.client.Meta.
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

func (repo *metaRepository) Get(ctx context.Context, id int) (*ent.Meta, error) {
	m, err := repo.client.Meta.
		Query().
		Where(meta.ID(id), meta.DeletedAtIsNil()).
		First(ctx)
	return m, readError(err)
}

func (repo *metaRepository) List(ctx context.Context, offset int, limit int) ([]*ent.Meta, error) {
	metas, err := repo.client.Meta.
		Query().
		Where(meta.DeletedAtIsNil()).
		Limit(limit).
		Offset(offset).
		All(ctx)
	return metas, readError(err)
}

func (repo *metaRepository) Update(ctx context.Context, m ent.Meta) (*ent.Meta, error) {
	exist, err := repo.client.Meta.
		Query().
		Where(meta.ID(m.ID), meta.DeletedAtIsNil()).
		Exist(ctx)
	err = readError(err)
	if err != nil || !exist {
		return nil, err
	}
	return repo.client.Meta.
		UpdateOneID(m.ID).
		SetBase(m.Base).
		SetUpdatedAt(time.Now()).
		Save(ctx)
}

func (repo *metaRepository) Delete(ctx context.Context, id int) (*ent.Meta, error) {
	m, err := repo.client.Meta.
		UpdateOneID(id).
		SetDeletedAt(time.Now()).
		Save(ctx)
	return m, readError(err)
}

func (repo *metaRepository) Count(ctx context.Context) (int, error) {
	m, err := repo.client.Meta.Query().Count(ctx)
	return m, readError(err)
}
