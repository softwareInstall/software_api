package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"software_api/dto"
	_ "software_api/models"
	"software_api/pkg/app"
	"software_api/pkg/e"
	"software_api/service"
	"strconv"
)

func AddSoftwareVersion(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form dto.AddSoftwareVersionDto
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR_ID_NOT_INT, nil)
	}

	softwareVersionService := service.SoftwareVersionService{
		SoftwareID:   id,
		Version:      form.Version,
		Description:  form.Description,
		ObjectName:   form.ObjectName,
		OriginalName: form.OriginalName,
	}
	exists, err := softwareVersionService.ExistByVersion()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_SOFTWARE_VERSION_FAIL, nil)
		return
	}
	if exists {
		appG.Response(http.StatusOK, e.ERROR_EXIST_SOFTWARE_VERSION, nil)
		return
	}

	err = softwareVersionService.Add()
	if err != nil {
		fmt.Printf("%v", err)
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_SOFTWARE_VERSION_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
