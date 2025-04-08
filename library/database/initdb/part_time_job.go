/*
 * @Author: wxw9868 wxw9868@163.com
 * @Date: 2025-01-02 16:08:43
 * @LastEditors: wxw9868 wxw9868@163.com
 * @LastEditTime: 2025-01-15 14:37:36
 * @FilePath: /study/initdb/part_time_job.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package initdb

import (
	"time"

	"gorm.io/gorm"
)

// GormModel 基础模型
type GormModel struct {
	ID        uint           `gorm:"primarykey"`
	CreatedAt time.Time      `gorm:"column:created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index"`
}

// Settings 系统设置表
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

// Region 地区表
type Region struct {
	GormModel
	Name     string `gorm:"column:name;type:longtext;comment:"`
	Level    int8   `gorm:"column:level;type:tinyint(4);default:0;comment:"`
	ParentID int    `gorm:"index:parent_id;column:parent_id;type:int(10);comment:"`
	RegionID int    `gorm:"column:region_id;type:int(11);comment:"`
}

// Ad 广告表
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

// Article 文章表
type Article struct {
	GormModel
	Title   string `gorm:"column:title;type:varchar(255);not null;comment:文章标题"` // 标题
	Author  string `gorm:"column:author;type:varchar(255);not null;comment:文章作者"`
	Content string `gorm:"column:content;type:longtext;not null;comment:文章内容"` // 内容
}

// Coupon 优惠券表
type Coupon struct {
	GormModel
	Name       string  `gorm:"column:name;type:varchar(255);comment:优惠券名称"`                                // 优惠券名称
	Level      int16   `gorm:"column:level;type:smallint(6);comment:优惠券等级"`                                // 优惠券等级
	FullAmount float64 `gorm:"column:full_amount;type:decimal(10,2);not null;comment:金额"`                  // 金额
	SendAmount float64 `gorm:"column:send_amount;type:decimal(10,2);not null;comment:满多少金额送多少: 例如满100送10"` // 满多少金额送多少：例如满100送10
	IsUse      int8    `gorm:"column:is_use;not null;default:1;comment:是否使用: 1已使用 2没使用"`                   // 是否使用：1已使用；2没使用
	Desc       string  `gorm:"column:desc;type:varchar(255);comment:优惠券介绍"`                                // 优惠券介绍
}

// CouponLog 优惠券使用记录表
type CouponLog struct {
	GormModel
	UserID   uint `gorm:"column:user_id;not null;comment:用户ID"`
	OrderID  uint `gorm:"column:order_id;not null;comment:订单ID"`
	CouponID uint `gorm:"column:coupon_id;not null;comment:优惠券ID"`
}
