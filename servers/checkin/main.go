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
	db.InitDatabase()

	// 初始化 Gin
	r := gin.Default()

	r.NoMethod(func(c *gin.Context) { c.JSON(405, gin.H{"msg": "method not allowed"}) })
	r.NoRoute(func(c *gin.Context) { c.JSON(404, gin.H{"msg": "not found"}) })

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// 路由
	r.POST("/checkin", handlers.Checkin)

	// 启动服务
	r.Run(":8080")
}
