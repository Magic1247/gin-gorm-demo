package main

import (
	"gin-gorm-demo/config"
	"gin-gorm-demo/controllers"
	"gin-gorm-demo/models"
	"github.com/gin-gonic/gin"
)

func main() {
	// 连接数据库
	config.ConnectDatabase()

	// 自动迁移数据库
	err := config.DB.AutoMigrate(&models.User{})
	if err != nil {
		return
	}

	// 创建路由
	r := gin.Default()

	// 路由配置
	r.GET("/users", controllers.GetUsers)
	r.POST("/users", controllers.CreateUser)
	r.GET("/users/:id", controllers.GetUserByID)
	r.PUT("/users/:id", controllers.UpdateUser)
	r.DELETE("/users/:id", controllers.DeleteUser)

	// 启动服务
	startErr := r.Run(":8080")
	if startErr != nil {
		return
	}
}
