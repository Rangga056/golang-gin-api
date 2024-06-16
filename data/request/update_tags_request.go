package request

type UpdateTagRequest struct {
	//validate is a struct tag used for validation
	Id   int    `validate:"required"`
	Name string `validate:"required, max=200, min=1" json:"name"`
}
