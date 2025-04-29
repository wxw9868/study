package database

import "time"

// JobCategory 职位分类表
type JobCategory struct {
	GormModel
	ParentID  uint   `gorm:"column:parent_id;index;not null;default:0;comment:父ID"`                             // Changed type to uint, added index
	Name      string `gorm:"column:name;type:varchar(100);uniqueIndex:idx_parent_name;not null;comment:职位分类名称"` // Added uniqueIndex
	Icon      string `gorm:"column:icon;type:varchar(100);comment:职位分类图标"`
	Image     string `gorm:"column:image;type:varchar(200);comment:分类图片路径"`
	Sorting   int    `gorm:"column:sorting;index;not null;default:0;comment:排序"`                                // Added index
	IsShow    int8   `gorm:"column:is_show;type:tinyint(1);index;not null;default:2;comment:是否显示: 1显示 2不显示"`    // Changed default to 2, added index
	Recommend int8   `gorm:"column:recommend;type:tinyint(1);index;not null;default:2;comment:金刚位显示: 1显示 2不显示"` // Changed default to 2, added index
}

// Job 职位表
type Job struct {
	GormModel
	UserID       uint      `gorm:"column:user_id;index;not null;comment:用户ID"`                       // Changed type to uint, added index
	CategoryID   uint      `gorm:"column:category_id;index;not null;comment:职位分类ID"`                 // Changed type to uint, added index
	Name         string    `gorm:"column:name;type:varchar(100);index;not null;comment:职位名称"`        // 添加索引
	RecruitNum   uint      `gorm:"column:recruit_num;not null;comment:招聘人数"`                         // Changed type to uint
	MinSalary    float64   `gorm:"column:min_salary;type:decimal(10,2);index;not null;comment:最低薪资"` // 添加索引
	MaxSalary    float64   `gorm:"column:max_salary;type:decimal(10,2);index;not null;comment:最高薪资"` // 添加索引
	Description  string    `gorm:"column:description;type:text;not null;comment:职位描述"`
	StartTime    time.Time `gorm:"column:start_time;index;not null;comment:开始工作时间"`      // Added index
	EndTime      time.Time `gorm:"column:end_time;index;not null;comment:结束工作时间"`        // Added index
	SettlementID uint      `gorm:"column:settlement_id;index;not null;comment:薪资结算方式ID"` // Changed type to uint, added index
	Province     uint      `gorm:"column:province;index;not null;comment:工作所在地：省"`       // Changed type to uint, added index
	City         uint      `gorm:"column:city;index;not null;comment:工作所在地：市"`           // Changed type to uint, added index
	District     uint      `gorm:"column:district;index;not null;comment:工作所在地：区/县"`     // Changed type to uint, added index
	Address      string    `gorm:"column:address;type:varchar(200);not null;comment:工作详细地址"`
	Liaison      string    `gorm:"column:liaison;type:varchar(50);not null;comment:联系人"`
	Mobile       string    `gorm:"column:mobile;type:varchar(20);index;not null;comment:手机号"`                             // Added index
	WechatID     string    `gorm:"column:wechat_id;type:varchar(30);index;comment:微信号"`                                   // Added index
	IsDiscuss    int8      `gorm:"column:is_discuss;type:tinyint(1);index;not null;default:1;comment:是否面议: 1是 2否"`        // 添加索引
	IsShow       int8      `gorm:"column:is_show;type:tinyint(1);index;not null;default:2;comment:是否上线: 1上线 2下线"`         // Added index, default 2
	IsTop        int8      `gorm:"column:is_top;type:tinyint(1);index;not null;default:2;comment:是否置顶: 1置顶 2取消"`          // Added index, default 2
	Sorting      int       `gorm:"column:sorting;index;not null;default:0;comment:排序"`                                    // Added index
	Status       int8      `gorm:"column:status;type:tinyint(1);index;not null;default:1;comment:审核状态: 1待审核 2审核通过 3审核失败"` // Added index, default 1
	Comment      string    `gorm:"column:comment;type:varchar(255);comment:备注"`
	ViewCount    uint      `gorm:"column:view_count;not null;default:0;comment:浏览次数"`
	ApplyCount   uint      `gorm:"column:apply_count;not null;default:0;comment:申请次数"`
	LastEditTime time.Time `gorm:"column:last_edit_time;index;comment:最后编辑时间"` // 添加索引
	// 添加复合索引，提高查询效率
	_ int `gorm:"index:idx_category_status_time,columns:category_id,status,start_time"`
}

