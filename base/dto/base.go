package dto

import (
	"github.com/mazti/restless/base/ent"
	shared "github.com/mazti/restless/shared/dto"
)

type CreateBaseReq struct {
	Base string `json:"base"`
}

func NewBaseResp(meta *ent.MetaSchema, EncodeID func(id int) (string, error)) (resp BaseResp, err error) {
	id, err := EncodeID(meta.ID)
	if err != nil {
		return resp, err
	}
	return BaseResp{
		ID:        id,
		Base:      meta.Base,
		CreatedAt: meta.CreatedAt.UnixNano() / 1000000,
		UpdatedAt: meta.UpdatedAt.UnixNano() / 1000000,
	}, nil
}

type BaseResp struct {
	ID        string `json:"id"`
	Base      string `json:"base"`
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

func (req CreateBaseReq) ToMeta() ent.MetaSchema {
	return ent.MetaSchema{
		Base: req.Base,
	}
}

func (req UpdateBaseReq) ToMeta(DecodeID func(id string) (int, error)) (meta ent.MetaSchema, err error) {
	id, err := DecodeID(req.ID)
	if err != nil {
		return meta, err
	}
	return ent.MetaSchema{
		ID:   id,
		Base: req.Base,
	}, nil
}
