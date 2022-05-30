package main

import (
	"github.com/cocovs/tiny-douyin/controller"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {

	controller.Setini()

	//主函数最后关闭数据库，否则可能会意外关闭
	defer controller.Close()

	r := gin.Default()

	initRouter(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
