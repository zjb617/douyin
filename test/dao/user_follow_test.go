package test

import (
	"douyin/dao"
	"fmt"
	"testing"
)

func TestInsertUserFollow(t *testing.T) {
	dao.Init()
	// for i := 2; i < 50; i++ {
	// 	err := dao.NewUserFollowDaoInstance().InsertUserFollow(1, int64(i))
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// }
	// for i := 2; i < 30; i++ {
	// 	err := dao.NewUserFollowDaoInstance().InsertUserFollow(int64(i), 1)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// }
	// for i := 100; i < 150; i++ {
	// 	err := dao.NewUserFollowDaoInstance().InsertUserFollow(int64(i), 1)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// }
	_ = dao.NewUserFollowDaoInstance().InsertUserFollow(1, 1)
}

func TestQueryFollowUidByUid(t *testing.T) {
	dao.Init()
	userIds, err := dao.NewUserFollowDaoInstance().QueryFollowByUid(1)
	if err != nil {
		fmt.Println(err)
	}
	for _, userId := range userIds {
		fmt.Println(userId)
	}
}

func TestQueryUidByFollowUid(t *testing.T) {
	dao.Init()
	userIds, err := dao.NewUserFollowDaoInstance().QueryUserByFollowUid(1)
	if err != nil {
		fmt.Println(err)
	}
	for _, userId := range userIds {
		fmt.Println(userId)
	}
}

func TestQueryFriendByUid(t *testing.T) {
	dao.Init()
	userIds, err := dao.NewUserFollowDaoInstance().QueryFriendByUid(1)
	if err != nil {
		fmt.Println(err)
	}
	for _, userId := range userIds {
		fmt.Println(userId)
	}
}

func TestQueryByUidAndFollowUid(t *testing.T) {
	dao.Init()
	userFollow, err := dao.NewUserFollowDaoInstance().QueryByUidAndFollowUid(1, 3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(userFollow)
}

func TestDeleteUserFollow(t *testing.T) {
	dao.Init()
	err := dao.NewUserFollowDaoInstance().DeleteUserFollow(1, 3)
	if err != nil {
		fmt.Println(err)
	}
}
