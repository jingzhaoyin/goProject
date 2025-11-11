package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// var DB *gorm.DB

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/student", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "查询学生信息成功",
		})
	})
	r.POST("/create_student", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "创建学生信息成功",
		})
	})
	r.PUT("/updata_student", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "更新学生信息成功",
		})
	})
	r.DELETE("/delete_student", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "删除学生信息成功",
		})
	})

	r.LoadHTMLGlob("./templates/*")
	r.GET("/demo", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"name": "admin",
			"pwd":  "123456",
		})
	})

	r.GET("/", func(c *gin.Context) {
		name := c.Query("name")
		pwd := c.Query("pwd")
		// fmt.Printf("name:%s ; pwd:%s",name,pwd)
		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"pwd":  pwd,
		})
	})

	r.Run() // 监听并在

}
