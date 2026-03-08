package pagination

import "github.com/1-AkM-0/empreGo-web/internal/validator"

type Filter struct {
	Page     int
	PageSize int
}

func (f Filter) Limit() int {
	return f.PageSize
}

func (f Filter) Offset() int {
	return (f.Page - 1) * f.PageSize
}

type Metadata struct {
	CurrentPage  int `json:"current_page,omitempty"`
	PageSize     int `json:"page_size,omitempty"`
	FirstPage    int `json:"first_page,omitempty"`
	LastPage     int `json:"last_page,omitempty"`
	TotalRecords int `json:"total_records,omitempty"`
}

func CalculateMetada(totalRecords, page, pageSize int) Metadata {
	if totalRecords == 0 {
		return Metadata{}
	}

	return Metadata{
		CurrentPage:  page,
		PageSize:     pageSize,
		FirstPage:    1,
		LastPage:     (totalRecords + pageSize - 1) / pageSize,
		TotalRecords: totalRecords,
	}
}

func ValidateFilter(v *validator.Validator, filter Filter) {
	v.Check(filter.Page > 0, "page", "deve ser maior que 0")
	v.Check(filter.Page <= 10_000_000, "page", "deve ser no máximo 10 milhões")
	v.Check(filter.PageSize > 0, "page_size", "deve ser maior que 0")
	v.Check(filter.PageSize <= 10, "page_size", "deve ser no máximo 10")
}
