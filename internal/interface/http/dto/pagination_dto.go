package dto

type PaginationRequest struct {
	Page int `json:"page" validate:"required,min=1"`
	Size int `json:"size" validate:"required,min=1,max=100"`
}

type PaginationResponse struct {
	CurrentPage int   `json:"current_page"`
	Size        int   `json:"size"`
	TotalCount  int64 `json:"total_count"`
	TotalPages  int   `json:"total_pages"`
	HasNext     bool  `json:"has_next"`
	HasPrev     bool  `json:"has_prev"`
}
