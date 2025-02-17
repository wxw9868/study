/*
 * @Author: wxw9868@163.com
 * @Date: 2025-02-17 18:43:06
 * @LastEditTime: 2025-02-17 19:07:47
 * @LastEditors: wxw9868@163.com
 * @FilePath: /web/handlers/checkin.go
 * @Description: 灵活就业服务平台
 */
package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wxw9868/checkin/db"
	"github.com/wxw9868/checkin/models"
	"github.com/wxw9868/checkin/utils"
)

// CheckinRequest 打卡请求结构体
type CheckinRequest struct {
	UserID uint    `form:"user_id" binding:"required"`
	Text   string  `form:"text" binding:"required"`
	Lat    float64 `form:"lat" binding:"required"`
	Lng    float64 `form:"lng" binding:"required"`
}

// Checkin 处理打卡请求
func Checkin(c *gin.Context) {
	var req CheckinRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 获取用户打卡范围
	var user models.User
	if err := db.DB().First(&user, req.UserID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "用户不存在"})
		return
	}

	// 校验是否在范围内
	distance := utils.Haversine(req.Lat, req.Lng, user.CenterLat, user.CenterLng)
	isInRange := distance <= float64(user.Radius)

	// 保存打卡记录
	checkin := models.Checkin{
		UserID:    req.UserID,
		Text:      req.Text,
		Lat:       req.Lat,
		Lng:       req.Lng,
		IsInRange: isInRange,
	}

	// 保存文件
	if file, err := c.FormFile("media"); err == nil {
		filePath, err := utils.SaveFile(file, "uploads")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "文件保存失败"})
			return
		}
		checkin.MediaURL = filePath
	}

	if err := db.DB().Create(&checkin).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "打卡记录保存失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":    "success",
		"isInRange": isInRange,
	})
}
