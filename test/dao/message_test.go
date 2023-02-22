package test

import (
	"douyin/dao"
	"fmt"
	"testing"
)

func TestInsertMessage(t *testing.T) {
	dao.Init()
	for i := 0; i < 100; i++ {
		err := dao.NewMessageDaoInstance().InsertMessage(1, 2, "hello from 1 to 2")
		if err != nil {
			fmt.Println(err)
		}
		err = dao.NewMessageDaoInstance().InsertMessage(2, 1, "hello from 2 to 1")
		if err != nil {
			fmt.Println(err)
		}
	}
}

func TestQueryMessageAfterByUidAndToUid(t *testing.T) {
	dao.Init()
	messages, err := dao.NewMessageDaoInstance().QueryMessageAfterByUidAndToUid(1, 2, 0)
	if err != nil {
		fmt.Println(err)
	}
	for _, message := range messages {
		fmt.Println(message)
	}

	messages, err = dao.NewMessageDaoInstance().QueryMessageAfterByUidAndToUid(2, 1, 0)
	if err != nil {
		fmt.Println(err)
	}
	for _, message := range messages {
		fmt.Println(message)
	}
}
