package test

import (
	"douyin/dao"
	"douyin/service"
	"fmt"
	"testing"
)

func TestQueryPublishList(t *testing.T) {
	dao.Init()
	publishList, err := service.QueryPublishList("testnametestpassword", 2)
	if err != nil {
		t.Error(err)
	}
	for _, publish := range *publishList {
		fmt.Printf("%#v\n", publish)
	}
}