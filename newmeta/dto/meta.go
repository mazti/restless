package dto

import (
	"github.com/mazti/restless/newmeta/ent"
	shared "github.com/mazti/restless/shared/dto"
	"github.com/speps/go-hashids"
)

type CreateMetaReq struct {
	Base string `json:"base"`
}

type MetaResp struct {
	ID     string `json:"id"`
	Base   string `json:"base"`
	Schema string `json:"schema"`
}

type UpdateMetaReq struct {
	ID   string `json:"id"`
	Base string `json:"base"`
}

type ListMetaResp struct {
	Metadata shared.ListMetadata `json:"metadata"`
	Results  []MetaResp        `json:"results"`
}

func (req CreateMetaReq) ToMeta() ent.Meta {
	return ent.Meta{
		Base: req.Base,
	}
}

func (req UpdateMetaReq) ToMeta(hash *hashids.HashID) (meta ent.Meta, err error) {
	data, err := hash.DecodeWithError(req.ID)
	if err != nil {
		return meta, err
	}
	return ent.Meta{
		ID:   data[0],
		Base: req.Base,
	}, nil
}

