package dao

import (
	"sync"

	"gorm.io/gorm"
)

type User struct {
	Id            int64  //`json:"id"`
	Username      string //`json:"username"`
	Password      string //`json:"password"`
	FollowCount   int64  //`json:"follow_count"`
	FollowerCount int64  //`json:"follower_count"`
}

func (User) TableName() string {
	return "user"
}

type UserDao struct {
}

var (
	userDao  *UserDao
	userOnce sync.Once
)

func NewUserDaoInstance() *UserDao {
	userOnce.Do(func() {
		userDao = &UserDao{}
	})
	return userDao
}

func (*UserDao) QueryUserById(id int64) (*User, error) {
	user := User{}
	err := DB.Where("id = ?", id).Limit(1).Find(&user).Error
	return &user, err
}

func (*UserDao) QueryUserByUsername(username string) (*User, error) {
	user := User{}
	err := DB.Where("username = ?", username).Limit(1).Find(&user).Error
	return &user, err
}

func (*UserDao) InsertUser(username, password string) (*User, error) {
	user := User{
		Username: username,
		Password: password,
	}
	return &user, DB.Create(&user).Error
}

func (*UserDao) UpdateUserFollowCountAddOne(id int64) error {
	return DB.Model(&User{}).Where("id = ?", id).Update("follow_count", gorm.Expr("follow_count + 1")).Error
}

func (*UserDao) UpdateUserFollowCountMinusOne(id int64) error {
	return DB.Model(&User{}).Where("id = ?", id).Update("follow_count", gorm.Expr("follow_count - 1")).Error
}

func (*UserDao) UpdateUserFollowerCountAddOne(id int64) error {
	return DB.Model(&User{}).Where("id = ?", id).Update("follower_count", gorm.Expr("follower_count + 1")).Error
}

func (*UserDao) UpdateUserFollowerCountMinusOne(id int64) error {
	return DB.Model(&User{}).Where("id = ?", id).Update("follower_count", gorm.Expr("follower_count - 1")).Error
}
