package initdb

import (
	"time"

	"gorm.io/gorm"
)

type GormModel struct {
	ID        uint           `gorm:"primarykey"`
	CreatedAt time.Time      `gorm:"column:created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index"`
}

// Ad [...]
type Ad struct {
	GormModel
	AdName    string    `gorm:"column:ad_name;type:varchar(255);not null;comment:广告名称"`   // 广告名称
	AdURL     string    `gorm:"column:ad_url;type:varchar(255);not null;comment:链接地址"`    // 链接地址
	ImgURL    string    `gorm:"column:img_url;type:varchar(255);not null;comment:图片地址"`   // 图片地址
	StartTime time.Time `gorm:"column:start_time;not null;comment:开始时间·"`                 // 开始时间·
	EndTime   time.Time `gorm:"column:end_time;not null;comment:结束时间"`                    // 结束时间
	IsShow    int8      `gorm:"column:is_show;not null;default:1;comment:是否显示: 1显示 2不显示"` // 是否显示：1显示；2不显示
	IsFree    int8      `gorm:"column:is_free;not null;comment:是否免费: 1是 2不是"`             // 是否免费：1是；2不是
	Sort      int       `gorm:"column:sort;not null;default:0;comment:排序"`                // 排序
	AdType    int8      `gorm:"column:ad_type;not null;comment:广告类型: 1banner广告 2普通广告"`    // 广告类型：1banner广告；2普通广告
}

// Article [...]
type Article struct {
	GormModel
	Title   string `gorm:"column:title;type:varchar(255);not null;comment:文章标题"` // 标题
	Author  string `gorm:"column:author;type:varchar(255);not null;comment:文章作者"`
	Content string `gorm:"column:content;type:longtext;not null;comment:文章内容"` // 内容
}

// Coupon [...]
type Coupon struct {
	GormModel
	Name       string  `gorm:"column:name;type:varchar(255);comment:优惠券名称"`                                // 优惠券名称
	Level      int16   `gorm:"column:level;type:smallint(6);comment:优惠券等级"`                                // 优惠券等级
	FullAmount float64 `gorm:"column:full_amount;type:decimal(10,2);not null;comment:金额"`                  // 金额
	SendAmount float64 `gorm:"column:send_amount;type:decimal(10,2);not null;comment:满多少金额送多少: 例如满100送10"` // 满多少金额送多少：例如满100送10
	IsUse      int8    `gorm:"column:is_use;not null;default:1;comment:是否使用: 1已使用 2没使用"`                   // 是否使用：1已使用；2没使用
	Desc       string  `gorm:"column:desc;type:varchar(255);comment:优惠券介绍"`                                // 优惠券介绍
}

// CouponLog [...]
type CouponLog struct {
	GormModel
	UserID   uint `gorm:"column:user_id;not null;comment:用户ID"`
	OrderID  uint `gorm:"column:order_id;not null;comment:订单ID"`
	CouponID uint `gorm:"column:coupon_id;not null;comment:优惠券ID"`
}

