package model

import (
	"math"
	"strconv"
)

type PageTable struct {
	Fields    string
	Table     string
	Where     string
	PageIndex int
	PageSize  int
	Order     string
	GroupBy   string
}

type PageData struct {
	PageIndex  int
	PageSize   int
	PageNumber int
	Count      int
	Data       interface{}
}

const PageSize = 10

func GetPageIndex(pi string) int {
	pageIndex, err := strconv.Atoi(pi)
	if err != nil {
		pageIndex = 1
	}
	return pageIndex
}

func GetPageSize(ps string) int {
	pageSize, err := strconv.Atoi(ps)
	if err != nil {
		pageSize = PageSize
	}
	if pageSize > 100 {
		//防止大量数据查询
		pageSize = 100
	}
	return pageSize
}

type Count struct {
	count int
}

func GetPageNumber(count, pageSize int) int {
	return int(math.Ceil(float64(count) / float64(pageSize)))
}
