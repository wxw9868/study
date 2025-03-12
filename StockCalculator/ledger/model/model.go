package model

import (
	"log"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Avatar   string `gorm:"type:varchar(255);not null;comment:头像"`
	Username string `gorm:"type:varchar(60);not null;comment:用户名"`
	Password string `gorm:"type:varchar(255);not null;comment:登录密码"`
	Nickname string `gorm:"type:varchar(60);default:系统用户;not null;comment:昵称"`
	Gender   int8   `gorm:"type:smallint;not null;default:0;comment:性别:0保密 1男 2女"`
	Phone    string `gorm:"type:varchar(60);not null;comment:手机"`
	Email    string `gorm:"type:varchar(60);not null;comment:邮箱"`
	Intro    string `gorm:"type:text;not null;comment:简介"`
	Status   int8   `gorm:"type:smallint;not null;comment:状态:1正常 2停用 3锁定 4拉黑"`
	Sorting  uint   `gorm:"type:smallint;not null;comment:排序"`
}

type UserLoginLog struct {
	gorm.Model
	InfoId        uint      `gorm:"comment:访问ID" json:"infoId"`
	LoginName     string    `gorm:"comment:登录账号" json:"loginName"`
	Ipaddr        string    `gorm:"comment:登录IP地址" json:"ipAddr"`
	LoginLocation string    `gorm:"comment:登录地点" json:"loginLocation"`
	Browser       string    `gorm:"comment:浏览器类型" json:"browser"`
	Os            string    `gorm:"comment:操作系统" json:"os"`
	Status        int8      `gorm:"comment:登录状态(1成功 2失败)" json:"status"`
	Msg           string    `gorm:"comment:提示消息" json:"msg"`
	LoginTime     time.Time `gorm:"comment:登录时间" json:"loginTime"`
	Module        string    `gorm:"comment:登录模块" json:"module"`
}

// 货币信息
type Currency struct {
	gorm.Model
	Code        string `gorm:"size:3;unique;not null;comment:货币代码(ISO 4217)"` // 货币代码，如 CNY, USD
	Name        string `gorm:"not null;comment:货币名称"`                         // 货币名称，如 人民币, 美元
	Symbol      string `gorm:"size:10;comment:货币符号"`                          // 货币符号，如 ¥, $
	CountryCode string `gorm:"size:2;comment:国家/地区代码(ISO 3166-1 alpha-2)"`    // 国家/地区代码
	CountryName string `gorm:"comment:国家/地区名称"`                               // 国家/地区名称
	IsActive    bool   `gorm:"default:true;comment:是否启用"`                     // 是否启用
}

// InitCurrencyData 初始化货币数据
func InitCurrencyData(db *gorm.DB) error {
	// 这里使用 GORM 的批量插入功能
	currencies := []Currency{
		{Code: "CNY", Name: "人民币", Symbol: "¥", CountryCode: "CN", CountryName: "中国", IsActive: true},
		{Code: "USD", Name: "美元", Symbol: "$", CountryCode: "US", CountryName: "美国", IsActive: true},
		{Code: "EUR", Name: "欧元", Symbol: "€", CountryCode: "EU", CountryName: "欧盟", IsActive: true},
		{Code: "JPY", Name: "日元", Symbol: "¥", CountryCode: "JP", CountryName: "日本", IsActive: true},
		{Code: "GBP", Name: "英镑", Symbol: "£", CountryCode: "GB", CountryName: "英国", IsActive: true},
		{Code: "HKD", Name: "港币", Symbol: "HK$", CountryCode: "HK", CountryName: "香港", IsActive: true},
	}

	for _, currency := range currencies {
		if err := db.FirstOrCreate(&currency, Currency{Code: currency.Code}).Error; err != nil {
			log.Printf("货币数据写入失败：%s", err.Error())
		}
	}
	return nil
}

type Ledger struct {
	gorm.Model
	UserId     uint      `gorm:"not null;comment:用户ID"` // 账本ID
	Name       string    `gorm:"not null;comment:帐本名称"`
	Amount     float64   `gorm:"not null;comment:资产金额"`
	Date       time.Time `gorm:"not null;comment:开始日期"`
	CurrencyId uint      `gorm:"not null;comment:币种ID"`
}

// 资产记录
type AssetRecord struct {
	gorm.Model
	LedgerId     uint      `gorm:"not null;comment:账本ID"`                              // 账本ID
	Amount       float64   `gorm:"not null;comment:资产金额"`                              // 资产金额
	Date         time.Time `gorm:"not null;comment:记录日期"`                              // 记录日期
	Profit       float64   `gorm:"type:decimal(10,2);not null;default 0;comment:累计收益"` // 累计收益
	AnnualReturn float64   `gorm:"not null;default 0;comment:年化收益率(%)"`                // 年化收益率(%)
	Notes        string    `gorm:"comment:备注"`                                         // 备注
}
