package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"software_api/dto"
	_ "software_api/models"
	"software_api/pkg/app"
	"software_api/pkg/e"
	"software_api/pkg/util"
	"software_api/service"
)

func Login(c *gin.Context) {
	var (
		appG  = app.Gin{C: c}
		form  dto.LoginDto
		token string
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	authService := service.LoginService{
		Username: form.Username,
		Password: form.Password,
	}
	exists, err := authService.GetUser()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH, nil)
		return
	}

	if exists {
		token, err = util.GenerateToken(form.Username)
	} else {
		token = ""
	}

	appG.Response(http.StatusOK, e.SUCCESS, token)
}

func CheckToken(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)

	authService := service.LoginService{
		Token: c.Query("token"),
	}
	isValid := authService.CheckToken()

	appG.Response(http.StatusOK, e.SUCCESS, isValid)
}
