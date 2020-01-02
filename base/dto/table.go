package dto

import (
	"fmt"
	"strings"
)

type CreateTableReq struct {
	Schema  string   `json:"schema"`
	Name    string   `json:"name"`
	Columns []Column `json:"columns"`
}

func (req *CreateTableReq) ToQuery() string {
	var query []string
	query = append(query, "CREATE TABLE")
	query = append(query, fmt.Sprintf("`%s`.`%s`", req.Schema, req.Name))
	query = append(query, "(")

	var columns []string
	for _, c := range req.Columns {
		columns = append(columns, c.toQuery())
	}
	query = append(query, strings.Join(columns, ","))
	query = append(query, ")")

	return strings.Join(query, " ")
}

type Column struct {
	Name       string `json:"name"`
	Attributes string `json:"attributes"`
}

func (req *Column) toQuery() string {
	return fmt.Sprintf("%s %s", req.Name, req.Attributes)
}
