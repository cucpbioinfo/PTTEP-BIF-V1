package types

type BaseResponse struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}

type HttpError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
