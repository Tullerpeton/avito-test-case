package models

import "time"

type Transaction struct {
	UserId        uint64    `json:"-"`
	TypeOperation string    `json:"type_operation"`
	Context       string    `json:"context"`
	CreatedAt     time.Time `json:"created_at"`
	Value         float64   `json:"value"`
}

type PaginatorTransaction struct {
	Page          int    `json:"page"`
	CountRecords  int    `json:"count_records"`
	SortField     string `json:"sort_field"`
	SortDirection string `json:"sort_direction"`
}

var (
	TransactionSumSort       = "sum"
	TransactionDateAddedSort = "date"

	PaginatorASC  = "asc"
	PaginatorDESC = "desc"
)
