package main


import (
	"github.com/cocovs/tiny-douyin/controller"
	"github.com/gin-gonic/gin"
	"fmt"
)


//路由入口
func initRouter(r *gin.Engine) {
	//数据库连接
	err := controller.InitMySQL()
	if err != nil {
		fmt.Println("mysql err")
		return
	}

	
	// 提供静态资源的公共服务
	r.Static("/static", "./public")
	//创建路由组
	apiRouter := r.Group("/douyin")

	// basic apis 基础接口
	apiRouter.GET("/feed/", controller.Feed)                //视频流接口
	apiRouter.GET("/user/", controller.UserInfo)            //用户信息
	apiRouter.POST("/user/register/", controller.Register)  //用户注册接口
	apiRouter.POST("/user/login/", controller.Login)        //用户登录接口
	apiRouter.POST("/publish/action/", controller.Publish)  //视频投稿
	apiRouter.GET("/publish/list/", controller.PublishList) //发布列表

	// extra apis - I 进阶接口1
	apiRouter.POST("/favorite/action/", controller.FavoriteAction)
	apiRouter.GET("/favorite/list/", controller.FavoriteList)
	apiRouter.POST("/comment/action/", controller.CommentAction)
	apiRouter.GET("/comment/list/", controller.CommentList)

	// extra apis - II 进阶接口2
	apiRouter.POST("/relation/action/", controller.RelationAction)
	apiRouter.GET("/relation/follow/list/", controller.FollowList)
	apiRouter.GET("/relation/follower/list/", controller.FollowerList)
}
