package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Debug(c *gin.Context) {
	var input struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": input})
}

func GetUrlParams(c *gin.Context) {
	if name := c.Param("name"); name != "" {
		c.JSON(http.StatusOK, gin.H{"param": name})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"error": "Param 'name' not found"})
	return

}

func GetQueryParams(c *gin.Context) {
	// 带默认值的查询参数
	hobby := c.DefaultQuery("hobby", "play video games")
	if name := c.Query("name"); name != "" {

		c.String(http.StatusOK, "My name is  %s, my hobby is %s", name, hobby)
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
	return
}

func GetPostForm(c *gin.Context) {
	// 带默认值的表单参数
	message := c.DefaultPostForm("message", "不要回答，不要回答，不要回答")
	if name := c.PostForm("name"); name != "" {
		c.JSON(http.StatusOK, gin.H{"name": name, "message": message})
	}
}

func GetQueryAndFrom(c *gin.Context) {
	id := c.DefaultQuery("id", "1247")
	page := c.Query("page")
	name := c.PostForm("name")
	if page != "" && name != "" {
		c.JSON(http.StatusOK, gin.H{"id": id, "page": page, "name": name})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"error": "内部错误，请检查参数后重试"})

}

func UploadFile(c *gin.Context) {
	file, _ := c.FormFile("file")
	log.Println(file.Filename)
	err := c.SaveUploadedFile(file, "./"+file.Filename)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "上传成功"})
}

func UploadFiles(c *gin.Context) {
	// 多文件
	form, _ := c.MultipartForm()
	files := form.File["upload[]"]

	for _, file := range files {
		fmt.Println(file.Filename)
		// 上传文件到指定的路径
		// c.SaveUploadedFile(file, dst)
	}
	c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
}
