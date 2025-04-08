package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	CenterLat float64 `json:"center_lat"` // 打卡范围中心纬度
	CenterLng float64 `json:"center_lng"` // 打卡范围中心经度
	Radius    int     `json:"radius"`     // 打卡范围半径（米）
}
