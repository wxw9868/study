package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type Database struct {
	DB    *gorm.DB
	SqlDB *sql.DB
	Config
}

type Config struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
	Name string `mapstructure:"name"`
	User string `mapstructure:"user"`
	Pass string `mapstructure:"pass"`
}

func New() *Database {
	return &Database{
		Config: Config{
			Host: "127.0.0.1",
			Port: "3306",
			Name: "job",
			User: "root",
			Pass: "123456789",
		},
	}
}

func (conf *Database) CreateDatabase() error {
	source := fmt.Sprintf("%s:%s@tcp(%s:%s)/", conf.User, conf.Pass, conf.Host, conf.Port)
	sqlDB, err := sql.Open("mysql", source)
	if err != nil {
		return err
	}
	defer sqlDB.Close()

	_, err = sqlDB.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci", conf.Name))
	if err != nil {
		return err
	}
	return nil
}

func (conf *Database) ConnectDatabase() error {
	source := fmt.Sprintf("%s:%s@tcp(%s:%s)", conf.User, conf.Pass, conf.Host, conf.Port)
	dsn := fmt.Sprintf("%s/%s?charset=utf8mb4&parseTime=True&loc=Local", source, conf.Name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
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
		return err
	}
	conf.DB = db
	return nil
}

func (conf *Database) ConnectionPool() error {
	sqlDB, err := conf.DB.DB()
	if err != nil {
		return err
	}
	conf.SqlDB = sqlDB

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	return nil
}

func DB() *gorm.DB {
	return n.DB
}

func SqlDB() *sql.DB {
	return n.SqlDB
}

func Close() error {
	return n.SqlDB.Close()
}

var n = New()

func init() {
	if err := n.CreateDatabase(); err != nil {
		panic(err)
	}

	if err := n.ConnectDatabase(); err != nil {
		panic(err)
	}

	if err := n.ConnectionPool(); err != nil {
		panic(err)
	}

	err := DB().AutoMigrate(
		&Settings{},
		&University{},
		&Region{},

		&Order{},

		&Ad{},
		&Article{},

		&Coupon{},
		&CouponLog{},

		&User{},
		&UserLoginLog{},
		&UserBalance{},
		&UserBalanceLog{},
		&UserCertification{},

		&JobCategory{},
		&Job{},
		&JobDeliver{},
		&JobDeliveryMeter{},
		&JobPromotion{},
		&JobSettlementType{},
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("数据库初始化成功")
}
