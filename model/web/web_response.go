package web

type WebResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"response"`
	Data   interface{} `json:"data"`
}
