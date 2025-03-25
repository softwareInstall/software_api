package util

import (
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"time"
)

// Setup Initialize the util
func Setup() {
}

func IntDefaultQuery(c *gin.Context, key string, defaultValue int) int {
	value := c.Query(key)
	if value == "" {
		return defaultValue
	}
	result, err := strconv.Atoi(value)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

// FormatTimePtrCustom 将带 "T" 的时间字符串格式化为 "YYYY-MM-DD HH:MM:SS"
func FormatTimePtrCustom(t *time.Time) string {
	const layout = "2006-01-02 15:04:05"
	if t == nil {
		return ""
	}
	return t.Format(layout)
}
