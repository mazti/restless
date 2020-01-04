package dto

import (
	"github.com/mazti/restless/base/ent"
	shared "github.com/mazti/restless/shared/dto"
)

type CreateBaseReq struct {
	Base string `json:"base"`
}

type BaseResp struct {
	ID   string `json:"id"`
	Base string `json:"base"`
	Schema string `json:"schema"`
}

type UpdateBaseReq struct {
	ID   string `json:"id"`
	Base string `json:"base"`
	Schema string `json:"schema"`
}

type ListBaseResp struct {
	Metadata shared.ListMetadata `json:"metadata"`
	Results  []BaseResp         `json:"results"`
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
