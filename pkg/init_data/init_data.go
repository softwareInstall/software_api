package init_data

import (
	"software_api/service"
)

func Setup() {
	InitUser()
}

func InitUser() {
	loginService := service.LoginService{}
	loginService.CreateUserIfNotExists()
}
