package controller

//视频流
import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FeedResponse struct {
	Response
	//视频列表
	VideoList []Video `json:"video_list,omitempty"`
	//本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
	NextTime int64 `json:"next_time,omitempty"`
}

//投稿时间倒叙，从最近的开始30个视频
//该视频列中发布最早的视频时间 作为下次视频列表的最近时间
// 视频流
func Feed(c *gin.Context) {
	latest_time, _ := strconv.ParseInt(c.Query("latest_time"), 10, 64)
	//token := c.Query("token")
	videosList, next_time := NewVideoList(latest_time) //获取视频列表函数
	//获取发布最早时间

	
	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0}, //状态码，0-成功，其他值-失败
		VideoList: videosList,              //视频列表
		//本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
		NextTime: next_time,
	})
}
