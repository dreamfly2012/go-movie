package router

import (
	"go-movie/controller"
	"go-movie/util"
	"html/template"
	"net/http"

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

	r.GET("/", controller.Index)
	r.GET("/about", controller.About)
	r.GET("/article/:id", controller.Article)
	r.GET("/page/:id", controller.Index)

	admin := r.Group("/admin")

	secrect := myConfig.Read("server", "secrect")

	store := cookie.NewStore([]byte(secrect))

	admin.Use(sessions.Sessions("mysession", store))
	admin.Use(checkAuth)

	r.GET("/admin/login", controller.Alogin)
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
