package model

type Empty struct{}

type BasicResponse struct {
	Message string `json:"message"`
}

type Cursor struct {
	After  string `json:"after,omitempty" query:"after,omitempty" example:""`
	Before string `json:"before,omitempty" query:"before,omitempty" example:""`
}

type Sort struct {
	Key   string `query:"key" example:""`
	Order string `query:"order" validate:"omitempty,oneof=asc desc" example:"asc"`
}

type PaginationRequest struct {
	Cursor Cursor   `query:"cursor"`
	Keys   []string `query:"-"`
	Limit  int      `query:"limit"`
	Sort   Sort     `query:"sort"`
}

type PaginationResponse struct {
	Cursor Cursor `json:"cursor"`
	Total  int64  `json:"total"`
}
