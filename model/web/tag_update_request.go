package web

type TagUpdateRequest struct {
	Id    int    `validate:"required" json:"id"`
	Label string `validate:"required, max=20, min=1" json:"label"`
}
