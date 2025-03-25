package util

import (
	"errors"
	"fmt"
	"software_api/pkg/setting"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// JwtSecret 应该从环境变量或安全配置中加载
var JwtSecret []byte

// Claims 定义了 JWT 的声明结构
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// init 函数用于初始化 JwtSecret
func init() {
	JwtSecret = []byte(setting.AppSetting.JwtSecret)
}

// GenerateToken 生成用于认证的 JWT Token
func GenerateToken(username string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
			Issuer:    "gin-blog",
			// 可选：添加 Subject 和 ID
			Subject: username,
			ID:      fmt.Sprintf("%d", time.Now().UnixNano()),
		},
	}

	// 使用 HS256 签名方法创建 Token
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 生成签名字符串
	token, err := tokenClaims.SignedString(JwtSecret)
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}

	return token, nil
}

// ParseToken 解析并验证 JWT Token
func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// 确保使用了预期的签名方法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return JwtSecret, nil
	})

	if err != nil {
		// 处理解析错误
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	// 验证 Token 是否有效
	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	// 断言 Claims 类型
	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	return claims, nil
}

// ValidateToken 是一个辅助函数，用于仅验证 Token 的有效性而不解析其内容
func ValidateToken(tokenString string) error {
	_, err := ParseToken(tokenString)
	return err
}
