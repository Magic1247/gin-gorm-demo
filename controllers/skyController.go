package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type loginInput struct {
	Username string `form:"username" json:"username" xml:"username"  binding:"required,StartWith=sky"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func LoginSky(c *gin.Context) {
	var loginInput loginInput
	if err := c.ShouldBind(&loginInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if loginInput.Username == "" || loginInput.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "参数错误"})
		return
	}
	if loginInput.Username == "sky001" && loginInput.Password == "123" {
		c.JSON(http.StatusOK, gin.H{"message": "登陆成功"})
		return
	} else {
		c.JSON(http.StatusFailedDependency, gin.H{"message": "用户名或密码错误"})
		return
	}
}
