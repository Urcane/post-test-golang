package web

type TagCreateRequest struct {
	Label string `validate:"required, max=20, min=1" json:"label"`
}
