package test

import (
	"douyin/dao"
	"douyin/service"
	"fmt"
	"testing"
)

func TestComment(t *testing.T) {
	dao.Init()
	comment, err := service.Comment("testnametestpassword", 1, "nice trick")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v\n",comment.User)
	fmt.Printf("%#v\n", comment)
}

func TestQueryCommentList(t *testing.T) {
	dao.Init()
	commentList, err := service.QueryCommentList("testnametestpassword", 1)
	if err != nil {
		fmt.Println(err)
	}
	for _, comment := range *commentList {
		fmt.Printf("%#v\n", comment.User)
		fmt.Printf("%#v\n", comment)
	}
}