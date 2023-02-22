package test

import (
	"testing"
	"fmt"
	"douyin/dao"
	"douyin/service"
)

func TestGetFeedInfo(t *testing.T) {
	dao.Init()
	feedInfo, err := service.GetFeedInfo(1676261653, "testnametestpassword")
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%#v\n", feedInfo.NextTime)
	for _, video := range feedInfo.VideoList {
		fmt.Printf("%#v\n", video.Author)
		fmt.Printf("id:%d, playUrl:%s, coverUrl:%s, favoriteCount:%d, commentCount:%d, isFavorite:%t, title:%s\n", video.Id, video.PlayUrl, video.CoverUrl, video.FavoriteCount, video.CommentCount, video.IsFavorite, video.Title)
	}
}