package dto

import (
	shared "github.com/mazti/restless/shared/dto"
)

type CreateMetaReq struct {
	Name string `json:"name" validate:"string,min=1,required"`
}

type MetaResp struct {
	ID   string `json:"id" validate:"required"`
	Name string `json:"name" validate:"string,min=1,required"`
}

type UpdateMetaReq struct {
	ID   string `json:"id" validate:"required"`
	Name string `json:"name" validate:"string,min=1,required"`
}

type ListMetaResp struct {
	Metadata shared.ListMetadata `json:"metadata"`
	Results  []MetaResp        `json:"results"`
}
