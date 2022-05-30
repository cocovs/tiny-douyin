package controller

//公共的结构体

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

//视频信息
type Video struct {
	Id      int64  `json:"id,omitempty"`
	UserId  int64  `json:"user_id,omitempty"` //用户id
	Author  User   `json:"author"`
	PlayUrl string `json:"play_url,omitempty"`
	//PlayUrl       string `json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int64  `json:"favorite_count,omitempty"`
	CommentCount  int64  `json:"comment_count,omitempty"`
	//是否收藏针对每个特定的用户，因此每个用户users表中
	//应该额外有一列代表收藏的视频有哪些
	IsFavorite bool `json:"is_favorite,omitempty"`
	//title string `json:"is_favorite,omitempty"` //标题
	ReleaseTime int64 `json:"Release_Time,omitempty"`
}

//评论信息
type Comment struct {
	Id         int64  `json:"id,omitempty"`
	User       User   `json:"user"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
}

//用户信息
type User struct {
	Id            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
	//新加
	//Username string `json:"username,omitempty"`
	Token    string `json:"token,omitempty"`
	Password string `json:"password,omitempty"`
}
