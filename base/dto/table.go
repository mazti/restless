package dto

import (
	"fmt"
	"github.com/mazti/restless/base/ent"
)

func NewTableResp(schemaId int, table ent.MetaTable, EncodeID func(int) (string, error)) (resp *TableResp, err error) {
	id, err := EncodeID(table.ID)
	if err != nil {
		return resp, err
	}
	sId, err := EncodeID(schemaId)
	if err != nil {
		return resp, err
	}
	return &TableResp{
		ID:        id,
		SchemaID:  sId,
		Name:      table.Name,
		CreatedAt: table.CreatedAt.UnixNano() / 1000000,
		UpdatedAt: table.UpdatedAt.UnixNano() / 1000000,
	}, nil
}

type TableResp struct {
	ID        string `json:"id"`
	SchemaID  string `json:"schema_id"`
	Name      string `json:"name"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type CreateTableReq struct {
	Schema  string   `json:"schema"`
	Name    string   `json:"name"`
	Columns []Column `json:"columns"`
}

func (req CreateTableReq) ToTable(DecodeID func(string) (int, error)) (schemaId int, table ent.MetaTable, err error) {
	schemaId, err = DecodeID(req.Schema)
	if err != nil {
		return 0, table, err
	}
	table = ent.MetaTable{
		Name: req.Name,
	}
	return schemaId, table, nil
}

type Column struct {
	Name       string `json:"name"`
	Attributes string `json:"attributes"`
}

func (req *Column) ToQuery() string {
	return fmt.Sprintf("%s %s", req.Name, req.Attributes)
}
