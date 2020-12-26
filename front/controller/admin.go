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
			"data": "",
			"message":err,
		})
	}

	db := db.Init()

	result := db.Model(model.Admin{}).Where("username = ? and password = ?", admin.Username, admin.Password).First(&user).RecordNotFound()

	if result == true {
		c.JSON(200, gin.H{
			"code": 20001,
			"message":"用户名密码错误",
			"data": "",
		})
	} else {
		c.JSON(200, gin.H{
			"code": 20000,
			"message": "success",
			"data": "",
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

//CategoryList
func CategoryList(c *gin.Context){
	categories := []model.Category{}

	pageindex := c.DefaultQuery("currentPage", "1")

	currentPage, _ := strconv.Atoi(pageindex)

	db := db.Init()
	var totalNum int

	pagesize := 10
	db.Model(model.Category{}).Count(&totalNum)
	totalPage := totalNum / pagesize

	db.Model(model.Category{}).Order("id desc").Offset((currentPage - 1) * pagesize).Limit(pagesize).Find(&categories)

	c.JSON(200, gin.H{
		"code": 20000,
		"data": gin.H{
			"total": totalPage,
			"data":  categories,
		},
	})
}


//CategoryGet 获取
func CategoryGet(c *gin.Context) {
	category := &model.Category{}

	id := c.DefaultQuery("id", "1")

	db := db.Init()

	db.Model(model.Category{}).Where("id = ? ", id).First(&category)

	c.JSON(200, gin.H{
		"code": 20000,
		"data": category,
		"message":"success",
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

//CategoryEdit 修改分类
func CategoryEdit(c *gin.Context) {
	category := &model.Category{}

	err := c.ShouldBindJSON(category)

	if err != nil {
		c.JSON(200, gin.H{
			"code": 20001,
			"data": err,
		})

		return
	}


	db := db.Init()

	db.Model(model.Category{}).Updates(&category)

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
