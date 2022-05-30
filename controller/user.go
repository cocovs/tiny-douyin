package controller

//用户
import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 用户信息   key 为用户名++密码
// test data: username=zhanglei, password=douyin
var usersLoginInfo = map[string]User{
	"zhangleidouyin": {
		Id:            1,
		Name:          "zhanglei",
		FollowCount:   10,
		FollowerCount: 5,
		IsFollow:      true,
	},
}

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User User `json:"user"`
}

//注册接口
func Register(c *gin.Context) {
	//Query返回参数值，如果不存在的话则返回空字符串""
	username := c.Query("username")
	password := c.Query("password")

	token := username + password

	_, exist := QueryUsersToken(token)
	if exist {
		//数据库中找到token 用户存在
		c.JSON(http.StatusOK, UserLoginResponse{
			//用户已存在
			Response: Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
	} else {
		//不存在 添加新用户传入name 和 password
		NowId := NewUser(token, username, password).Id
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   NowId,
			Token:    username + password,
		})
	}

}

//登录
func Login(c *gin.Context) {
	//登录接口 拿到用户的用户名和密码生成 token 在 users数据库中搜索
	username := c.Query("username")
	password := c.Query("password")
	token := username + password

	var userNow User
	userNow, exist := QueryUsersToken(token)

	if exist {
		//存在
		fmt.Println("exist = true")
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   userNow.Id,
			Token:    token,
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}

}

///douyin/user/
//登陆后用户鉴权一次再进入主页
//登录：客户端将用户名＋密码组合成token 发送过来
func UserInfo(c *gin.Context) {
	token := c.Query("token")
	//Query返回参数值，如果不存在的话则返回空字符串""
	userNow, exist := QueryUsersToken(token)
	if exist { //用户存在
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 0},
			User:     userNow,
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}

}
