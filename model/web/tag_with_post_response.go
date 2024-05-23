package web

type TagWithPostResponse struct {
	Id    int            `json:"id"`
	Label string         `json:"label"`
	Posts []PostResponse `json:"posts"`
}
