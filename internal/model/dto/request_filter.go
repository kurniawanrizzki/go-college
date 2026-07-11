package dto

type CollegeFilter struct {
	NIM      string `json:"nim"`
	Name     string `json:"name"`
	Semester int64  `json:"semester"`
	SortBy   string `json:"sort_by"`
	SortDir  string `json:"sort_dir"`
	Page     int64  `json:"page"`
	PerPage  int64  `json:"per_page"`
	Limit    int64  `json:"-"`
	Offset   int64  `json:"-"`
}

type CourseFilter struct {
	Code    string `json:"code"`
	Name    string `json:"name"`
	SKS     int64  `json:"sks"`
	SortBy  string `json:"sort_by"`
	SortDir string `json:"sort_dir"`
	Page    int64  `json:"page"`
	PerPage int64  `json:"per_page"`
	Limit   int64  `json:"-"`
	Offset  int64  `json:"-"`
}
