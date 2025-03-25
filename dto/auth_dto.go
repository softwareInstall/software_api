package dto

import (
	"github.com/go-playground/validator/v10"
)

// LoginDto 结构体
type LoginDto struct {
	Username string `json:"username" validate:"required,max=255"`
	Password string `json:"password" validate:"required,max=255"`
}

// Validate 验证 AddSoftwareDto
func (add LoginDto) Validate() error {
	validate := validator.New()
	// 可以在这里添加不同的验证器实例或规则
	return validate.Struct(add)
}
