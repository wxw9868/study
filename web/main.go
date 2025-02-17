/*
 * @Author: wxw9868@163.com
 * @Date: 2024-07-18 18:45:27
 * @LastEditTime: 2025-02-17 19:07:27
 * @LastEditors: wxw9868@163.com
 * @FilePath: /web/main.go
 * @Description: 灵活就业服务平台
 */
// main.go
package main

import (
	"github.com/wxw9868/checkin/db"
	"github.com/wxw9868/checkin/handlers"
	"github.com/wxw9868/checkin/models"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化数据库
	// db.InitDatabase()
	db.InitDB()
	defer db.CloseDB()

	// 自动迁移表结构
	db.DB().AutoMigrate(&models.User{}, &models.Checkin{})

	// 初始化 Gin
	r := gin.Default()

	// 路由
	r.POST("/checkin", handlers.Checkin)

	// 启动服务
	r.Run(":8080")
}
