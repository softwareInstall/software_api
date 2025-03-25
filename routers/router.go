package routers

import (
	"github.com/gin-gonic/gin"
	"software_api/middleware/pagination"
	v1 "software_api/routers/api/v1"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	apiV1 := r.Group("/api/v1")
	{
		apiV1.Use(pagination.HandlePagination())
		software := apiV1.Group("/software")
		{
			software.GET("/menu", v1.GetSoftwareMenu)
			software.GET("/list", v1.GetSoftwareList)
			software.POST("", v1.AddSoftware)
			software.PUT("/:id", v1.EditSoftware)
			software.POST("/:id/version", v1.AddSoftwareVersion)
		}
		auth := apiV1.Group("/auth")
		{
			auth.POST("/login", v1.Login)
			auth.GET("/check_token", v1.CheckToken)
		}
		oss := apiV1.Group("/oss")
		{
			oss.POST("/upload", v1.Upload)
		}
	}

	return r
}
