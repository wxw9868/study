package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/wxw9868/checkin/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var db *gorm.DB
var sqlDB *sql.DB
var err error

func InitDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", "root", "123456789", "127.0.0.1", "3306", "checkin")
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "pt_", // 表名前缀，`User`表为`t_users`
			SingularTable: true,  // 使用单数表名，启用该选项后，`User` 表将是`user`
		},
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold: time.Second, // 慢 SQL 阈值
				LogLevel:      logger.Info, // Log level
				Colorful:      false,       // 禁用彩色打印
			},
		),
	})
	if err != nil {
		panic(err)
	}

	sqlDB, err = db.DB()
	if err != nil {
		panic(err)
	}

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	log.Println("数据库初始化成功")
}

func DB() *gorm.DB {
	return db
}

func CloseDB() error {
	return sqlDB.Close()
}

func InitDatabase() {
	sqlDB, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/", "root", "123456789", "127.0.0.1", "3306"))
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()

	// CREATE DATABASE my_database_name CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
	// 创建数据库的SQL命令
	createDBSQL := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci", "checkin")
	// 执行SQL命令创建数据库
	_, err = sqlDB.Exec(createDBSQL)
	if err != nil {
		log.Fatal(err)
	}

	user := models.User{
		CenterLat: 39.9042,
		CenterLng: 116.4074,
		Radius:    100,
	}
	if err := db.FirstOrCreate(&user).Error; err != nil {
		fmt.Printf("创建用户失败：%v\n", err)
		panic(err)
	}
	fmt.Printf("创建用户成功：%v\n", user)
}
