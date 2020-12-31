package db

import (
	"go-movie/util"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

//Init DB
func Init() *gorm.DB {
	myConfig := new(util.Config)
	myConfig.InitConfig("config.ini")
	database := myConfig.Read("database", "db")
	username := myConfig.Read("database", "username")
	password := myConfig.Read("database", "password")
	host := myConfig.Read("database", "host")
	port := myConfig.Read("database", "port")

	//"user:password@/dbname?charset=utf8&parseTime=True&loc=Local"

	db, err := gorm.Open("mysql", username+":"+password+"@tcp("+host+":"+port+")/"+database+"?charset=utf8&parseTime=True&loc=Local")
	//defer db.Close()
	if err != nil {
		panic(err)
	} else {
		// 全局禁用表名复数
		db.SingularTable(true) // 如果设置为true,`User`的默认表名为`user`,使用`TableName`设置的表名不受影响

		// 一般不会直接用CreateTable创建表
		// 检查模型`User`表是否存在，否则为模型`User`创建表
		// if !db.HasTable(&User{}) {
		// 	if err := db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{}).Error; err != nil {
		// 		panic(err)
		// 	}
		// }
	}
	
	return db
}