// Job [...]
type Job struct {
	GormModel
	UserID       uint      `gorm:"column:user_id;not null;comment:用户ID"`                     // 用户id
	CatID        uint      `gorm:"column:cat_id;not null;comment:职位分类"`                      // 职位分类
	Name         string    `gorm:"column:name;type:varchar(255);not null;comment:职位名称"`      // 职位名称
	Number       uint      `gorm:"column:number;not null;comment:招聘人数"`                      // 招聘人数
	IsDiscuss    int8      `gorm:"column:is_discuss;not null;default:1;comment:是否面议: 1否 2是"` // 是否面议：1否；2是
	MaxSalary    float64   `gorm:"column:max_salary;not null;comment:最高薪资"`                  // 最高薪资
	MinSalary    float64   `gorm:"column:min_salary;not null;comment:最低薪资"`                  // 最低薪资
	SettlementID uint      `gorm:"column:settlement_id;not null;comment:薪资结算方式ID"`           // 薪资结算方式ID
	StartTime    time.Time `gorm:"column:start_time;not null;comment:开始工作时间"`                // 开始工作时间
	EndTime      time.Time `gorm:"column:end_time;not null;comment:结束工作时间"`                  // 结束工作时间
	Description  string    `gorm:"column:description;type:longtext;not null;comment:职位描述"`   // 职位描述
	Province     int       `gorm:"column:province;type:int(11);not null;comment:工作所在地：省"`    // 工作所在地：省
	City         int       `gorm:"column:city;type:int(11);not null;comment:工作所在地：市"`        // 工作所在地：市
	District     int       `gorm:"column:district;type:int(11);not null;comment:工作所在地：区/县"`  // 工作所在地：区/县
	Address      string    `gorm:"column:address;type:varchar(255);not null;comment:工作详细地址"` // 工作详细地址
	Liaison      string    `gorm:"column:liaison;type:varchar(255);not null;comment:联系人"`    // 联系人
	Mobile       string    `gorm:"column:mobile;type:varchar(255);not null;comment:手机号"`     // 手机号
	WechatId     string    `gorm:"column:wechat_id;type:varchar(255);comment:微信号"`           // 微信号
	Status       int8      `gorm:"column:status;not null;comment:审核状态: 1待审核 2审核通过 3审核失败"`    // 审核状态：1待审核，2审核通过；3审核失败
	Comment      string    `gorm:"column:comment;type:text;comment:备注"`                      // 备注
	IsShow       int8      `gorm:"column:is_show;not null;default:2;comment:是否上线: 1上线 2下线"`  // 是否上线：1上线；2下线
	HideTime     time.Time `gorm:"column:hide_time;comment:下线时间"`                            // 下线时间
	ShowTime     time.Time `gorm:"column:show_time;comment:上线时间"`                            // 上线时间
	IsTop        int8      `gorm:"column:is_top;not null;default:2;comment:是否置顶: 1置顶 2取消"`   // 是否置顶：1置顶；2取消
	Sort         int       `gorm:"column:sort;type:int(11);not null;default:0;comment:排序"`   // 排序
}

// JobCategory [...]
type JobCategory struct {
	GormModel
	ParentID  uint   `gorm:"column:parent_id;not null;default:0;comment:父ID"`             // 父ID
	Name      string `gorm:"column:name;type:varchar(255);not null;comment:职位分类名称"`       // 职位类别名称
	Icon      string `gorm:"column:icon;type:varchar(255);comment:职位分类图标"`                // 分类图片路径
	Image     string `gorm:"column:image;type:varchar(255);comment:分类图片路径"`               // 分类图片路径
	Sort      int    `gorm:"column:sort;not null;comment:排序"`                             // 排序
	IsShow    int8   `gorm:"column:is_show;not null;default:0;comment:是否显示: 2不显示 1显示"`    // 是否显示：2不显示；1显示
	Recommend int8   `gorm:"column:recommend;not null;default:0;comment:金刚位显示: 2不显示 1显示"` // 金刚位显示：2不显示；1显示
}

// JobDeliver [...]
type JobDeliver struct {
	GormModel
	JobID       uint `gorm:"column:job_id;not null;comment:岗位ID"`
	HireUserID  uint `gorm:"column:hire_user_id;not null;comment:发布岗位的用户ID"`
	ApplyUserID uint `gorm:"column:apply_user_id;not null;comment:申请岗位的用户ID"`
	Status      int8 `gorm:"column:status;not null;comment:状态: 1被查看 2已录取 3已拒绝 4待处理"`
}

// JobDeliveryMeter [...]
type JobDeliveryMeter struct {
	GormModel
	DeliveryID uint `gorm:"column:delivery_id;not null;comment:用户投递表ID"`         // 用户投递表ID
	PayStatus  int8 `gorm:"column:pay_status;not null;comment:扣费状态: 1已扣费 2没有扣费"` // 扣费状态：1已扣费；2没有扣费
}

