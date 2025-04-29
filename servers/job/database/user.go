package database

import "time"

// User 用户表
type User struct {
	GormModel
	Avatar      string     `gorm:"column:avatar;type:varchar(128);comment:头像"`
	Username    string     `gorm:"column:username;type:varchar(64);uniqueIndex;not null;comment:用户名"`
	Nickname    string     `gorm:"column:nickname;type:varchar(64);index;comment:昵称"` // 添加索引
	Password    string     `gorm:"column:password;type:char(60);not null;comment:密码（加密存储）"`
	Sex         int8       `gorm:"column:sex;type:tinyint(1);index;not null;default:0;comment:性别: 0保密 1男 2女"` // 添加索引
	Birthday    *time.Time `gorm:"column:birthday;type:date;index;comment:生日"`                                // 添加索引
	Mobile      string     `gorm:"column:mobile;type:char(11);uniqueIndex;not null;comment:手机号"`
	Email       string     `gorm:"column:email;type:varchar(128);uniqueIndex;not null;comment:邮箱"`
	WechatId    string     `gorm:"column:wechat_id;type:varchar(32);index;comment:微信号"`
	QQ          string     `gorm:"column:qq;type:varchar(20);index;comment:QQ号"`
	Education   int8       `gorm:"column:education;type:tinyint(1);index;not null;default:0;comment:学历: 1小学 2初中 3高中 4大专 5本科 6研究生"` // 添加索引
	Degree      int8       `gorm:"column:degree;type:tinyint(1);index;not null;default:0;comment:学位: 1学士 2硕士 3博士"`                 // 添加索引
	Usertype    int8       `gorm:"column:usertype;type:tinyint(1);index;not null;default:1;comment:用户类型: 1普通用户 2企业用户 3管理员"`
	Status      int8       `gorm:"column:status;type:tinyint(1);index;not null;default:1;comment:用户状态: 1正常 2拉黑"`
	Intro       string     `gorm:"column:intro;type:text;comment:简介"`
	LastLoginIP string     `gorm:"column:last_login_ip;type:varchar(45);comment:最后登录IP"`
	LastLoginAt *time.Time `gorm:"column:last_login_at;index;comment:最后登录时间"` // 添加索引
}

// UserLoginLog 用户登录日志表
type UserLoginLog struct {
	GormModel
	UserID        uint      `gorm:"column:user_id;index;not null;comment:用户ID"`
	LoginName     string    `gorm:"column:login_name;type:varchar(64);index;not null;comment:登录账号"` // Added index
	IPAddr        string    `gorm:"column:ip_addr;type:varchar(45);index;not null;comment:登录IP地址"`  // Added index
	LoginLocation string    `gorm:"column:login_location;type:varchar(128);comment:登录地点"`
	Browser       string    `gorm:"column:browser;type:varchar(64);comment:浏览器类型"`
	OS            string    `gorm:"column:os;type:varchar(64);comment:操作系统"`
	Status        int8      `gorm:"column:status;type:tinyint(1);index;not null;default:1;comment:登录状态: 1成功 2失败"` // Added index
	Msg           string    `gorm:"column:msg;type:varchar(128);comment:提示消息"`
	LoginTime     time.Time `gorm:"column:login_time;index;not null;comment:登录时间"` // Added index
	Module        string    `gorm:"column:module;type:varchar(32);comment:登录模块"`
}

// UserBalance 用户余额表
type UserBalance struct {
	GormModel
	UserID    uint    `gorm:"column:user_id;uniqueIndex;not null;comment:用户ID"`
	Balance   float64 `gorm:"column:balance;type:decimal(10,2);not null;default:0.00;comment:账户余额"`
	GiveMoney float64 `gorm:"column:give_money;type:decimal(10,2);default:0.00;comment:赠送金额"`
}

// UserBalanceLog 用户余额日志表
type UserBalanceLog struct {
	GormModel
	UserID        uint      `gorm:"column:user_id;index;not null;comment:用户ID"`
	Action        int8      `gorm:"column:action;type:tinyint(1);index;not null;default:0;comment:行为: 1充值 2消费 3推广 4提现 5岗位扣费"` // Added index
	Amount        float64   `gorm:"column:amount;type:decimal(10,2);not null;comment:金额"`
	PaymentMethod string    `gorm:"column:payment_method;type:varchar(32);index;not null;comment:支付方式: alipay支付宝支付 wechatpay微信支付"` // Added index
	PaymentTime   time.Time `gorm:"column:payment_time;index;not null;comment:支付时间"`                                               // Added index
	PaymentStatus int8      `gorm:"column:payment_status;type:tinyint(1);index;not null;default:2;comment:支付状态: 1已支付 2待支付 3支付失败"`  // Added index, default 2
	TransactionID string    `gorm:"column:transaction_id;type:varchar(64);uniqueIndex;comment:第三方平台交易流水号"`                         // 改为唯一索引
	Note          string    `gorm:"column:note;type:varchar(255);comment:备注"`
	// 添加复合索引，提高查询效率
	_ int `gorm:"index:idx_user_action_time,columns:user_id,action,payment_time"`
}

// UserCertification 用户认证表
type UserCertification struct {
	GormModel
	UserID          uint       `gorm:"column:user_id;uniqueIndex;not null;comment:用户ID"`
	Type            int8       `gorm:"column:type;type:tinyint(1);index;not null;default:2;comment:认证类型: 1企业用户 2个人用户"` // 添加默认值
	Industry        uint       `gorm:"column:industry;index;not null;comment:所属行业"`
	CompanyName     string     `gorm:"column:company_name;type:varchar(100);uniqueIndex;comment:公司名称"` // 增加字段长度
	CompanyLogo     string     `gorm:"column:company_logo;type:varchar(200);comment:公司logo"`           // 增加字段长度
	BusinessLicense string     `gorm:"column:business_license;type:varchar(200);comment:营业执照"`         // 增加字段长度
	Realname        string     `gorm:"column:realname;type:varchar(32);index;comment:真实姓名"`            // 添加索引
	Idcard          string     `gorm:"column:idcard;type:char(18);uniqueIndex;comment:身份证号"`
	IdcardFront     string     `gorm:"column:idcard_front;type:varchar(200);comment:身份证正面"` // 增加字段长度
	IdcardBack      string     `gorm:"column:idcard_back;type:varchar(200);comment:身份证反面"`  // 增加字段长度
	Intro           string     `gorm:"column:intro;type:text;comment:简介"`
	Status          int8       `gorm:"column:status;type:tinyint(1);index;not null;default:0;comment:状态: 0待审核 1通过 2失败"` // 添加not null
	Note            string     `gorm:"column:note;type:text;comment:备注"`
	ApproveTime     *time.Time `gorm:"column:approve_time;index;comment:审核时间"` // 添加审核时间字段
	ApproveBy       uint       `gorm:"column:approve_by;index;comment:审核人ID"`  // 添加审核人字段
}
