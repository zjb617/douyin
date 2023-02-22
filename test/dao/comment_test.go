package test

import (
	"douyin/dao"
	"fmt"
	"testing"
)

func TestInsertComment(t *testing.T) {
	dao.Init()
	for i := 1; i < 100; i++ {
		comment, err := dao.NewCommentDaoInstance().InsertComment(1, int64(i%9+1), fmt.Sprintf("testcomment%d", i))
		if err != nil {
			t.Error(err)
		}
		fmt.Printf("%#v\n", comment)
	}
}

func TestQueryCommentByVideoId(t *testing.T) {
	dao.Init()
	comments, err := dao.NewCommentDaoInstance().QueryCommentByVideoId(1)
	if err != nil {
		t.Error(err)
	}
	for _, comment := range comments {
		fmt.Printf("%#v\n", comment)
	}
}

func TestDeleteComment(t *testing.T) {
	dao.Init()
	for i := 0; i < 10; i++ {
		err := dao.NewCommentDaoInstance().DeleteComment(int64(i))
		if err != nil {
			t.Error(err)
		}
	}
}
