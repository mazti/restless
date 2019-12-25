package dto

type CreateSchemaReq struct {
	Name string `json:"name" validate:"string,min=1,required"`
}
