package test

import (
	"douyin/dao"
	"douyin/service"
	"fmt"
	"testing"
)

func TestSendMessage(t *testing.T) {
	dao.Init()
	err := service.SendMessage("testnametestpassword", 3, "hello from 1 to 3")
	if err != nil {
		fmt.Println(err)
	}
}

func TestQueryMessage(t *testing.T) {
	dao.Init()
	messageList, err := service.QueryMessage("testnametestpassword", 2)
	if err != nil {
		fmt.Println(err)
	}
	for _, message := range *messageList {
		fmt.Printf("%#v\n", message)
	}
}
