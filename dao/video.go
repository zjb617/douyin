package dao

import (
	"sync"
	"time"

	"gorm.io/gorm"
)

type Video struct {
	Id            int64
	Uid           int64
	PlayUrl       string
	CoverUrl      string
	FavoriteCount int64
	CommentCount  int64
	Title         string
	CreateTime    int64
}

func (Video) TableName() string {
	return "video"
}

type VideoDao struct {
}

var (
	videoDao  *VideoDao
	videoOnce sync.Once
)

func NewVideoDaoInstance() *VideoDao {
	videoOnce.Do(func() {
		videoDao = &VideoDao{}
	})
	return videoDao
}

func (*VideoDao) QueryVideoByUid(uid int64) ([]*Video, error) {
	videos := make([]*Video, 0)
	err := DB.Where("uid = ?", uid).Find(&videos).Error
	return videos, err
}

func (*VideoDao) QueryVideoBefore(latestTime int64) ([]*Video, error) {
	videos := make([]*Video, 30)
	err := DB.Where("create_time < ?", latestTime).Order("create_time desc").Limit(30).Find(&videos).Error
	return videos, err
}

func (*VideoDao) InsertVideo(uid int64, playUrl, coverUrl, title string) error {
	video := Video{
		Uid:        uid,
		PlayUrl:    playUrl,
		CoverUrl:   coverUrl,
		Title:      title,
		CreateTime: time.Now().Unix(),
	}
	return DB.Create(&video).Error
}

func (*VideoDao) UpdateVideoFavoriteCountAddOne(videoId int64) error {
	return DB.Model(&Video{}).Where("id = ?", videoId).Update("favorite_count", gorm.Expr("favorite_count + 1")).Error
}

func (*VideoDao) UpdateVideoFavoriteCountMinusOne(videoId int64) error {
	return DB.Model(&Video{}).Where("id = ?", videoId).Update("favorite_count", gorm.Expr("favorite_count - 1")).Error
}

func (*VideoDao) UpdateVideoCommentCountAddOne(videoId int64) error {
	return DB.Model(&Video{}).Where("id = ?", videoId).Update("comment_count", gorm.Expr("comment_count + 1")).Error
}

func (*VideoDao) UpdateVideoCommentCountMinusOne(videoId int64) error {
	return DB.Model(&Video{}).Where("id = ?", videoId).Update("comment_count", gorm.Expr("comment_count - 1")).Error
}
