package dto

import (
	"github.com/mazti/restless/base/ent"
	shared "github.com/mazti/restless/shared/dto"
)

type CreateBaseReq struct {
	Base string `json:"base"`
}

func NewBaseResp(meta *ent.Meta, EncodeID func(id int) (string, error)) (resp BaseResp, err error) {
	id, err := EncodeID(meta.ID)
	if err != nil {
		return resp, err
	}
	return BaseResp{
		ID:        id,
		Base:      meta.Base,
		Schema:    meta.Schema,
		CreatedAt: meta.CreatedAt.Unix(),
		UpdatedAt: meta.UpdatedAt.Unix(),
	}, nil
}

type BaseResp struct {
	ID        string `json:"id"`
	Base      string `json:"base"`
	Schema    string `json:"schema"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type UpdateBaseReq struct {
	ID   string `json:"id"`
	Base string `json:"base"`
}

type ListBaseResp struct {
	Metadata shared.ListMetadata `json:"metadata"`
	Results  []BaseResp          `json:"results"`
}

func (req CreateBaseReq) ToMeta() ent.Meta {
	return ent.Meta{
		Base: req.Base,
	}
}

func (req UpdateBaseReq) ToMeta(DecodeID func(id string) (int, error)) (meta ent.Meta, err error) {
	id, err := DecodeID(req.ID)
	if err != nil {
		return meta, err
	}
	return ent.Meta{
		ID:   id,
		Base: req.Base,
	}, nil
}
