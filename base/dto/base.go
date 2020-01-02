package dto

import (
	mazti "github.com/tiennv147/mazti-commons/dto"
)

type CreateBaseReq struct {
	Name string `json:"name"`
}

type BaseResp struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UpdateBaseReq struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ListBaseResp struct {
	Metadata mazti.ListMetadata `json:"metadata"`
	Results  []BaseResp         `json:"results"`
}
