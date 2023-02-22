package test

import (
	"douyin/dao"
	"douyin/service"
	"fmt"
	"testing"
)

func TestRegister(t *testing.T) {
	dao.Init()
	userinfo, err := service.Register("这是三十二个中文这是三十二个中文这是三十二个中文这是三十二个中文", "testpassword")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v", userinfo)
}

func TestLogin(t *testing.T) {
	dao.Init()
	userinfo, err := service.Login("testname", "testpassword")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v", userinfo)
}

func TestQueryUserDetailInfo(t *testing.T) {
	dao.Init()
	userDetailInfo, err := service.QueryUserDetailInfo(1, "testnametestpassword")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v", userDetailInfo)
}