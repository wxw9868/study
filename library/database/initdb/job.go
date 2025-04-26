package initdb

import "time"

// JobCategory 职位分类表
type JobCategory struct {
	GormModel
	ParentID  uint   `gorm:"column:parent_id;not null;default:0;comment:父ID"`
	Name      string `gorm:"column:name;type:varchar(255);not null;comment:职位分类名称"`
	Icon      string `gorm:"column:icon;type:varchar(255);comment:职位分类图标"`
	Image     string `gorm:"column:image;type:varchar(255);comment:分类图片路径"`
	Sort      int    `gorm:"column:sort;not null;comment:排序"`
	IsShow    int8   `gorm:"column:is_show;not null;default:0;comment:是否显示: 2不显示 1显示"`
	Recommend int8   `gorm:"column:recommend;not null;default:0;comment:金刚位显示: 2不显示 1显示"`
}

// Job 职位表
type Job struct {
	GormModel
	UserID       uint      `gorm:"column:user_id;not null;comment:用户ID"`
	CateID       uint      `gorm:"column:cate_id;not null;comment:职位分类"`
	Name         string    `gorm:"column:name;type:varchar(255);not null;comment:职位名称"`
	Number       uint      `gorm:"column:number;not null;comment:招聘人数"`
	MinSalary    float64   `gorm:"column:min_salary;not null;comment:最低薪资"`
	MaxSalary    float64   `gorm:"column:max_salary;not null;comment:最高薪资"`
	Description  string    `gorm:"column:description;type:longtext;not null;comment:职位描述"`
	StartTime    time.Time `gorm:"column:start_time;not null;comment:开始工作时间"`
	EndTime      time.Time `gorm:"column:end_time;not null;comment:结束工作时间"`
	SettlementID uint      `gorm:"column:settlement_id;not null;comment:薪资结算方式ID"`
	Province     uint      `gorm:"column:province;type:int(11);not null;comment:工作所在地：省"`
	City         uint      `gorm:"column:city;type:int(11);not null;comment:工作所在地：市"`
	District     uint      `gorm:"column:district;type:int(11);not null;comment:工作所在地：区/县"`
	Address      string    `gorm:"column:address;type:varchar(255);not null;comment:工作详细地址"`
	Liaison      string    `gorm:"column:liaison;type:varchar(60);not null;comment:联系人"`
	Mobile       string    `gorm:"column:mobile;type:varchar(20);not null;comment:手机号"`
	WechatId     string    `gorm:"column:wechat_id;type:varchar(30);comment:微信号"`
	IsDiscuss    int8      `gorm:"column:is_discuss;not null;default:1;comment:是否面议: 1是 2否"`
	IsShow       int8      `gorm:"column:is_show;not null;default:2;comment:是否上线: 1上线 2下线"`
	HideTime     time.Time `gorm:"column:hide_time;comment:下线时间"`
	ShowTime     time.Time `gorm:"column:show_time;comment:上线时间"`
	IsTop        int8      `gorm:"column:is_top;not null;default:2;comment:是否置顶: 1置顶 2取消"`
	Sort         int       `gorm:"column:sort;type:int(11);not null;default:0;comment:排序"`
	Status       int8      `gorm:"column:status;not null;comment:审核状态: 1待审核 2审核通过 3审核失败"`
	Comment      string    `gorm:"column:comment;type:text;comment:备注"`
}

// JobDeliver 职位投递表
type JobDeliver struct {
	GormModel
	JobID       uint `gorm:"column:job_id;not null;comment:岗位ID"`
	HireUserID  uint `gorm:"column:hire_user_id;not null;comment:发布岗位的用户ID"`
	ApplyUserID uint `gorm:"column:apply_user_id;not null;comment:申请岗位的用户ID"`
	Status      int8 `gorm:"column:status;not null;comment:状态: 1被查看 2已录取 3已拒绝 4待处理"`
}

// JobDeliveryMeter 职位投递扣费记录表
type JobDeliveryMeter struct {
	GormModel
	DeliveryID uint `gorm:"column:delivery_id;not null;comment:用户投递表ID"`
	PayStatus  int8 `gorm:"column:pay_status;not null;comment:扣费状态: 1已扣费 2没有扣费"`
}

// JobPromotion 职位推广表
type JobPromotion struct {
	GormModel
	UserID          uint      `gorm:"column:user_id;not null;comment:用户ID"`
	JobID           uint      `gorm:"column:job_id;not null;comment:岗位ID"`
	PromotionFee    float64   `gorm:"column:promotion_fee;type:decimal(10,2);not null;comment:推广费用"`
	PromotionType   int8      `gorm:"column:promotion_type;not null;comment:推广类型: 1banner广告 2普通广告 3列表置顶"`
	PromotionStatus int8      `gorm:"column:promotion_status;not null;comment:申请状态: 1待审核 2通过 3不通过"`
	StartTime       time.Time `gorm:"column:start_time;type:datetime;not null;comment:开始时间"`
	EndTime         time.Time `gorm:"column:end_time;type:datetime;not null;comment:结束时间"`
}

// JobSettlementType 薪资结算方式表
type JobSettlementType struct {
	GormModel
	SettlementName string `gorm:"column:settlement_name;type:varchar(30);not null;comment:结算名称"`
	Description    string `gorm:"column:description;type:varchar(255);comment:说明"`
	Sort           int    `gorm:"column:sort;type:int(11);not null;default:1;comment:排序"`
}

// Order 订单表
type Order struct {
	GormModel
	UserID         uint      `gorm:"column:user_id;not null;comment:用户ID"`
	OrderSn        string    `gorm:"column:order_sn;not null;comment:订单号"`
	OrderStatus    int8      `gorm:"column:order_status;not null;comment:订单状态: 1已确认 2已取消 3已完成 4已作废"`
	JobPromotionID uint      `gorm:"column:job_promotion_id;default:0;comment:职位推广ID"`
	OrderAmount    float64   `gorm:"column:order_amount;type:decimal(10,2);not null;comment:订单金额"`
	OrderType      int8      `gorm:"column:order_type;not null;comment:订单类型: 1充值 2消费 3推广 4提现"`
	OrderDesc      string    `gorm:"column:order_desc;comment:订单说明"`
	PaymentMethod  string    `gorm:"column:payment_method;not null;comment:支付方式: alipay支付宝支付 wechatpay微信支付"`
	PaymentStatus  int8      `gorm:"column:payment_status;not null;default:0;comment:支付状态: 1已支付 2待支付 3支付失败"`
	PaymentTime    time.Time `gorm:"column:payment_time;not null;comment:支付时间"`
	TransactionID  string    `gorm:"column:transaction_id;type:varchar(255);comment:第三方平台交易流水号"`
}