// JobPromotion [...]
type JobPromotion struct {
	GormModel
	UserID          uint      `gorm:"column:user_id;not null;comment:用户ID"`
	JobID           uint      `gorm:"column:job_id;not null;comment:岗位ID"`
	PromotionFee    float64   `gorm:"column:promotion_fee;type:decimal(10,2);not null;comment:推广费用"`      // 推广费用
	PromotionType   int8      `gorm:"column:promotion_type;not null;comment:推广类型: 1banner广告 2普通广告 3列表置顶"` // 推广类型：1banner广告；2普通广告；3列表置顶
	PromotionStatus int8      `gorm:"column:promotion_status;not null;comment:申请状态: 1待审核 2通过 3不通过"`       // 申请状态：1待审核；2通过；3不通过
	StartTime       time.Time `gorm:"column:start_time;type:datetime;not null;comment:开始时间"`              // 开始时间
	EndTime         time.Time `gorm:"column:end_time;type:datetime;not null;comment:结束时间"`                // 结束时间
}

// JobSettlementType [...]
type JobSettlementType struct {
	GormModel
	SettlementName string `gorm:"column:settlement_name;type:varchar(30);not null;comment:结算名称"` // 结算名称
	Description    string `gorm:"column:description;type:varchar(255);comment:说明"`               // 说明
	Sort           int    `gorm:"column:sort;type:int(11);not null;default:1;comment:排序"`        // 排序
}

// Order [...]
type Order struct {
	GormModel
	UserID         uint      `gorm:"column:user_id;not null;comment:用户ID"`
	OrderSn        string    `gorm:"column:order_sn;not null;comment:订单号"`                                   // 订单号
	OrderStatus    int8      `gorm:"column:order_status;not null;comment:订单状态: 1已确认 2已取消 3已完成 4已作废"`         // 订单状态：1已确认，2已取消，3已完成，4已作废
	JobPromotionID uint      `gorm:"column:job_promotion_id;default:0;comment:职位推广ID"`                       // 职位推广ID
	OrderAmount    float64   `gorm:"column:order_amount;type:decimal(10,2);not null;comment:订单金额"`           // 订单金额
	OrderType      int8      `gorm:"column:order_type;not null;comment:订单类型: 1充值 2消费 3推广 4提现"`               // 订单类型：1充值；2消费；3推广；4提现
	OrderDesc      string    `gorm:"column:order_desc;comment:订单说明"`                                         // 订单说明
	PaymentMethod  string    `gorm:"column:payment_method;not null;comment:支付方式: alipay支付宝支付 wechatpay微信支付"` // 支付方式：alipay支付宝支付；wechatpay微信支付
	PaymentStatus  int8      `gorm:"column:payment_status;not null;default:0;comment:支付状态: 1已支付 2待支付 3支付失败"` // 支付状态：1已支付；2待支付；3支付失败；
	PaymentTime    time.Time `gorm:"column:payment_time;not null;comment:支付时间"`                              // 支付时间
	TransactionID  string    `gorm:"column:transaction_id;type:varchar(255);comment:第三方平台交易流水号"`             // 第三方平台交易流水号
}

// Region [...]
type Region struct {
	GormModel
	Name     string `gorm:"column:name;type:longtext;comment:"`
	Level    int8   `gorm:"column:level;type:tinyint(4);default:0;comment:"`
	ParentID int    `gorm:"index:parent_id;column:parent_id;type:int(10);comment:"`
	RegionID int    `gorm:"column:region_id;type:int(11);comment:"`
}

