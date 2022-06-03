package controller

//收藏
import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"net/http"
	"time"
)

// FavoriteAction 点赞操作
func FavoriteAction(c *gin.Context) {
	token := c.PostForm("token")
	videoId := c.PostForm("video_id")
	actionType := c.PostForm("action_type")

	var user User
	DB.Where("token = ?", token).First(&user)

	if (user == User{}) {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}

	var video Video
	DB.Where("video_id = ?", videoId).First(&video)

	if video.IsFavorite = actionType == "1"; video.IsFavorite {
		video.FavoriteCount++
	} else {
		video.FavoriteCount--
	}
	c.JSON(http.StatusOK, Response{StatusCode: 0})

	//var rdb redis.Client
	//err := initRedis(&rdb)
	//if err != nil {
	//	c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "redis连接失败"})
	//}

}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	var VideoList []Video
	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: VideoList,
	})
}

func initRedis(redisDb *redis.Client) (err error) {
	redisDb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "123456",
		DB:       0,
	})
	timeoutCtx, cancelFunc := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelFunc()
	_, err = redisDb.Ping(timeoutCtx).Result()
	if err != nil {
		return err
	}
	return nil
}
