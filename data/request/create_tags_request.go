package request

type CreateTagsRequests struct {
	//validate is a struct tag used for validation
	Name string `validate:"required, min=1, max=10" json:"name"`
}
