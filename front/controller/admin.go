package controller

import (
	"go-movie/db"
	"go-movie/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//Login 登录页面
func Alogin(c *gin.Context) {
	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the index.html template
		"login.html",
		// Pass the data that the page uses (in this case, 'title')
		gin.H{
			"title": "登录",
		},
	)
}

func APlogin(c *gin.Context) {
	admin := &model.Admin{}
	user := model.Admin{}
	err := c.ShouldBindJSON(admin)

	if err != nil {
		c.JSON(200, gin.H{
			"code": 20001,
			"data": err,
		})
	}

	db := db.Init()

	result := db.Model(model.Admin{}).Where("username = ? and password = ?", admin.Username, admin.Password).First(&user).RecordNotFound()

	if result == true {
		c.JSON(200, gin.H{
			"code": 20001,
			"data": "用户名密码错误",
		})
	} else {
		c.JSON(200, gin.H{
			"code": 20000,
			"data": "admin-token",
		})
	}

}

// AInfo 用户信息
func AInfo(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 20000,
		"data": "admin-token",
	})
}

//AList 列表
func AList(c *gin.Context) {
	posts := []model.Post{}

	pageindex := c.DefaultQuery("currentPage", "1")

	currentPage, _ := strconv.Atoi(pageindex)

	db := db.Init()
	var totalNum int

	pagesize := 10
	db.Model(model.Post{}).Count(&totalNum)
	totalPage := totalNum / pagesize

	db.Model(model.Post{}).Order("id desc").Offset((currentPage - 1) * pagesize).Limit(pagesize).Find(&posts)

	c.JSON(200, gin.H{
		"code": 20000,
		"data": gin.H{
			"total": totalPage,
			"data":  posts,
		},
	})
}

//CategoryAdd 添加分类
func CategoryAdd(c *gin.Context) {
	category := &model.Category{}

	err := c.ShouldBindJSON(category)

	if err != nil {
		c.JSON(200, gin.H{
			"code": 20001,
			"data": err,
		})
	}

	db := db.Init()

	db.Model(model.Category{}).Create(category)

	c.JSON(200, gin.H{
		"code": 20000,
		"data": "success",
	})

}

//Aindex 首页
func Aindex(c *gin.Context) {
	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the index.html template
		"index.html",
		// Pass the data that the page uses (in this case, 'title')
		gin.H{
			"title": "登录",
		},
	)
}
