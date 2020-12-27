package controller

import (
	"context"
	"fmt"
	"go-movie/db"
	"go-movie/model"
	"go-movie/util"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tencentyun/cos-go-sdk-v5"
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
			"code":    20001,
			"data":    "",
			"message": err,
		})
	}

	db := db.Init()

	result := db.Model(model.Admin{}).Where("username = ? and password = ?", admin.Username, admin.Password).First(&user).RecordNotFound()

	if result == true {
		c.JSON(200, gin.H{
			"code":    20001,
			"message": "用户名密码错误",
			"data":    "",
		})
	} else {
		c.JSON(200, gin.H{
			"code":    20000,
			"message": "success",
			"data":    "",
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

//CategoryList 分类列表
func CategoryList(c *gin.Context) {
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
		"code":    20000,
		"data":    category,
		"message": "success",
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

//PostGet 文章获取
func PostGet(c *gin.Context) {
	post := &model.Post{}

	id := c.DefaultQuery("id", "1")

	db := db.Init()

	db.Model(model.Post{}).Where("id = ? ", id).First(&post)

	c.JSON(200, gin.H{
		"code":    20000,
		"data":    post,
		"message": "success",
	})

}

//PostAdd 文章添加
func PostAdd(c *gin.Context) {
	post := &model.Post{}

	err := c.ShouldBindJSON(post)

	if err != nil {
		c.JSON(200, gin.H{
			"code": 20001,
			"data": err,
		})
	}

	db := db.Init()

	db.Model(model.Post{}).Create(post)

	c.JSON(200, gin.H{
		"code": 20000,
		"data": "success",
	})
}

//PostEdit 文章修改
func PostEdit(c *gin.Context) {
	post := &model.Post{}

	err := c.ShouldBindJSON(post)

	if err != nil {
		c.JSON(200, gin.H{
			"code": 20001,
			"data": err,
		})

		return
	}

	db := db.Init()

	db.Model(model.Post{}).Omit("id").Where("id = ? ", post.ID).Updates(post)

	c.JSON(200, gin.H{
		"code": 20000,
		"data": "success",
	})

}

//Upload 文件上传
func Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err == nil {
		Path := "./upload"
		t := time.Now()
		date := t.Format("20060102")
		pathTmp := Path + "/" + date + "/"
		if isDirExists(pathTmp) {
			fmt.Println("目录存在")
		} else {
			fmt.Println("目录不存在")
			err := os.Mkdir(pathTmp, 0777)
			if err != nil {
				//log.Fatal(err)
				c.JSON(200, gin.H{"code": 200001, "msg": "创建目录失败"})
			}
		}
		//文件名
		fileName := strconv.FormatInt(time.Now().Unix(), 10) + strconv.Itoa(rand.Intn(999999-100000)+100000) + path.Ext(file.Filename)
		uperr := c.SaveUploadedFile(file, pathTmp+fileName)

		fmt.Println(fileName)
		fmt.Println(pathTmp)

		url := uploadYun(fileName, pathTmp+fileName)

		if uperr == nil {
			c.JSON(200, gin.H{"code": 20000, "msg": "上传成功", "data": url})
		} else {
			c.JSON(200, gin.H{"code": 20002, "msg": "上传失败"})
		}

	} else {
		c.JSON(200, gin.H{"code": 20003, "msg": "上传失败"})
	}

}

//目录是否存在
func isDirExists(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func uploadYun(filename string, filepath string) string {
	// 将 examplebucket-1250000000 和 COS_REGION 修改为真实的信息
	myConfig := new(util.Config)
	myConfig.InitConfig("config.ini")
	bucket := myConfig.Read("server", "bucket")
	region := myConfig.Read("server", "region")
	SECRETID := myConfig.Read("server", "COS_SECRETID")
	SECRETKEY := myConfig.Read("server", "COS_SECRETKEY")
	u, _ := url.Parse("https://" + bucket + ".cos." + region + ".myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  SECRETID,
			SecretKey: SECRETKEY,
		},
	})
	// 对象键（Key）是对象在存储桶中的唯一标识。
	// 例如，在对象的访问域名 `examplebucket-1250000000.cos.COS_REGION.myqcloud.com/test/objectPut.go` 中，对象键为 test/objectPut.go

	_, err := c.Object.PutFromFile(context.Background(), filename, filepath, nil)
	if err != nil {
		panic(err)
	}

	return "https://" + bucket + ".cos." + region + ".myqcloud.com/" + filename
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