// JobDeliver 职位投递表
type JobDeliver struct {
	GormModel
	JobID       uint `gorm:"column:job_id;index;not null;comment:岗位ID"`                                              // Changed type to uint, added index
	HireUserID  uint `gorm:"column:hire_user_id;index;not null;comment:发布岗位的用户ID"`                                   // Changed type to uint, added index
	ApplyUserID uint `gorm:"column:apply_user_id;index;not null;comment:申请岗位的用户ID"`                                  // Changed type to uint, added index
	Status      int8 `gorm:"column:status;type:tinyint(1);index;not null;default:4;comment:状态: 1被查看 2已录取 3已拒绝 4待处理"` // Added index, default 4
	// 添加复合唯一索引，确保一个用户不能重复投递同一个职位
	_ int `gorm:"uniqueIndex:idx_job_apply,columns:job_id,apply_user_id"`
}

// JobDeliveryMeter 职位投递扣费记录表
type JobDeliveryMeter struct {
	GormModel
	DeliveryID uint `gorm:"column:delivery_id;uniqueIndex;not null;comment:职位投递表ID"`                        // Changed type to uint, added uniqueIndex
	Status     int8 `gorm:"column:status;type:tinyint(1);index;not null;default:2;comment:扣费状态: 1已扣费 2待扣费"` // Added index, default 2
}

// JobPromotion 职位推广表
type JobPromotion struct {
	GormModel
	UserID    uint      `gorm:"column:user_id;index;not null;comment:用户ID"`    // Changed type to uint, added index
	JobID     uint      `gorm:"column:job_id;index;not null;comment:岗位ID"`     // Changed type to uint, added index
	StartTime time.Time `gorm:"column:start_time;index;not null;comment:开始时间"` // Added index
	EndTime   time.Time `gorm:"column:end_time;index;not null;comment:结束时间"`   // Added index
	Fee       float64   `gorm:"column:fee;type:decimal(10,2);not null;comment:推广费用"`
	Type      int8      `gorm:"column:type;type:tinyint(1);index;not null;comment:推广类型: 1banner广告 2普通广告 3列表置顶"`    // Added index
	Status    int8      `gorm:"column:status;type:tinyint(1);index;not null;default:2;comment:申请状态: 1通过 2待审核 3失败"` // Added index, default 2
}

// JobSettlementType 薪资结算方式表
type JobSettlementType struct {
	GormModel
	Name        string `gorm:"column:name;type:varchar(30);uniqueIndex;not null;comment:结算名称"` // Added uniqueIndex
	Description string `gorm:"column:description;type:varchar(100);comment:说明"`
	Sorting     int    `gorm:"column:sorting;index;not null;default:0;comment:排序"` // Added index
}

// Order 订单表
type Order struct {
	GormModel
	UserID         uint      `gorm:"column:user_id;index;not null;comment:用户ID"`                        // Changed type to uint, added index
	JobPromotionID uint      `gorm:"column:job_promotion_id;index;default:0;comment:职位推广ID"`            // Changed type to uint, added index
	OrderSn        string    `gorm:"column:order_sn;type:varchar(50);uniqueIndex;not null;comment:订单号"` // Added uniqueIndex
	Amount         float64   `gorm:"column:amount;type:decimal(10,2);not null;comment:订单金额"`
	Type           int8      `gorm:"column:type;type:tinyint(1);index;not null;comment:订单类型: 1充值 2消费 3推广 4提现"`                 // Added index
	Status         int8      `gorm:"column:status;type:tinyint(1);index;not null;default:1;comment:订单状态: 1已确认 2已取消 3已完成 4已作废"` // Added index, default 1
	Desc           string    `gorm:"column:order_desc;type:varchar(255);comment:订单说明"`
	PaymentMethod  string    `gorm:"column:payment_method;type:varchar(20);index;not null;comment:支付方式: alipay支付宝支付 wechatpay微信支付"` // Added index
	PaymentTime    time.Time `gorm:"column:payment_time;index;comment:支付时间"`                                                        // Added index, removed not null
	PaymentStatus  int8      `gorm:"column:payment_status;type:tinyint(1);index;not null;default:2;comment:支付状态: 1已支付 2待支付 3支付失败"`  // Added index, default 2
	TransactionID  string    `gorm:"column:transaction_id;type:varchar(100);uniqueIndex;comment:第三方平台交易流水号"`                        // Changed to uniqueIndex
	CreateIP       string    `gorm:"column:create_ip;type:varchar(45);comment:创建订单的IP地址"`
}
