package dao

import "sync"

type UserFollow struct {
	Id        int64
	Uid       int64
	FollowUid int64
}

func (UserFollow) TableName() string {
	return "userfollow"
}

type UserFollowDao struct {
}

var (
	userFollowDao  *UserFollowDao
	userFollowOnce sync.Once
)

func NewUserFollowDaoInstance() *UserFollowDao {
	userFollowOnce.Do(func() {
		userFollowDao = &UserFollowDao{}
	})
	return userFollowDao
}

func (*UserFollowDao) InsertUserFollow(uid, followUid int64) error {
	userFollow := UserFollow{
		Uid:       uid,
		FollowUid: followUid,
	}
	return DB.Create(&userFollow).Error
}

func (*UserFollowDao) QueryFollowByUid(uid int64) ([]*User, error) {
	users := make([]*User, 0)
	err := DB.Table("user").Select("user.*").Joins("inner join userfollow on user.id = userfollow.follow_uid").Where("userfollow.uid = ?", uid).Find(&users).Error
	return users, err
}

func (*UserFollowDao) QueryUserByFollowUid(followUid int64) ([]*User, error) {
	users := make([]*User, 0)
	err := DB.Table("user").Select("user.*").Joins("inner join userfollow on user.id = userfollow.uid").Where("userfollow.follow_uid = ?", followUid).Find(&users).Error
	return users, err
}

func (*UserFollowDao) QueryByUidAndFollowUid(uid, followUid int64) (*UserFollow, error) {
	userFollow := UserFollow{}
	err := DB.Where("uid = ? and follow_uid = ?", uid, followUid).Limit(1).Find(&userFollow).Error
	return &userFollow, err
}

func (*UserFollowDao) QueryFriendByUid(uid int64) ([]*User, error) {
	users := make([]*User, 0)
	// DB.Table("userfollow").Select("userfollow.follow_uid").Joins("inner join userfollow as uf on userfollow.follow_uid = uf.uid and uf.follow_uid = userfollow.uid").Where("userfollow.uid = ?", uid).Find(&userIds)
	err := DB.Table("user").Select("user.*").Joins("inner join userfollow on user.id = userfollow.follow_uid").Joins("inner join userfollow as uf on userfollow.follow_uid = uf.uid and uf.follow_uid = userfollow.uid").Where("userfollow.uid = ?", uid).Find(&users).Error
	return users, err
}

func (*UserFollowDao) DeleteUserFollow(uid, followUid int64) error {
	return DB.Where("uid = ? and follow_uid = ?", uid, followUid).Delete(UserFollow{}).Error
}
