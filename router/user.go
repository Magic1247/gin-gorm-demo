package router

import (
	"gin-gorm-demo/controllers"
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (u *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("/user")
	userRouter.GET("/users", controllers.GetUsers)
	userRouter.POST("/users", controllers.CreateUser)
	userRouter.GET("/users/:id", controllers.GetUserByID)
	userRouter.PUT("/users/:id", controllers.UpdateUser)
	userRouter.DELETE("/users/:id", controllers.DeleteUser)
}
