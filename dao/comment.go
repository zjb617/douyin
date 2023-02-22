package dao

import (
	"sync"
	"time"
)

type Comment struct {
	Id         int64
	Uid        int64
	VideoId    int64
	Content    string
	CreateDate int64
}

func (Comment) TableName() string {
	return "comment"
}

type CommentDao struct {
}

var (
	commentDao  *CommentDao
	commentOnce sync.Once
)

func NewCommentDaoInstance() *CommentDao {
	commentOnce.Do(func() {
		commentDao = &CommentDao{}
	})
	return commentDao
}

func (*CommentDao) QueryCommentByVideoId(videoId int64) ([]*Comment, error) {
	comment := make([]*Comment, 0)
	err := DB.Where("video_id = ?", videoId).Order("create_date desc").Find(&comment).Error
	return comment, err
}

func (*CommentDao) InsertComment(uid, videoId int64, content string) (*Comment, error) {
	comment := Comment{
		Uid:        uid,
		VideoId:    videoId,
		Content:    content,
		CreateDate: time.Now().Unix(),
	}
	return &comment, DB.Create(&comment).Error
}

func (*CommentDao) DeleteComment(id int64) error {
	return DB.Where("id = ?", id).Delete(Comment{}).Error
}
