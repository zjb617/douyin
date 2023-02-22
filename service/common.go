package service

type BasicResponse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

type UserRegisterOrLoginInfo struct {
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}

type FeedInfo struct {
	NextTime  int64       `json:"next_time"`
	VideoList []VideoInfo `json:"video_list"`
}

type VideoInfo struct {
	Id            int64          `json:"id"`
	Author        UserDetailInfo `json:"author"`
	PlayUrl       string         `json:"play_url"`
	CoverUrl      string         `json:"cover_url"`
	FavoriteCount int64          `json:"favorite_count"`
	CommentCount  int64          `json:"comment_count"`
	IsFavorite    bool           `json:"is_favorite"`
	Title         string         `json:"title"`
}

type UserDetailInfo struct {
	Id            int64  `json:"id"`
	Name          string `json:"name"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}

type CommentInfo struct {
	Id         int64          `json:"id"`
	User       UserDetailInfo `json:"user"`
	Content    string         `json:"content"`
	CreateDate string         `json:"create_date"`
}

type MessageInfo struct {
	Id         int64  `json:"id"`
	ToUserId   int64  `json:"to_user_id"`
	FromUserId int64  `json:"from_user_id"`
	Content    string `json:"content"`
	CreateTime int64  `json:"create_time"`
}
