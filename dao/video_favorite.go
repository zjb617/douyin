package dao

import "sync"

type VideoFavorite struct {
	Id      int64
	Uid     int64
	VideoId int64
}

func (VideoFavorite) TableName() string {
	return "videofavorite"
}

type VideoFavoriteDao struct {
}

var (
	videoFavoriteDao  *VideoFavoriteDao
	videoFavoriteOnce sync.Once
)

func NewVideoFavoriteDaoInstance() *VideoFavoriteDao {
	videoFavoriteOnce.Do(func() {
		videoFavoriteDao = &VideoFavoriteDao{}
	})
	return videoFavoriteDao
}

func (*VideoFavoriteDao) QueryVideoFavoriteByUid(uid int64) ([]*Video, error) {
	videos := make([]*Video, 0)
	err := DB.Table("video").Select("video.*").Joins("inner join videofavorite on video.id = videofavorite.video_id").Where("videofavorite.uid = ?", uid).Find(&videos).Error
	return videos, err
}

func (*VideoFavoriteDao) QueryByUidAndVideoId(uid, videoId int64) (*VideoFavorite, error) {
	videoFavorite := VideoFavorite{}
	err := DB.Where("uid = ? and video_id = ?", uid, videoId).Limit(1).Find(&videoFavorite).Error
	return &videoFavorite, err
}

func (*VideoFavoriteDao) InsertVideoFavorite(uid, videoId int64) error {
	videoFavorite := VideoFavorite{
		Uid:     uid,
		VideoId: videoId,
	}
	return DB.Create(&videoFavorite).Error
}

func (*VideoFavoriteDao) DeleteVideoFavorite(uid, videoId int64) error {
	return DB.Where("uid = ? and video_id = ?", uid, videoId).Delete(VideoFavorite{}).Error
}
