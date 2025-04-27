package initdb

import "time"

// User 用户表
type User struct {
	GormModel
	Avatar    string    `gorm:"column:avatar;type:varchar(255);comment:头像"`
	Username  string    `gorm:"column:username;type:varchar(120);uniqueIndex;comment:用户名"`
	Nickname  string    `gorm:"column:nickname;type:varchar(120);uniqueIndex;comment:昵称"`
	Password  string    `gorm:"column:password;type:varchar(255);comment:密码"`
	Sex       int8      `gorm:"column:sex;not null;size:1;default:0;comment:性别: 0保密 1男 2女"`
	Birthday  time.Time `gorm:"column:birthday;comment:生日"`
	Mobile    string    `gorm:"column:mobile;type:varchar(20);uniqueIndex;comment:手机号"`
	Email     string    `gorm:"column:email;type:varchar(20);comment:邮箱"`
	Usertype  int8      `gorm:"column:usertype;size:1;not null;default:0;comment:用户类型: 1普通用户 2企业用户 3管理员"`
	WechatId  string    `gorm:"column:wechat_id;type:varchar(60);comment:微信号"`
	QQ        string    `gorm:"column:qq;type:varchar(60);comment:QQ"`
	Education int8      `gorm:"column:education;size:1;not null;default:0;comment:学历: 1小学 2初中 3高中 4大专 5本科 6研究生"`
	Degree    int8      `gorm:"column:degree;size:1;comment:学位: 1学士 2硕士 3博士"`
	Intro     string    `gorm:"column:intro;comment:简介"`
}

// UserLoginLog 用户登陆日志表
type UserLoginLog struct {
	GormModel
	InfoId        int64     `gorm:"comment:访问ID" json:"infoId"`
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

// UserBalance 用户余额表
type UserBalance struct {
	GormModel
	UserID    uint    `gorm:"column:user_id;not null;comment:用户ID"`
	Balance   float64 `gorm:"column:balance;type:decimal(10,2);not null;default:0;comment:账户余额"`
	GiveMoney float64 `gorm:"column:give_money;type:decimal(10,2);default:0;comment:满赠额度"`
}

// UserBalanceLog 用户余额日志表
type UserBalanceLog struct {
	GormModel
	UserID        uint      `gorm:"column:user_id;not null;comment:用户ID"`
	Amount        float64   `gorm:"column:amount;type:decimal(10,2);not null;comment:金额"`
	PaymentMethod string    `gorm:"column:payment_method;not null;comment:支付方式: alipay支付宝支付 wechatpay微信支付"`
	PaymentStatus int8      `gorm:"column:payment_status;not null;size:1;default:0;comment:支付状态: 1已支付 2待支付 3支付失败"`
	PaymentTime   time.Time `gorm:"column:payment_time;not null;comment:支付时间"`
	TransactionID string    `gorm:"column:transaction_id;type:varchar(255);comment:第三方平台交易流水号"`
	Status        int8      `gorm:"column:status;size:1;not null;default:0;comment:行为: 1充值 2消费 3推广 4提现 5岗位扣费"`
}

// UserCertification 用户认证表
type UserCertification struct {
	GormModel
	UserID          uint   `gorm:"column:user_id;not null;comment:用户ID"`
	Type            int8   `gorm:"column:type;size:1;comment:认证类型: 1企业用户 2个人用户"`
	Industry        uint   `gorm:"column:industry;comment:所属行业: job_category表ID"`
	CompanyName     string `gorm:"column:company_name;type:varchar(255);uniqueIndex;comment:公司名称"`
	CompanyLogo     string `gorm:"column:company_logo;type:varchar(255);comment:公司logo"`
	BusinessLicense string `gorm:"column:business_license;type:varchar(255);comment:营业执照"`
	Realname        string `gorm:"column:realname;type:varchar(255);comment:真实姓名"`
	Idcard          string `gorm:"column:idcard;type:varchar(30);uniqueIndex;comment:身份证号"`
	IdcardFront     string `gorm:"column:idcard_front;type:varchar(255);comment:身份证正面"`
	IdcardBack      string `gorm:"column:idcard_back;type:varchar(255);comment:身份证反面"`
	Province        uint   `gorm:"column:province;comment:省"`
	City            uint   `gorm:"column:city;comment:市"`
	District        uint   `gorm:"column:district;comment:区/县"`
	Address         string `gorm:"column:address;type:varchar(255);comment:详细地址"`
	Intro           string `gorm:"column:intro;type:text;comment:简介"`
	Status          int8   `gorm:"column:status;size:1;default:1;comment:状态: 1待审核 2审核通过 3审核失败"`
	Note            string `gorm:"column:note;type:text;comment:备注"`
}
