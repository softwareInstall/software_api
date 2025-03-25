package models

import (
	"gorm.io/gorm"
	"log"
)

type User struct {
	Model

	Username string `json:"username"`
	Password string `json:"password"`
}

// GetUserPwByName 根据用户名获取密码
func GetUserPwByName(username string) (string, error) {
	var user User
	err := db.Select("password").Where(User{Username: username}).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return "", err
	}

	return user.Password, nil
}

// CreateUserIfNotExists 创建用户
func CreateUserIfNotExists(username, password string) {
	var count int64
	db.Model(&User{}).Where("username = ?", username).Count(&count)

	if count == 0 {
		user := User{
			Username: username,
			Password: password,
		}

		if err := db.Create(&user).Error; err != nil {
			log.Printf("创建用户失败: %v", err)
		} else {
			log.Printf("创建用户成功: %s", username)
		}
	} else {
		log.Printf("用户已经存在: %s", username)
	}
}
