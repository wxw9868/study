package initdb

import (
	"time"

	"gorm.io/gorm"
)

// GormModel 基础模型
type GormModel struct {
	ID        uint           `gorm:"primarykey"`
	CreatedAt time.Time      `gorm:"column:created_at;index"` // Added index
	UpdatedAt time.Time      `gorm:"column:updated_at;index"` // Added index
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index"`
}

// Settings 系统设置表
type Settings struct {
	GormModel
	WechatShow      int8    `gorm:"column:wechat_show;type:tinyint(1);not null;default:2;comment:微信显示: 1展示 2隐藏"`
	Price           float64 `gorm:"column:price;type:decimal(10,2);not null;default:0.00;comment:投递一次岗位的收费价格"`
	CustomerService string  `gorm:"column:customer_service;type:varchar(255);comment:客户服务"` // Increased length to 255
}

// University 全国高等学校名单表
type University struct {
	GormModel
	SchoolName          string `gorm:"column:school_name;type:varchar(100);uniqueIndex;not null;comment:学校名称"`
	SchoolIdentifier    string `gorm:"column:school_identifier;type:varchar(30);uniqueIndex;not null;comment:学校标识码"`
	CompetentDepartment string `gorm:"column:competent_department;type:varchar(100);index;comment:主管部门"` // Added index
	Location            string `gorm:"column:location;type:varchar(100);index;comment:所在地"`              // Added index
	SchoolLevel         string `gorm:"column:school_level;type:varchar(20);index;comment:办学层次"`          // Added index
	Note                string `gorm:"column:note;type:varchar(255);comment:备注"`
}

// Region 地区表
type Region struct {
	GormModel
	Name     string `gorm:"column:name;type:varchar(100);not null;comment:地区名称"`
	Level    int8   `gorm:"column:level;type:tinyint(1);index;default:0;comment:地区级别"` // Added index
	ParentID uint   `gorm:"column:parent_id;index;comment:父级ID"`                       // Changed type to uint, kept index
	RegionID uint   `gorm:"column:region_id;uniqueIndex;comment:地区ID"`                 // Changed type to uint, kept uniqueIndex
}

// Ad 广告表
type Ad struct {
	GormModel
	Type      int8      `gorm:"column:type;type:tinyint(1);index;not null;comment:广告类型: 1banner广告 2普通广告"` // Added index
	Name      string    `gorm:"column:name;type:varchar(100);not null;comment:广告名称"`
	URL       string    `gorm:"column:url;type:varchar(255);not null;comment:链接地址"`
	Image     string    `gorm:"column:image;type:varchar(255);not null;comment:图片地址"`
	StartTime time.Time `gorm:"column:start_time;index;not null;comment:开始时间"`                                 // Added index
	EndTime   time.Time `gorm:"column:end_time;index;not null;comment:结束时间"`                                   // Added index
	IsShow    int8      `gorm:"column:is_show;type:tinyint(1);index;not null;default:1;comment:是否显示: 1显示 2隐藏"` // Added index
	IsFree    int8      `gorm:"column:is_free;type:tinyint(1);index;not null;default:2;comment:是否免费: 1是 2不是"`  // Added index
	Sorting   int       `gorm:"column:sorting;index;not null;default:0;comment:排序"`                            // Added index
}

// Article 文章表
type Article struct {
	GormModel
	Title   string `gorm:"column:title;type:varchar(255);not null;comment:文章标题"`
	Author  string `gorm:"column:author;type:varchar(100);not null;comment:文章作者"`
	Content string `gorm:"column:content;type:longtext;not null;comment:文章内容"`
	IsShow  int8   `gorm:"column:is_show;type:tinyint(1);index;not null;default:1;comment:是否显示: 1显示 2隐藏"` // Added index
	IsTop   int8   `gorm:"column:is_top;type:tinyint(1);index;not null;default:2;comment:是否置顶: 1置顶 2取消"`  // Added index
}

// Coupon 优惠券表
type Coupon struct {
	GormModel
	Type       int8      `gorm:"column:type;type:tinyint(1);index;not null;default:1;comment:优惠券类型: 1满减券 2折扣券"` // Added index
	Name       string    `gorm:"column:name;type:varchar(100);uniqueIndex;not null;comment:优惠券名称"`              // Added uniqueIndex
	Level      int16     `gorm:"column:level;type:smallint(4);index;default:1;comment:优惠券等级"`                   // Added index
	FullAmount float64   `gorm:"column:full_amount;type:decimal(10,2);not null;comment:满减金额"`
	SendAmount float64   `gorm:"column:send_amount;type:decimal(10,2);not null;comment:赠送金额"`
	StartTime  time.Time `gorm:"column:start_time;index;not null;comment:开始时间"`                               // Added index
	EndTime    time.Time `gorm:"column:end_time;index;not null;comment:结束时间"`                                 // Added index
	Status     int8      `gorm:"column:status;type:tinyint(1);index;not null;default:1;comment:状态: 1正常 2已过期"` // Added index
	Desc       string    `gorm:"column:desc;type:varchar(255);comment:优惠券介绍"`
}

// CouponLog 优惠券使用记录表
type CouponLog struct {
	GormModel
	UserID     uint      `gorm:"column:user_id;index;not null;comment:用户ID"`    // Changed type to uint, added index
	OrderID    uint      `gorm:"column:order_id;index;not null;comment:订单ID"`   // Changed type to uint, added index
	CouponID   uint      `gorm:"column:coupon_id;index;not null;comment:优惠券ID"` // Changed type to uint, added index
	Amount     float64   `gorm:"column:amount;type:decimal(10,2);not null;default:0.00;comment:优惠金额"`
	Status     int8      `gorm:"column:status;type:tinyint(1);index;not null;default:1;comment:使用状态: 1已使用 2已作废"`
	CreateTime time.Time `gorm:"column:create_time;index;not null;comment:使用时间"`
	// 添加复合唯一索引，确保一个订单只能使用一张优惠券
	_ int `gorm:"uniqueIndex:idx_user_order_coupon,columns:user_id,order_id,coupon_id"`
}
