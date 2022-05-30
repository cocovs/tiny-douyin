package controller

//加载配置文件
import (
	"fmt"
	"os"

	"gopkg.in/ini.v1"
)

type MySQLini struct {
	User     string `ini:"user"`
	Password string `ini:"password"`
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	DB       string `ini:"db"`
}

type MyOSSini struct {
	SecId    string `ini:"secretid"`
	SectKey  string `ini:"secretkey"`
	MyOSSUrl string `ini:"myossurl"`
}

var (
	MyCnf MySQLini
	MyOss MyOSSini
)

//加载配置文件
func Setini() {
	//从config文件映射到结构体
	cfg, err := ini.Load("config.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	//MapTo将文件映射到给定的结构。
	cfg.Section("mysql").MapTo(&MyCnf)
	cfg.Section("oss").MapTo(&MyOss)

}
