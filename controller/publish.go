package controller

//发布
import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type VideoListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}

//用户发布视频
func Publish(c *gin.Context) {
	token := c.PostForm("token")

	//拿到用户信息
	UserOK, exist := QueryUsersToken(token)
	if !exist {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	//拿到文件
	file, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	myfile, _ := file.Open()

	FileName := file.Filename
	//上传视频
	err = UpOSS(myfile, FileName)

	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
	} else {
		//上传视频信息至mysql
		NewVideo(file.Filename, UserOK.Id)

		c.JSON(http.StatusOK, Response{
			StatusCode: 0,
			StatusMsg:  file.Filename + " uploaded successfully  wuhu！！！！",
		})
	}

}

//根据每个用户的信息来返回视频列表 请求参数token user_id
func PublishList(c *gin.Context) {
	//c.Query()
	token := c.Query("token")
	var VideoList []Video
	VideoList = UserVideoList(token, VideoList)

	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: VideoList,
	})
}
