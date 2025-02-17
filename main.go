/*
 * @Author: wxw9868@163.com
 * @Date: 2024-07-04 10:26:38
 * @LastEditTime: 2025-02-17 18:37:42
 * @LastEditors: wxw9868@163.com
 * @FilePath: /study/main.go
 * @Description: 灵活就业服务平台
 */
package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run("0.0.0.0:80") // 监听并在 0.0.0.0:8080 上启动服务
}
