package controller

import (
	"fmt"
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

	categories := []model.Category{}
	db.Model(model.Category{}).Order("id desc").Find(&categories)

	tags := []model.Tag{}
	db.Model(model.Tag{}).Order("id desc").Find(&tags)

	latest := []model.Post{}
	db.Model(model.Post{}).Where("status = ?", 0).Order("id desc").Limit(3).Find(&latest)

	c.HTML(

		http.StatusOK,

		"index.html",

		gin.H{
			"title":        "电影宇宙",
			"posts":        posts,
			"tags":         tags,
			"latest":       latest,
			"categories":   categories,
			"prev_link":    prevLink,
			"next_link":    nextLink,
			"current_page": pageindex,
			"total_page":   totalPage,
		},
	)
}

//

//Category  分类控制器
func Category(c *gin.Context) {
	db := db.Init()

	pagesize := 5

	posts := []model.Post{}
	var category = model.Category{}
	var pageindex int
	var prevID int
	var nextID int
	var prevLink string
	var nextLink string
	var totalNum int
	id := c.Param("id")
	url := c.Param("url")
	if id != "" {
		pageindex, _ = strconv.Atoi(id)
	} else {
		pageindex = 1
	}

	db.Model(model.Category{}).Where("url = ?", url).First(&category)

	prevID = pageindex - 1
	if prevID == 0 {
		prevLink = "/" + category.URL + "/1"
	} else {
		prevLink = "/" + category.URL + "/" + strconv.Itoa(prevID)
	}
	nextID = pageindex + 1
	db.Model(model.Post{}).Where("status = ? and cid = ? ", 0, category.ID).Count(&totalNum)
	totalPage := totalNum / pagesize

	if totalNum > 0 && totalPage == 0 {
		totalPage = 1
	}

	if nextID > totalPage {
		nextLink = "/" + category.URL + "/" + strconv.Itoa(totalPage)
	} else {
		nextLink = "/" + category.URL + "/" + strconv.Itoa(nextID)
	}

	db.Model(model.Post{}).Where("status = ? and cid = ? ", 0, category.ID).Order("id desc").Offset((pageindex - 1) * pagesize).Limit(pagesize).Find(&posts)

	categories := []model.Category{}
	db.Model(model.Category{}).Order("id desc").Find(&categories)

	tags := []model.Tag{}
	db.Model(model.Tag{}).Order("id desc").Find(&tags)

	latest := []model.Post{}
	db.Model(model.Post{}).Where("status = ?", 0).Order("id desc").Limit(3).Find(&latest)

	c.HTML(

		http.StatusOK,

		"index.html",

		gin.H{
			"title":        "电影宇宙",
			"posts":        posts,
			"latest":       latest,
			"tags":         tags,
			"categories":   categories,
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

	categories := []model.Category{}
	db.Model(model.Category{}).Order("id desc").Find(&categories)
	fmt.Println(categories)
	c.HTML(
		http.StatusOK,
		"single.html",
		gin.H{
			"title":      post.Title,
			"categories": categories,
			"post":       post,
		},
	)
}
