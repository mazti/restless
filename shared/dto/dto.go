package dto

type ListMetadata struct {
	Count  int   `json:"count"`
	Offset int   `json:"offset"`
	Limit  int   `json:"limit"`
	Total  int `json:"total"`
}

type ListResp struct {
	Metadata ListMetadata  `json:"metadata"`
	Results  []interface{} `json:"results"`
}

type ListRequest struct {
	Limit  int
	Offset int
}
