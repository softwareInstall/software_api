package service

import (
	"log"
	"software_api/models"
	"software_api/pkg/setting"
	"software_api/pkg/util"
)

type LoginService struct {
	Username string
	Password string
	Token    string
	Filter   map[string]interface{}
}

func (l *LoginService) GetUser() (bool, error) {
	hashedPassword, err := models.GetUserPwByName(l.Username)
	if err != nil {
		return false, err
	}
	exists := util.CheckPassword(l.Password, hashedPassword)
	return exists, err
}

func (l *LoginService) CreateUserIfNotExists() {
	// 创建初始用户
	initialUsername := setting.UserSetting.Username
	initialPassword := setting.UserSetting.Password
	hashedPassword, err := util.HashPassword(initialPassword)
	if err != nil {
		log.Printf("CreateUserIfNotExists, fail: %v", err)
	}
	models.CreateUserIfNotExists(initialUsername, hashedPassword)
}

func (l *LoginService) CheckToken() bool {
	err := util.ValidateToken(l.Token)
	if err != nil {
		log.Printf("CheckToken fail': %v", err)
		return false
	}
	return true
}
