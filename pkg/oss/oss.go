package oss

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"io"
	"log"
	"time"

	"software_api/pkg/setting"
)

// ManagerOss 是一个封装了 MinIO 客户端和相关操作的结构体
type ManagerOss struct {
	client        *minio.Client
	defaultBucket string
}

var ManagerOssObj *ManagerOss

// NewManager 初始化 Manager 并返回一个实例
func NewManager(endpoint, accessKeyID, secretAccessKey string) (*ManagerOss, error) {
	// 创建 MinIO 客户端
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: false, // 使用 HTTPS
	})
	if err != nil {
		return nil, fmt.Errorf("无法连接到 MinIO 服务器: %w", err)
	}
	fmt.Println("成功连接到 MinIO 服务器")

	// 返回 OssManager 实例
	return &ManagerOss{client: client, defaultBucket: setting.OssSetting.BucketName}, nil
}

// CreateBucketIfNotExists 创建桶（如果不存在）
func (m *ManagerOss) CreateBucketIfNotExists(ctx context.Context, bucketName string) error {
	// 检查桶是否存在
	exists, err := m.client.BucketExists(ctx, bucketName)
	if err != nil {
		return fmt.Errorf("检查桶是否存在失败: %w", err)
	}
	if exists {
		fmt.Printf("桶 %s 已存在", bucketName)
		return nil
	}

	// 创建桶
	if err := m.client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{}); err != nil {
		return fmt.Errorf("创建桶失败: %w", err)
	}
	fmt.Printf("桶 %s 创建成功", bucketName)
	return nil
}

// UploadFile 上传文件到 MinIO
func (m *ManagerOss) UploadFile(ctx context.Context, objectName string, file io.Reader, fileSize int64) error {
	// 上传文件
	info, err := m.client.PutObject(ctx, m.defaultBucket, objectName, file, fileSize, minio.PutObjectOptions{})
	if err != nil {
		return fmt.Errorf("上传文件失败: %w", err)
	}
	fmt.Printf("文件上传成功: %s (%d 字节)", objectName, info.Size)
	return nil
}

// DownloadFile 从 MinIO 下载文件
func (m *ManagerOss) DownloadFile(ctx context.Context, objectName, downloadFilePath string) error {
	// 下载文件
	err := m.client.FGetObject(ctx, m.defaultBucket, objectName, downloadFilePath, minio.GetObjectOptions{})
	if err != nil {
		return fmt.Errorf("下载文件失败: %w", err)
	}
	fmt.Printf("文件下载成功: %s (%d 字节)", downloadFilePath)
	return nil
}

// DeleteObject 删除 MinIO 中的对象
func (m *ManagerOss) DeleteObject(ctx context.Context, objectName string) error {
	// 删除对象
	if err := m.client.RemoveObject(ctx, m.defaultBucket, objectName, minio.RemoveObjectOptions{}); err != nil {
		return fmt.Errorf("删除对象失败: %w", err)
	}
	fmt.Printf("对象删除成功: %s", objectName)
	return nil
}

// SetDefaultBucket 不建议使用配置之外的桶，我希望这些配置能在固定配置中可见，而不是动态的,能防止很多其他情况
func (m *ManagerOss) SetDefaultBucket(bucketName string) error {
	if bucketName == "" {
		return fmt.Errorf("bucket name cannot be empty")
	}
	m.defaultBucket = bucketName
	return nil
}

// GetFileUrl 获取文件的零时下载地址
func (m *ManagerOss) GetFileUrl(ctx context.Context, objectName string) (string, error) {
	fileUrl, err := m.client.PresignedGetObject(ctx, m.defaultBucket, objectName, setting.OssSetting.FileExpires*time.Minute, nil)
	if err != nil {
		log.Printf("生成预签名链接失败: %v", err)
	} else {
		log.Printf("预签名下载链接:: %v", fileUrl)
	}
	return fileUrl.String(), err
}

// Setup 初始化
func Setup() {
	// 初始化 Manager
	var err error
	ManagerOssObj, err = NewManager(setting.OssSetting.Endpoint, setting.OssSetting.AccessKeyID, setting.OssSetting.SecretAccessKey)
	if err != nil {
		log.Fatalf("初始化 OssManager 失败: %v", err)
	}
	ctx := context.Background()
	// 创建桶（如果不存在）
	if err = ManagerOssObj.CreateBucketIfNotExists(ctx, setting.OssSetting.BucketName); err != nil {
		log.Fatalf("创建桶失败: %v", err)
	}
}
