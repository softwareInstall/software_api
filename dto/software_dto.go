package dto

import (
	"github.com/go-playground/validator/v10"
)

type Validatable interface {
	Validate() error
}

// SoftwareBase 基础结构体
type SoftwareBase struct {
	ID   int
	Name string `form:"name,omitempty" validate:"required,max=100"`
}

// AddSoftwareDto 添加软件的数据传输对象
type AddSoftwareDto struct {
	SoftwareBase
	Platform int `form:"platform,omitempty" validate:"required"`
}

// EditSoftwareDto 编辑软件的表单结构体
type EditSoftwareDto struct {
	SoftwareBase
}

// Validate 验证 AddSoftwareDto
func (add AddSoftwareDto) Validate() error {
	validate := validator.New()
	// 可以在这里添加不同的验证器实例或规则
	return validate.Struct(add)
}

// Validate 验证 EditSoftwareDto
func (edit EditSoftwareDto) Validate() error {
	validate := validator.New()
	// 可以在这里添加不同的验证器实例或规则
	return validate.Struct(edit)
}

type VersionDto struct {
	ID                int    `json:"id"`
	SoftwareID        int    `json:"software_id"`
	Version           string `json:"version"`
	ReleaseDate       string `json:"release_date"`
	Description       string `json:"description"`
	DownloadURL       string `json:"download_url"`
	LatestReleaseDate string `json:"latest_release_date"`
	LatestDownloadURL string `json:"latest_download_url"`
	LatestVersion     string `json:"latest_version"`
}

type SoftwareDto struct {
	ID                int          `json:"id"`
	Name              string       `json:"name"`
	LatestReleaseDate string       `json:"latest_release_date"`
	LatestDownloadURL string       `json:"latest_download_url"`
	LatestVersion     string       `json:"latest_version"`
	Versions          []VersionDto `json:"versions"`
}

type SoftwareNameIdDto struct {
	Value string `json:"value"`
	Label string `json:"label"`
}
