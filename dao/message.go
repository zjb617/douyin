package dao

import (
	"sync"
	"time"
)

type Message struct {
	Id          int64
	Uid         int64
	ToUid       int64
	Content     string
	Create_time int64
}

func (Message) TableName() string {
	return "message"
}

type MessageDao struct {
}

var (
	messageDao  *MessageDao
	messageOnce sync.Once
)

func NewMessageDaoInstance() *MessageDao {
	messageOnce.Do(func() {
		messageDao = &MessageDao{}
	})
	return messageDao
}

func (*MessageDao) InsertMessage(uid, toUid int64, content string) error {
	message := Message{
		Uid:         uid,
		ToUid:       toUid,
		Content:     content,
		Create_time: time.Now().Unix(),
	}
	return DB.Create(&message).Error
}

// func (*MessageDao) QueryMessageByUidAndToUid(uid, toUid int64) ([]*Message, error) {
// 	message := make([]*Message, 0)
// 	err := DB.Where("uid = ? and to_uid = ?", uid, toUid).Or("uid = ? and to_uid = ?", toUid, uid).Find(&message).Error
// 	return message, err
// }

func (*MessageDao) QueryMessageAfterByUidAndToUid(uid, toUid, latestTime int64) ([]*Message, error) {
	message := make([]*Message, 0)
	err := DB.Where("uid = ? and to_uid = ? and create_time >= ?", uid, toUid, latestTime).Or("uid = ? and to_uid = ? and create_time >= ?", toUid, uid, latestTime).Find(&message).Error
	return message, err
}
