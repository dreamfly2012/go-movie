package router

import (
	"go-movie/controller"
	"go-movie/util"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Init 路由初始化
func Init() {
	r := gin.Default()
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	r.LoadHTMLGlob("templates/**/*")
	r.StaticFS("/static", http.Dir("./static"))
	r.StaticFS("/ueditor", http.Dir("./ueditor"))

	r.GET("/", controller.Index)
	r.GET("/about", controller.About)
	r.GET("/article/:id", controller.Article)
	r.GET("/page/:id", controller.Index)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	myConfig := new(util.Config)
	myConfig.InitConfig("config.ini")
	host := myConfig.Read("server", "host")
	port := myConfig.Read("server", "port")

	//r.Use(TlsHandler())
	r.Run(host + ":" + port)
}
