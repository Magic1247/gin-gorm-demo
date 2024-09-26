package router

import (
	"gin-gorm-demo/controllers"
	"github.com/gin-gonic/gin"
)

type StudyRouter struct{}

func (s *StudyRouter) InitStudyRouter(Router *gin.RouterGroup) {
	studyRouter := Router.Group("/study")

	studyRouter.POST("/debug", controllers.Debug)
	studyRouter.GET("/paramTest/:name", controllers.GetUrlParams)
	studyRouter.GET("/queryTest", controllers.GetQueryParams)

	studyRouter.POST("/postForm", controllers.GetPostForm)

	studyRouter.POST("/queryAndParam", controllers.GetQueryAndFrom)

	studyRouter.POST("/upload", controllers.UploadFile)
	studyRouter.POST("/uploads", controllers.UploadFiles)
}
