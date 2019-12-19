package dto

import (
	cdto "github.com/tiennv147/restless/commons/dto"
)

type CreateChannelReq struct {
	Name     string `json:"name" validate:"string,min=1,required"`
	Provider string `json:"provider" validate:"string,min=1,required"`
}

type ChannelResp struct {
	ID       int64  `json:"id" validate:"required"`
	Name     string `json:"name" validate:"string,min=1,required"`
	Provider string `json:"provider" validate:"string,min=1,required"`
}

type UpdateChannelReq struct {
	ID       int64  `json:"id" validate:"required"`
	Name     string `json:"name" validate:"string,min=1,required"`
	Provider string `json:"provider" validate:"string,min=1,required"`
}

type ListChannelResp struct {
	Metadata cdto.ListMetadata `json:"metadata"`
	Results  []ChannelResp     `json:"results"`
}
