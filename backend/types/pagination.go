package types

type PaginationQuery struct {
	PageNumber int `json:"pageNumber"`
	PageSize   int `json:"pageSize"`
}

type PaginationResponse struct {
	PageNumber int `json:"pageNumber"`
	PageSize   int `json:"pageSize"`
	Total      int `json:"total"`
}
