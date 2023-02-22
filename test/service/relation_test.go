package test

import (
	"douyin/dao"
	"douyin/service"
	"fmt"
	"testing"
)

func TestFollowAction(t *testing.T) {
	dao.Init()
	err := service.FollowAction("testnametestpassword", 3)
	if err != nil {
		fmt.Println(err)
	}
}

func TestCancelFollow(t *testing.T) {
	dao.Init()
	err := service.CancelFollow("testnametestpassword", 3)
	if err != nil {
		fmt.Println(err)
	}
}

func TestQueryFollowList(t *testing.T) {
	dao.Init()
	userList, err := service.QueryFollowList(1, "testnametestpassword")
	if err != nil {
		fmt.Println(err)
	}
	for _, user := range *userList {
		fmt.Printf("%#v\n", user)
	}
}

func TestQueryFollowerList(t *testing.T) {
	dao.Init()
	userList, err := service.QueryFollowerList(1, "testnametestpassword")
	if err != nil {
		fmt.Println(err)
	}
	for _, user := range *userList {
		fmt.Printf("%#v\n", user)
	}
}

func TestQueryFriendList(t *testing.T) {
	dao.Init()
	userList, err := service.QueryFriendList(1, "testnametestpassword")
	if err != nil {
		fmt.Println(err)
	}
	for _, user := range *userList {
		fmt.Printf("%#v\n", user)
	}
}
