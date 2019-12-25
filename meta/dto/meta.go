package dto

import (
	cdto "github.com/tiennv147/mazti-commons/dto"
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
	Metadata cdto.ListMetadata `json:"metadata"`
	Results  []MetaResp        `json:"results"`
}
