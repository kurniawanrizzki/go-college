package dto

type Pagination struct {
	Page       int64
	PerPage    int64
	PageCount  int64
	TotalCount int64
}

type PaginatedResp struct {
	Items      any   `json:"items"`
	Page       int64 `json:"page"`
	PerPage    int64 `json:"per_page"`
	PageCount  int64 `json:"page_count"`
	TotalCount int64 `json:"total_count"`
}
