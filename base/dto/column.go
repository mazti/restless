package dto

import (
	"encoding/json"
	"github.com/mazti/restless/base/ent"
)

type CreateColumnReq struct {
	Table      string     `json:"table"`
	Name       string     `json:"name"`
	Type       string     `json:"type"`
	Default    string     `json:"default"`
	TypeOption TypeOption `json:"typeOptions"`
}

type TypeOption struct {
	Format    string `json:"format"`
	Precision string `json:"precision"`
}

func NewColumnResp(col *ent.MetaColumn, EncodeID func(int) (string, error)) (resp *ColumnResp, err error) {
	id, err := EncodeID(col.ID)
	if err != nil {
		return resp, err
	}
	var typeOption = &TypeOption{}
	if col.TypeOption != "" {
		if err := json.Unmarshal([]byte(col.TypeOption), typeOption); err != nil {
			return resp, err
		}
	}

	return &ColumnResp{
		ID:         id,
		Name:       col.Name,
		Type:       col.Type,
		Default:    col.Default,
		TypeOption: typeOption,
		CreatedAt:  col.CreatedAt.UnixNano() / 1000000,
		UpdatedAt:  col.UpdatedAt.UnixNano() / 1000000,
	}, nil
}

type ColumnResp struct {
	ID         string     `json:"id"`
	Name       string     `json:"name"`
	Type       string     `json:"type"`
	Default    string     `json:"default"`
	TypeOption *TypeOption `json:"typeOption"`
	CreatedAt  int64      `json:"created_at"`
	UpdatedAt  int64      `json:"updated_at"`
}

type UpdateColumnReq struct {
	ID          string     `json:"id"`
	Table       string     `json:"table"`
	Name        string     `json:"name"`
	Type        string     `json:"type"`
	Default     string     `json:"default"`
	TypeOption TypeOption `json:"typeOptions"`
}

type ListColumnResp struct {
	Results []ColumnResp `json:"results"`
}

func (req CreateColumnReq) ToMeta(DecodeID func(string) (int, error)) (tableID int, column *ent.MetaColumn, err error) {

	tableID, err = DecodeID(req.Table)
	if err != nil {
		return 0, nil, err
	}

	column = &ent.MetaColumn{
		Name:    req.Name,
		Type:    req.Type,
		Default: req.Default,
	}
	typeOptions, err := json.Marshal(req.TypeOption)
	if err != nil {
		return 0, nil, err
	}
	column.TypeOption = string(typeOptions)
	return tableID, column, nil
}

func (req UpdateColumnReq) ToMeta(DecodeID func(string) (int, error)) (tableID int, meta *ent.MetaColumn, err error) {
	tableID, err = DecodeID(req.ID)
	if err != nil {
		return 0, nil, err
	}
	colID, err := DecodeID(req.ID)
	if err != nil {
		return 0, nil, err
	}
	column := &ent.MetaColumn{
		ID:      colID,
		Name:    req.Name,
		Type:    req.Type,
		Default: req.Default,
	}
	typeOption, err := json.Marshal(req.TypeOption)
	if err != nil {
		return 0, nil, err
	}
	column.TypeOption = string(typeOption)
	return tableID, column, nil
}
