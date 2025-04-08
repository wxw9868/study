package database

import (
	"log"
	"os"
	"strings"
	"time"

	"study/StockCalculator/ledger/config"
	"study/StockCalculator/ledger/model"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

// 初始化数据库
func init() {
	var err error
	db, err = gorm.Open(sqlite.Open(config.Config().Database.DBPath), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "video_",                        // table name prefix, table for `User` would be `t_users`
			SingularTable: true,                            // use singular table name, table for `User` would be `user` with this option enabled
			NoLowerCase:   true,                            // skip the snake_casing of names
			NameReplacer:  strings.NewReplacer("ID", "id"), // use name replacer to change struct/field name before convert it to db name
		},
		PrepareStmt: true,
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:             time.Second,   // Slow SQL threshold
				LogLevel:                  logger.Silent, // Log level
				IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
				ParameterizedQueries:      true,          // Don't include params in the SQL log
				Colorful:                  false,         // Disable color
			},
		),
	})
	if err != nil {
		log.Fatalf("数据库链接失败: %s\n", err)
	}

	// 获取通用数据库对象 sql.DB ，然后使用其提供的功能
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	log.Println("数据库连接成功")

	// 迁移 schema
	if err = db.AutoMigrate(
		&model.Ledger{}, &model.AssetRecord{}, &model.Currency{},
	); err != nil {
		log.Fatalf("迁移 schema 失败: %s\n", err)
	}
}

func DB() *gorm.DB {
	return db
}
