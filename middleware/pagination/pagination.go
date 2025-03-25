package pagination

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Pagination struct {
	Page     int
	PageSize int
}

func HandlePagination() gin.HandlerFunc {
	return func(c *gin.Context) {
		pagination := Pagination{
			Page:     1,
			PageSize: 10,
		}

		// 获取查询参数
		pageNum := c.Query("page")
		pageSize := c.Query("page_size")

		if pageNum != "" {
			if p, err := strconv.Atoi(pageNum); err == nil && p > 0 {
				pagination.Page = p
			} else {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
				return
			}
		}

		if pageSize != "" {
			if ps, err := strconv.Atoi(pageSize); err == nil && ps > 0 {
				if ps > 100 { // 设置最大值
					ps = 100
				}
				pagination.PageSize = ps
			} else {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid page size"})
				return
			}
		}

		// 将分页信息存入上下文
		c.Set("pagination", pagination)
		c.Next()
	}
}
