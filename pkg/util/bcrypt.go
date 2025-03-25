package util

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword 使用 bcrypt 对密码进行哈希
func HashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}

// CheckPassword 比对输入密码与哈希密码是否匹配
func CheckPassword(inputPassword, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(inputPassword))
	return err == nil
}
