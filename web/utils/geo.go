/*
 * @Author: wxw9868@163.com
 * @Date: 2025-02-17 18:41:44
 * @LastEditTime: 2025-02-17 18:41:56
 * @LastEditors: wxw9868@163.com
 * @FilePath: /study/web/utils/geo.go
 * @Description: 灵活就业服务平台
 */
package utils

import "math"

// Haversine 公式计算两点间距离（单位：米）
func Haversine(lat1, lng1, lat2, lng2 float64) float64 {
	const R = 6371000 // 地球半径（米）
	dLat := (lat2 - lat1) * (math.Pi / 180)
	dLng := (lng2 - lng1) * (math.Pi / 180)
	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(lat1*(math.Pi/180))*math.Cos(lat2*(math.Pi/180))*
			math.Sin(dLng/2)*math.Sin(dLng/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return R * c
}
