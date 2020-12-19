package controller

import (
	"fmt"
	"go-movie/db"
	"go-movie/model"
	"net/http"

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
	user := &model.User{}
	err := c.ShouldBindJSON(user)

	if err != nil {
		c.JSON(200, gin.H{
			"code": 20001,
			"data": err,
		})
	}

	db := db.Init()

	db.Model(model.User{}).Where("username = ? and password = ?", user.Username, user.Password).First(&user)

	fmt.Println(user)

	c.JSON(200, gin.H{
		"code":     20001,
		"username": user.Username,
		"password": user.Password,
	})
}

//Index 首页
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
