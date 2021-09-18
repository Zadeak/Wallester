package repository

import (
	"errors"
	"gorm.io/gorm"
	"math"
)

type Pagination struct {
	Limit        int         `json:"limit,omitempty;query:limit"`
	Page         int         `json:"page,omitempty;query:page"`
	NextPage     int         `json:"next-page,omitempty;query:page"`
	PreviousPage int         `json:"previous-page,omitempty;query:page"`
	FirstPage    int         `json:"first_page,omitempty;query:page"`
	LastPage     int         `json:"last-page,omitempty;query:page"`
	Sort         string      `json:"sort,omitempty;query:sort"`
	TotalRows    int64       `json:"total_rows"`
	TotalPages   int         `json:"total_pages"`
	Rows         interface{} `json:"rows"`
}

func (p *Pagination) getFirstPage() int {
	return 1
}
func (p *Pagination) getLastPage() int {
	return p.TotalPages
}
func (p *Pagination) getNextPage() (int, error) {
	if p.Page+1 > p.TotalPages {
		return -1, errors.New("ERROR: current page is last")
	}
	return p.Page + 1, nil
}
func (p *Pagination) getPreviousPage() (int, error) {
	if p.Page-1 < 0 {
		return -1, errors.New("ERROR: current page is first page")
	}
	return p.Page - 1, nil
}

func (p *Pagination) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}
func (p *Pagination) GetLimit() int {
	if p.Limit == 0 {
		p.Limit = 10
	}
	return p.Limit
}

func (p *Pagination) GetPage() int {
	if p.Page == 0 {
		p.Page = 1
	}
	return p.Page
}

func (p *Pagination) GetSort() string {
	if p.Sort == "" {
		p.Sort = "Id ASC"
	}
	return p.Sort
}

func paginate(value interface{}, pagination *Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var totalRows int64
	db.Model(value).Count(&totalRows)
	pagination.TotalRows = totalRows
	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.GetLimit())))
	pagination.Page = pagination.GetPage()
	pagination.TotalPages = totalPages
	pagination.NextPage, _ = pagination.getNextPage()
	pagination.PreviousPage, _ = pagination.getPreviousPage()
	pagination.FirstPage = pagination.getFirstPage()
	pagination.LastPage = pagination.getLastPage()
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort())
	}
}