// Settings [...]
type Settings struct {
	GormModel
	WechatShow      int8    `gorm:"column:wechat_show;not null;default:2;comment:微信显示: 1展示 2隐藏"`            // 微信显示：1展示；2隐藏
	Price           float64 `gorm:"column:price;type:decimal(10,2);not null;default:0;comment:投递一次岗位的收费价格"` // 投递一次岗位的收费价格
	CustomerService string  `gorm:"column:customer_service;comment:客户服务"`                                   // 客户服务
}

// University 全国高等学校名单表
type University struct {
	GormModel
	SchoolName          string `gorm:"column:school_name;type:varchar(60);uniqueIndex;comment:学校名称"`
	SchoolIdentifier    string `gorm:"column:school_identifier;type:varchar(20);uniqueIndex;comment:学校标识码"`
	CompetentDepartment string `gorm:"column:competent_department;type:varchar(60);comment:主管部门"`
	Location            string `gorm:"column:location;type:varchar(60);comment:所在地"`
	SchoolLevel         string `gorm:"column:location;type:varchar(10);comment:办学层次"`
	Note                string `gorm:"column:note;type:varchar(255);comment:备注"`
}

// User 用户表
type User struct {
	GormModel
	Username  string `gorm:"column:username;type:varchar(120);uniqueIndex;comment:用户名"`
	Password  string `gorm:"column:password;type:varchar(255);comment:密码"`
	Mobile    string `gorm:"column:mobile;type:varchar(20);uniqueIndex;comment:手机号"`
	Email     string `gorm:"column:email;type:varchar(20);comment:邮箱"`
	Avatar    string `gorm:"column:avatar;type:varchar(255);comment:头像"`
	Usertype  int8   `gorm:"column:usertype;not null;size:1;default:0;comment:用户类型: 1普通用户 2企业用户 3管理员"`
	Sex       int8   `gorm:"column:sex;not null;size:1;default:0;comment:性别: 0保密 1男 2女"`
	WechatId  string `gorm:"column:wechat_id;type:varchar(60);comment:微信号"`
	QQ        string `gorm:"column:qq;type:varchar(60);comment:QQ"`
	Education int8   `gorm:"column:education;size:1;comment:学历: 1小学 2初中 3高中 4大专 5本科 6研究生"`
	Degree    int8   `gorm:"column:degree;size:1;comment:学历: 1学士 2硕士 3博士"`
	Intro     string `gorm:"column:intro;comment:简介"`
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
	Status        int       `gorm:"comment:登录状态(1成功 2失败)" json:"status"`
	Msg           string    `gorm:"comment:提示消息" json:"msg"`
	LoginTime     time.Time `gorm:"comment:登录时间" json:"loginTime"`
	Module        string    `gorm:"comment:登录模块" json:"module"`
}

// UserBalance [...]
type UserBalance struct {
	GormModel
	UserID    uint    `gorm:"column:user_id;not null;comment:用户ID"`
	Balance   float64 `gorm:"column:balance;type:decimal(10,2);not null;default:0;comment:账户余额"` // 账户余额
	GiveMoney float64 `gorm:"column:give_money;type:decimal(10,2);default:0;comment:满赠的额度"`      // 满赠的额度
}

// UserBalanceLog [...]
type UserBalanceLog struct {
	GormModel
	UserID        uint      `gorm:"column:user_id;not null;comment:用户ID"`
	Amount        float64   `gorm:"column:amount;type:decimal(10,2);not null;comment:金额"`
	Action        int8      `gorm:"column:action;not null;size:1;comment:行为: 1充值 2消费 3推广 4提现 5岗位扣费"`
	PaymentMethod string    `gorm:"column:payment_method;not null;comment:支付方式: alipay支付宝支付 wechatpay微信支付"`
	PaymentStatus int8      `gorm:"column:payment_status;not null;size:1;default:0;comment:支付状态: 1已支付 2待支付 3支付失败"`
	PaymentTime   time.Time `gorm:"column:payment_time;not null;comment:支付时间"`
	TransactionID string    `gorm:"column:transaction_id;type:varchar(255);comment:第三方平台交易流水号"`
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
