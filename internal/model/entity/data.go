package entity

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"sync"
)

// 数据表结构

// UserTable 用户表
type UserTable struct {
	Id       int    `gorm:"primary_key;auto_increment;not null"`
	Username string `gorm:"type:varchar(255);not null"`
	Password string `gorm:"type:varchar(255);not null"`
}

// MusicResourceTable 音乐资源表
type MusicResourceTable struct {
	Id        int    `gorm:"primary_key;auto_increment;not null"`
	Title     string `gorm:"type:varchar(255);not null"`
	MusicFile string `gorm:"type:varchar(255);not null"`
}

// TaskGroupTable 任务组表
type TaskGroupTable struct {
	Id   int    `gorm:"primary_key;auto_increment;not null"`
	Name string `gorm:"type:varchar(255);not null"`
}

// TaskTable 任务表
type TaskTable struct {
	Id        int    `gorm:"primary_key;auto_increment;not null"`
	MusicID   int    `gorm:"type:int;not null"`
	MusicName string `gorm:"type:varchar(255);not null"`
	GroupID   int    `gorm:"type:int;not null"`
	Loop      bool   `gorm:"type:bool;not null"`
}

var once sync.Once
var db *gorm.DB

// GetDatabase 获取数据库
func GetDatabase() *gorm.DB {
	once.Do(func() {
		db, _ = gorm.Open(sqlite.Open("./resource/database/data.db"), &gorm.Config{})
	})
	return db
}

// InitData 数据初始化
func InitData() {
	if GetDatabase().AutoMigrate(&UserTable{}, &MusicResourceTable{}, &TaskTable{}, &TaskGroupTable{}) != nil {
		panic("数据库初始化失败")
	}
	// 初始化管理员账号
	// 首先检查是否存在管理员账号
	var userTable UserTable
	if GetDatabase().Where("id = ?", 1).First(&userTable).RowsAffected == 0 {
		// 不存在管理员账号，创建一个
		GetDatabase().Create(&UserTable{Username: "admin", Password: "21232f297a57a5a743894a0e4a801fc3"})
		println("管理员账号已创建，用户名：admin，密码：admin")
	}
}
