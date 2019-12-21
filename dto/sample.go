package dto

import (
	cdto "github.com/tiennv147/mazti-commons/dto"
)

type CreateSampleReq struct {
	Name     string `json:"name" validate:"string,min=1,required"`
	Provider string `json:"provider" validate:"string,min=1,required"`
}

type SampleResp struct {
	ID       int64  `json:"id" validate:"required"`
	Name     string `json:"name" validate:"string,min=1,required"`
	Provider string `json:"provider" validate:"string,min=1,required"`
}

type UpdateSampleReq struct {
	ID       int64  `json:"id" validate:"required"`
	Name     string `json:"name" validate:"string,min=1,required"`
	Provider string `json:"provider" validate:"string,min=1,required"`
}

type ListSampleResp struct {
	Metadata cdto.ListMetadata `json:"metadata"`
	Results  []SampleResp      `json:"results"`
}
