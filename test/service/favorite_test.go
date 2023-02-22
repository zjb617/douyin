package test

import (
	"douyin/dao"
	"douyin/service"
	"fmt"
	"testing"
)

func TestFavoriteAction(t *testing.T) {
	dao.Init()
	err := service.FavoriteAction("testnametestpassword", 3)
	if err != nil {
		fmt.Println(err)
	}
}

func TestCancelFavorite(t *testing.T) {
	dao.Init()
	err := service.CancelFavorite("testnametestpassword", 3)
	if err != nil {
		fmt.Println(err)
	}
}

func TestQueryFavoriteList(t *testing.T) {
	dao.Init()
	videoList, err := service.QueryFavoriteList(1,"testnametestpassword")
	if err != nil {
		fmt.Println(err)
	}
	for _, video := range *videoList {
		fmt.Printf("%#v\n", video.Author)
		fmt.Printf("id:%d, playUrl:%s, coverUrl:%s, favoriteCount:%d, commentCount:%d, isFavorite:%t, title:%s\n", video.Id, video.PlayUrl, video.CoverUrl, video.FavoriteCount, video.CommentCount, video.IsFavorite, video.Title)
	}
}