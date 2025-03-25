package v1

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	_ "software_api/models"
	"software_api/pkg/app"
	"software_api/pkg/e"
	"software_api/service"
)

func Upload(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)

	file, _ := c.FormFile("file")
	log.Printf("上传文件名：%v", file.Filename)

	ossService := service.OssService{
		File: file,
		Ctx:  c,
	}
	fileInfo, err := ossService.UploadFile()
	if err != nil {
		log.Printf("上传文件失败：%v", err)
		appG.Response(http.StatusInternalServerError, e.ERROR_UPLOAD_FILE_FAIL, fileInfo)
	}

	appG.Response(http.StatusOK, e.SUCCESS, fileInfo)
}
