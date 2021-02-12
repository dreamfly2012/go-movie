package controller

import (
	"fmt"
	"go-movie/db"
	"go-movie/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//Index  首页控制器
func Index(c *gin.Context) {
	db := db.GetDb()
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

//Search article search
func Search(c *gin.Context) {
	keywords := c.Query("keywords")

	db := db.GetDb()
	pagesize := 5

	var pageindex int
	var prevID int
	var nextID int
	var prevLink string
	var nextLink string
	var totalNum int

	posts := []model.Post{}

	id := c.Query("page")
	if id != "" {
		pageindex, _ = strconv.Atoi(id)
	} else {
		pageindex = 1
	}

	prevID = pageindex - 1
	if prevID == 0 {
		prevLink = "/search/?keywords=" + keywords + "&page=1"
	} else {
		prevLink = "/search/?keywords=" + keywords + "&page=" + strconv.Itoa(prevID)
	}
	nextID = pageindex + 1
	db.Model(model.Post{}).Where("title like ?", "%"+keywords+"%").Count(&totalNum)
	fmt.Println(totalNum)
	fmt.Println(pagesize)
	totalPage := totalNum / pagesize
	if totalPage == 0 {
		totalPage = 1
	}
	if nextID > totalPage {
		nextLink = "/search/?keywords=" + keywords + "&page=" + strconv.Itoa(totalPage)
	} else {
		nextLink = "/search/?keywords=" + keywords + "&page=" + strconv.Itoa(nextID)
	}

	db.Model(model.Post{}).Where("title like ?", "%"+keywords+"%").Order("id desc").Offset((pageindex - 1) * pagesize).Limit(pagesize).Find(&posts)

	categories := []model.Category{}
	db.Model(model.Category{}).Order("id desc").Find(&categories)

	tags := []model.Tag{}
	db.Model(model.Tag{}).Order("id desc").Find(&tags)

	latest := []model.Post{}
	db.Model(model.Post{}).Where("status = ?", 0).Order("id desc").Limit(3).Find(&latest)

	c.HTML(

		http.StatusOK,

		"search.html",

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

//AjaxArticle 更新阅读
func AjaxArticle(c *gin.Context) {
	db := db.GetDb()

	id := c.Param("id")
	fmt.Println(id)

	db.Model(model.Post{}).Where("id = ?", id).UpdateColumn("views", gorm.Expr("views + ?", 1))

	c.JSON(200, gin.H{
		"code": 20000,
		"data": "success",
	})
}

//Category  分类控制器
func Category(c *gin.Context) {
	db := db.GetDb()
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
		prevLink = "/category/" + category.URL + "/1"
	} else {
		prevLink = "/category/" + category.URL + "/" + strconv.Itoa(prevID)
	}
	nextID = pageindex + 1
	db.Model(model.Post{}).Where("status = ? and cid = ? ", 0, category.ID).Count(&totalNum)
	totalPage := totalNum / pagesize

	if totalNum > 0 && totalPage == 0 {
		totalPage = 1
	}

	if nextID > totalPage {
		nextLink = "/category/" + category.URL + "/" + strconv.Itoa(totalPage)
	} else {
		nextLink = "/category/" + category.URL + "/" + strconv.Itoa(nextID)
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

//Tag  标签控制器
func Tag(c *gin.Context) {
	db := db.GetDb()
	pagesize := 5

	posts := []model.Post{}
	var tag = model.Tag{}
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

	db.Model(model.Tag{}).Where("url = ?", url).First(&tag)

	fmt.Println(tag)

	prevID = pageindex - 1
	if prevID == 0 {
		prevLink = "/tag/" + tag.URL + "/1"
	} else {
		prevLink = "/tag/" + tag.URL + "/" + strconv.Itoa(prevID)
	}
	nextID = pageindex + 1
	db.Model(model.Post{}).Where("status = ? and tid = ? ", 0, tag.ID).Count(&totalNum)
	totalPage := totalNum / pagesize

	if totalNum > 0 && totalPage == 0 {
		totalPage = 1
	}

	if nextID > totalPage {
		nextLink = "/tag/" + tag.URL + "/" + strconv.Itoa(totalPage)
	} else {
		nextLink = "/tag/" + tag.URL + "/" + strconv.Itoa(nextID)
	}

	db.Model(model.Post{}).Where("status = ? and tid = ? ", 0, tag.ID).Order("id desc").Offset((pageindex - 1) * pagesize).Limit(pagesize).Find(&posts)

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
	db := db.GetDb()
	id := c.Param("id")
	//fmt.Print(id)
	post := model.Post{}
	db.Model(model.Post{}).Where("id = ?", id).First(&post)

	categories := []model.Category{}
	db.Model(model.Category{}).Order("id desc").Find(&categories)

	tags := []model.Tag{}
	db.Model(model.Tag{}).Order("id desc").Find(&tags)

	latest := []model.Post{}
	db.Model(model.Post{}).Where("status = ?", 0).Order("id desc").Limit(3).Find(&latest)

	c.HTML(
		http.StatusOK,
		"single.html",
		gin.H{
			"title":      post.Title,
			"post":       post,
			"tags":       tags,
			"latest":     latest,
			"categories": categories,
		},
	)
}
