package dto

import (
	"github.com/go-playground/validator/v10"
)

// SoftwareVersionBase 基础结构体
type SoftwareVersionBase struct {
	Version      string `json:"version" validate:"required,max=255"`
	Description  string `json:"description" validate:"required,max=255"`
	ObjectName   string `json:"object_name" validate:"required,max=255"`
	OriginalName string `json:"original_name" validate:"required,max=255"`
}

// AddSoftwareVersionDto 添加软件的数据传输对象
type AddSoftwareVersionDto struct {
	SoftwareVersionBase
}

// Validate 验证 AddSoftwareDto
func (add AddSoftwareVersionDto) Validate() error {
	validate := validator.New()
	// 可以在这里添加不同的验证器实例或规则
	return validate.Struct(add)
}
