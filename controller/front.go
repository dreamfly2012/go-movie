package controller

import (
	"go-movie/db"
	"go-movie/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//Index  首页控制器
func Index(c *gin.Context) {
	db := db.Init()

	pagesize := 5

	posts := []model.Post{}

	var pageindex int
	var prevID int
	var nextID int
	var prevLink string
	var nextLink string
	var totalNum int
	id := c.Param("id")
	if id != "" {
		pageindex, _ = strconv.Atoi(id)
	} else {
		pageindex = 1
	}

	prevID = pageindex - 1
	if prevID == 0 {
		prevLink = "/page/1"
	} else {
		prevLink = "/page/" + strconv.Itoa(prevID)
	}
	nextID = pageindex + 1
	db.Model(model.Post{}).Where("status = ?", 0).Count(&totalNum)
	totalPage := totalNum / pagesize
	if nextID > totalPage {
		nextLink = "/page/" + strconv.Itoa(totalPage)
	} else {
		nextLink = "/page/" + strconv.Itoa(nextID)
	}

	db.Model(model.Post{}).Where("status = ?", 0).Order("id desc").Offset((pageindex - 1) * pagesize).Limit(pagesize).Find(&posts)

	c.HTML(

		http.StatusOK,

		"index.html",

		gin.H{
			"title":        "电影宇宙",
			"posts":        posts,
			"prev_link":    prevLink,
			"next_link":    nextLink,
			"current_page": pageindex,
			"total_page":   totalPage,
		},
	)
}

//About  关于
func About(c *gin.Context) {
	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the index.html template
		"about.html",
		// Pass the data that the page uses (in this case, 'title')
		gin.H{
			"title": "关于",
		},
	)
}

//Article 文章详情页
func Article(c *gin.Context) {
	db := db.Init()
	id := c.Param("id")
	//fmt.Print(id)
	post := model.Post{}
	db.Model(model.Post{}).Where("id = ?", id).First(&post)
	c.HTML(
		http.StatusOK,
		"single.html",
		gin.H{
			"title": post.Title,
			"post":  post,
		},
	)
}
