package util

import (
	"github.com/gin-gonic/gin"
	"software_api/middleware/pagination"
)

// GetPagination 从上下文中获取分页信息，并进行类型断言
func GetPagination(c *gin.Context) pagination.Pagination {
	paginationObj, exists := c.Get("pagination")
	if !exists {
		return pagination.Pagination{}
	}
	p, ok := paginationObj.(pagination.Pagination)
	if !ok {
		return pagination.Pagination{}
	}
	return p
}

// GetOffset get page parameters
func GetOffset(pagination pagination.Pagination) int {
	limit := 0
	page := pagination.Page
	if page > 0 {
		limit = (page - 1) * pagination.PageSize
	}

	return limit
}
