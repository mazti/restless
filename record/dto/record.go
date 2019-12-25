package dto

type SortParameter struct {
	Field     string `json:"field"`
	Direction string `json:"direction"`
}

type SelectRecordsReq struct {
	Fields          []string        `json:"fields"`
	FilterByFormula string          `json:"filterByFormula"`
	MaxRecords      int             `json:"maxRecords"`
	PageSize        int             `json:"pageSize"`
	Sort            []SortParameter `json:"sort"`
	View            string          `json:"view"`
}

type SelectRecordsResp struct {
	Records []interface{} `json:"records"`
	Offset  string        `json:"offset"`
}
