package main

import (
	"gin-gorm-demo/config"
	"gin-gorm-demo/models"
	"gin-gorm-demo/router"
	"gin-gorm-demo/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"log"
)

func main() {
	// 连接数据库
	config.ConnectDatabase()

	// 自动迁移数据库
	err := config.DB.AutoMigrate(&models.User{})
	if err != nil {
		return
	}

	r := gin.New()

	// 获取 Gin 的默认验证器引擎
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 注册自定义验证器
		err = v.RegisterValidation("StartWith", utils.StartWith)
		if err == nil {
			log.Println("自定义验证器注册成功")
		}

	} else {
		log.Println("自定义验证器注册失败")
	}

	// 使用中间件
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	v1 := r.Group("/v1")

	// 注册各模块路由
	studyRouter := router.StudyRouter{}
	studyRouter.InitStudyRouter(v1)

	userRouter := router.UserRouter{}
	userRouter.InitUserRouter(v1)

	skyRouter := router.SkyRouter{}
	skyRouter.InitSkyRouter(v1)

	// 启动服务
	startErr := r.Run(":8080")
	if startErr != nil {
		return
	}
}
