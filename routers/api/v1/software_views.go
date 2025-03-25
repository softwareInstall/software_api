package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"software_api/service"
	"strconv"

	"software_api/dto"
	_ "software_api/models"
	"software_api/pkg/app"
	"software_api/pkg/e"
	"software_api/pkg/util"
)

func AddSoftware(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form dto.AddSoftwareDto
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	softwareService := service.SoftwareService{
		Name:     form.Name,
		Platform: form.Platform,
	}
	exists, err := softwareService.ExistByFilter()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_SOFTWARE_FAIL, nil)
		return
	}
	if exists {
		appG.Response(http.StatusOK, e.ERROR_EXIST_SOFTWARE, nil)
		return
	}

	err = softwareService.Add()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_SOFTWARE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func GetSoftwareList(c *gin.Context) {
	appG := app.Gin{C: c}
	paginationObj := util.GetPagination(c)
	software := service.SoftwareService{
		Offset: util.GetOffset(paginationObj),
		Limit:  paginationObj.PageSize,
		Filter: map[string]interface{}{"platform": util.IntDefaultQuery(c, "platform", 1)},
	}
	softwareList, err := software.GetSimpleInfo()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_SOFTWARES_FAIL, nil)
		return
	}
	newSoftwareList := software.ConvertSoftwareModelsToNameDts(c, softwareList)
	count, err := software.Count()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_COUNT_SOFTWARE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"items": newSoftwareList,
		"count": count,
	})
}

func GetSoftwareMenu(c *gin.Context) {
	appG := app.Gin{C: c}
	paginationObj := util.GetPagination(c)
	software := service.SoftwareService{
		Offset: util.GetOffset(paginationObj),
		Limit:  paginationObj.PageSize,
		Filter: map[string]interface{}{"platform": util.IntDefaultQuery(c, "platform", 1)},
	}
	softwareList, err := software.GetAll()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_SOFTWARES_FAIL, nil)
		return
	}
	newSoftwareList := software.ConvertSoftwareModelsToDts(c, softwareList)
	count, err := software.Count()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_COUNT_SOFTWARE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"items": newSoftwareList,
		"total": count,
	})
}

func EditSoftware(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form = dto.EditSoftwareDto{}
	)
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR_ID_NOT_INT, nil)
	}
	form.ID = id

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	softwareService := service.SoftwareService{
		ID:   form.ID,
		Name: form.Name,
	}

	exists, err := softwareService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_SOFTWARE_FAIL, nil)
		return
	}

	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_SOFTWARE, nil)
		return
	}

	err = softwareService.Edit()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EDIT_SOFTWARE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
