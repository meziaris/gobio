package model

type ApiResponse struct {
	Meta MetaResponse `json:"meta"`
	Data interface{}  `json:"data"`
}

type MetaResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}
