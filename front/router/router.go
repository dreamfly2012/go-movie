package router

import (
	"go-movie/controller"
	"go-movie/util"
	"html/template"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

//Init 路由初始化
func Init() {
	myConfig := new(util.Config)
	myConfig.InitConfig("config.ini")
	r := gin.Default()
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})

	r.LoadHTMLGlob("templates/**/*")
	r.StaticFS("/static", http.Dir("./static"))
	r.StaticFS("/ueditor", http.Dir("./ueditor"))
	r.StaticFS("/upload", http.Dir("./upload"))

	//前端页面路由
	r.GET("/", controller.Index)
	r.GET("/about", controller.About)
	r.GET("/article/:id", controller.Article)
	r.GET("/page/:id", controller.Index)
	r.GET("/category/:url/:id", controller.Category)

	admin := r.Group("/admin")

	secrect := myConfig.Read("server", "secrect")

	store := cookie.NewStore([]byte(secrect))

	admin.Use(sessions.Sessions("mysession", store))
	admin.Use(checkAuth)

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true                                                                                                 //允许所有域名
	config.AllowMethods = []string{"GET", "POST", "OPTIONS"}                                                                      //允许请求的方法
	config.AllowHeaders = []string{"tus-resumable", "upload-length", "upload-metadata", "cache-control", "x-requested-with", "*"} //允许的Header
	r.Use(cors.New(config))

	//后台接口路由

	r.GET("/admin/login", controller.Alogin)
	r.POST("/admin/login", controller.APlogin)
	r.GET("/admin/info", controller.AInfo)
	r.GET("/admin/list", controller.AList)
	r.POST("/admin/category/add", controller.CategoryAdd)
	r.POST("/admin/category/edit", controller.CategoryEdit)
	r.POST("/admin/post/add", controller.PostAdd)
	r.POST("/admin/post/edit", controller.PostEdit)
	r.GET("/admin/post/get", controller.PostGet)
	r.GET("/admin/category/get", controller.CategoryGet)
	r.GET("/admin/category/list", controller.CategoryList)
	r.POST("/admin/post/upload", controller.Upload)

	admin.GET("/index", controller.Aindex)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	host := myConfig.Read("server", "host")
	port := myConfig.Read("server", "port")

	//r.Use(TlsHandler())
	r.Run(host + ":" + port)
}

func checkAuth(c *gin.Context) {
	session := sessions.Default(c)
	login := session.Get("login")

	if login != nil {
		c.Next()
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
	}
	//fmt.Println(id)

	//验证成功 调用c.Next()
	//验证失败 调用c.Abort()
}
