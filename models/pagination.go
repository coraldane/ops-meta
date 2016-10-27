package models

import (
	"net/http"
	"strconv"
)

var (
	PAGES_PER_TIME    = 5
	PAGE_SIZE_DEFAULT = 10
)

type PageInfo struct {
	PageSize  int `form:"pageSize"`
	PageIndex int `form:"pageIndex"`
	RowCount  int64
	Sort      string `form:"sort"`
	Order     string `form:"order"`
}

func (this *PageInfo) SetPageSize(size int) {
	this.PageSize = size
}

func (this *PageInfo) SetRowCount(total int64) {
	this.RowCount = total
}

func (this *PageInfo) GetStartIndex() int {
	return (this.PageIndex - 1) * this.PageSize
}

func NewPageInfo(request *http.Request) *PageInfo {
	pageIndex, _ := strconv.Atoi(request.FormValue("pageIndex"))
	pageSize, _ := strconv.Atoi(request.FormValue("pageSize"))
	if 0 >= pageIndex {
		pageIndex = 1
	}
	p := PageInfo{PageSize: pageSize, PageIndex: pageIndex}
	return &p
}
