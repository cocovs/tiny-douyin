package controller

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//全局变量DB
var (
	DB *gorm.DB
)

func InitMySQL() (err error) {
	//连接数据库(加载文件config中的配置)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		MyCnf.User, MyCnf.Password, MyCnf.Host, MyCnf.Port, MyCnf.DB)

	DB, _ = gorm.Open("mysql", dsn)
	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&Video{})
	//DB.AutoMigrate(&User{})
	// var uste User
	// DB.Where(&User{Id: 1}).First(&uste)
	// fmt.Printf("%#v\n", uste)
	return DB.DB().Ping()
}
func Close() {
	DB.Close()
}
