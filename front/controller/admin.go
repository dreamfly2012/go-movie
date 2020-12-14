package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//Login 登录页面
func Alogin(c *gin.Context){
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


//Index 首页
func Aindex(c *gin.Context){
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