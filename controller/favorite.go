package controller

//收藏
import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
)

// FavoriteAction 点赞操作
func FavoriteAction(c *gin.Context) {
	token := c.PostForm("token")
	videoId := c.PostForm("video_id")
	actionType := c.PostForm("action_type")

	userId := isUserExists(token)
	if userId == -1 {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}

	key := "video:liked:" + videoId

	if actionType == "1" {
		DB.Model(User{}).Where("id = ?", videoId).UpdateColumn("favorite_count", gorm.Expr("favorite_count + 1"))
		RDB.SAdd(ctx, key, userId)
	} else {
		DB.Model(User{}).Where("id = ?", videoId).UpdateColumn("favorite_count", gorm.Expr("favorite_count - 1"))
		RDB.SRem(ctx, key, userId)
	}
	c.JSON(http.StatusOK, Response{StatusCode: 0})

	//var rdb redis.Client
	//err := initRedis(&rdb)
	//if err != nil {
	//	c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "redis连接失败"})
	//}

}

func FavoriteList(c *gin.Context) {
	token := c.Query("token")
	userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)

	// 查询用户是否存在
	if isUserExists(token) == -1 {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User不存在"})
	}

	var video []Video
	result := DB.Find(&video)
	var videoList []Video
	// 填充videoList
	for i := int64(0); i < result.RowsAffected; i++ {
		if isVideoLike(video[i], userId) {
			videoList = append(videoList, video[i])
		}
	}

	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: videoList,
	})
}

// 根据token查询用户是否存在
func isUserExists(token string) int64 {
	var user User
	DB.Where("token = ?", token).First(&user)

	if (user == User{}) {
		return -1
	}
	return user.Id
}

// 查询video是否被当前user点赞并设置isFavorite
func isVideoLike(video Video, userId int64) bool {
	videoId := strconv.FormatInt(video.Id, 10)
	key := "video:liked:" + videoId
	isMember := RDB.SIsMember(ctx, key, userId)
	DB.Where("videoId = ?", videoId).Update("is_favorite", isMember.Val())
	return isMember.Val()
}
