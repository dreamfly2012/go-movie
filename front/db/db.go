package db

import (
	"go-movie/util"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

//DB is global connection
var DB *gorm.DB

//引入包的时候自动初始化
func init() {
	myConfig := new(util.Config)
	myConfig.InitConfig("config.ini")
	database := myConfig.Read("database", "db")
	username := myConfig.Read("database", "username")
	password := myConfig.Read("database", "password")
	host := myConfig.Read("database", "host")
	port := myConfig.Read("database", "port")

	//"user:password@/dbname?charset=utf8&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open("mysql", username+":"+password+"@tcp("+host+":"+port+")/"+database+"?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic("数据库连接失败")
	}

	DB.SingularTable(true)

	DB.DB().SetMaxOpenConns(10) //设置数据库连接池最大连接数
	DB.DB().SetMaxIdleConns(10)

}

//GetDb 返回全局变量DB
func GetDb() *gorm.DB {
	return DB
}
