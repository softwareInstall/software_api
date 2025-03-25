package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"software_api/pkg/setting"
)

var db *gorm.DB

type Model struct {
	ID        int            `gorm:"primary_key" json:"id"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type PrefixNamingStrategy struct {
	schema.NamingStrategy
	TablePrefix string
}

func (ns PrefixNamingStrategy) TableName(table string) string {
	return ns.TablePrefix + ns.NamingStrategy.TableName(table)
}

// Setup initializes the database instance
func Setup() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Name)

	namingStrategy := PrefixNamingStrategy{
		TablePrefix:    "", // 替换为您的前缀
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	}

	// 使用全局变量 db
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: namingStrategy,
	})
	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get generic database object: %v", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	err = db.AutoMigrate(&Software{}, &Version{}, &User{})
	// 自动迁移表结构
	if err != nil {
		log.Fatalf("models.Setup migration error: %v", err)
	}
}
