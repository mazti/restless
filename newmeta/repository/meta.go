package repository

import (
	"context"
	"github.com/mazti/restless/newmeta/ent"
	"time"
)

type MetaRepository interface {
	Create(ctx context.Context, meta ent.Meta) (*ent.Meta, error)
	Get(ctx context.Context, id int) (*ent.Meta, error)
	List(ctx context.Context, offset int, limit int) ([]*ent.Meta, error)
	Update(ctx context.Context, meta ent.Meta) (*ent.Meta, error)
	Delete(ctx context.Context, id int) error
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
	return repo.client.Meta.Create().
		SetBase(meta.Base).
		SetSchema(meta.Schema).
		Save(ctx)
}

func (repo *metaRepository) Get(ctx context.Context, id int) (*ent.Meta, error) {
	return repo.client.Meta.Get(ctx, id)
}

func (repo *metaRepository) List(ctx context.Context, offset int, limit int) ([]*ent.Meta, error) {
	return repo.client.Meta.
		Query().
		Limit(limit).
		Offset(offset).
		All(ctx)
}

func (repo *metaRepository) Update(ctx context.Context, meta ent.Meta) (*ent.Meta, error) {
	return repo.client.Meta.UpdateOneID(meta.ID).
		SetBase(meta.Base).
		SetUpdatedAt(time.Now()).
		Save(ctx)
}

func (repo *metaRepository) Delete(ctx context.Context, id int) error {
	return repo.client.Meta.DeleteOneID(id).Exec(ctx)
}

func (repo *metaRepository) Count(ctx context.Context) (int, error) {
	return repo.client.Meta.Query().Count(ctx)
}
