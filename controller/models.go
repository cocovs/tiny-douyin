package controller

//数据操作
import (
	"context"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/tencentyun/cos-go-sdk-v5"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//鉴权 搜索用户是否在users表中
//true:存在用户 false：无用户
func QueryUsersToken(token string) (User, bool) {
	var userNow User
	//DB.Where("Token = ?", token).First(&userNew)
	DB.Where("token = ?", token).First(&userNow)
	if userNow.Token != "" {
		//不为空，已有用户
		return userNow, true
	} else {
		//空 没有用户
		return userNow, false
	}

}

//上传用户信息到users表
func NewUser(token string, username string, password string) (newUser User) {
	//fmt.Println("这里是新增用户")
	newUser = User{
		Name:          username,
		FollowCount:   0,
		FollowerCount: 0,
		IsFollow:      false,
		Token:         token,
		Password:      password,
	}

	DB.Create(&newUser) //传入该结构体
	return newUser
}

//上传视频信息到videos表
func NewVideo(dataName string, userID int64) /*(newVideo Video)*/ {
	//需要传入用户id 和 视频访问url
	newVideo := Video{
		UserId:  userID,
		PlayUrl: MyOss.MyOSSUrl + "/douyin_video/" + dataName,
		//CoverUrl: ,
		ReleaseTime: time.Now().Unix(),
	}

	DB.Create(&newVideo)
	//return newVideo
}

//feed流视频列表
func NewVideoList(latest_time int64) ([]Video, int64) {

	//1.1返回小于最近时间的最近的三个视频
	var videosList []Video
	var count int
	DB.Limit(3).Where("release_time < ?", latest_time).Order("id desc").Find(&videosList)

	//2.1根据UserId从数据库中取出用户信息User结构体
	count = len(videosList)
	//遍历为每个结构体赋予对应用户结构体
	var timeNow int64
	for index, value := range videosList {
		userid := value.UserId
		//从数据库中找到用户id
		var usernow User
		DB.Where("id > ?", userid).First(&usernow)
		//分别为每个赋予用户结构体
		value.Author = usernow

		//如果该视频为视频库中最后一个视频则返回 现在时间
		if index == count-1 {
			timeNow = value.ReleaseTime
			if value.Id == 1 {
				timeNow = time.Now().Unix()
			}
		}
	}
	return videosList, timeNow
}

//获取用户发布视频列表
func UserVideoList(token string, VideoList []Video) []Video {
	//搜索数据库中所有  Video.UserId == User.Id  的
	//该用户信息
	var userNow User
	DB.Where("token = ?", token).First(&userNow)
	//该用户所有视频
	DB.Where("user_id = ?", userNow.Id).Find(&VideoList)
	//遍历加入用户信息
	for _, value := range VideoList {
		value.Author = userNow
	}
	return VideoList
}

//视频文件上传至对象存储桶内
func UpOSS(Myfile io.Reader, FileName string) error {
	//
	u, _ := url.Parse(MyOss.MyOSSUrl) //*url.URL
	b := &cos.BaseURL{BucketURL: u}   //*cos.BaseURL 访问桶列表
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  MyOss.SecId,   // 替换为用户的 SecretId，
			SecretKey: MyOss.SectKey, // 替换为用户的 SecretKey，
			//请登录访问管理控制台进行查看和管理，https://console.cloud.tencent.com/cam/capi
		},
	})
	//对象的存储位置 文件夹名/文件名
	name := "douyin_video/" + FileName
	_, err := c.Object.Put(context.Background(), name, Myfile, nil)
	return err
}
