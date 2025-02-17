/*
 * @Author: wxw9868@163.com
 * @Date: 2025-02-17 18:40:49
 * @LastEditTime: 2025-02-17 18:40:57
 * @LastEditors: wxw9868@163.com
 * @FilePath: /study/web/models/checkin.go
 * @Description: 灵活就业服务平台
 */
package models

import "gorm.io/gorm"

type Checkin struct {
	gorm.Model
	UserID    uint    `json:"user_id"`     // 用户ID
	Text      string  `json:"text"`        // 打卡文字
	Lat       float64 `json:"lat"`         // 打卡纬度
	Lng       float64 `json:"lng"`         // 打卡经度
	IsInRange bool    `json:"is_in_range"` // 是否在范围内
	MediaURL  string  `json:"media_url"`   // 图片/视频文件路径
}
