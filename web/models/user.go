/*
 * @Author: wxw9868@163.com
 * @Date: 2025-02-17 18:40:29
 * @LastEditTime: 2025-02-17 19:06:41
 * @LastEditors: wxw9868@163.com
 * @FilePath: /web/models/user.go
 * @Description: 灵活就业服务平台
 */
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
