package router

import (
	"gin-gorm-demo/controllers"
	"github.com/gin-gonic/gin"
)

type SkyRouter struct{}

func (s *SkyRouter) InitSkyRouter(r *gin.RouterGroup) {
	skyRouter := r.Group("/sky")
	skyRouter.POST("/login", controllers.LoginSky)
}
