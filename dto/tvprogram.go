package dto

import (
	cdto "github.com/tiennv147/restless/commons/dto"
)

type CreateTVProgramReq struct {
	Name      string `json:"name" validate:"string,min=1,required"`
	Provider  string `json:"provider" validate:"string,min=1,required"`
	ChannelID int64  `json:"ChannelID"`
}

type TVProgramResp struct {
	ID        int64  `json:"id" validate:"required"`
	Name      string `json:"name" validate:"string,min=1,required"`
	Provider  string `json:"provider" validate:"string,min=1,required"`
	ChannelID int64  `json:"ChannelID"`
}

type UpdateTVProgramReq struct {
	ID        int64  `json:"id" validate:"required"`
	Name      string `json:"name" validate:"string,min=1,required"`
	Provider  string `json:"provider" validate:"string,min=1,required"`
	ChannelID int64  `json:"ChannelID"`
}

type ListTVProgramResp struct {
	Metadata cdto.ListMetadata `json:"metadata"`
	Results  []ChannelResp     `json:"results"`
}
