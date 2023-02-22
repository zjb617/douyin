package test

import (
	"douyin/dao"
	"fmt"
	"testing"
	"time"
)

func TestInsertVideo(t *testing.T) {
	dao.Init()

	for j := 0; j < 1; j++ {
		for i := 1; i < 2; i++ {
			err := dao.NewVideoDaoInstance().InsertVideo(
				int64(i),
				"https://www.w3schools.com/html/movie.mp4",
				"https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
				fmt.Sprintf("testtitle%d", i),
			)
			if err != nil {
				t.Error(err)
			}
			time.Sleep(time.Millisecond * 100)
		}
	}
}

func TestQueryVideoByUid(t *testing.T) {
	dao.Init()
	videos, err := dao.NewVideoDaoInstance().QueryVideoByUid(1)
	if err != nil {
		t.Error(err)
	}
	for _, video := range videos {
		fmt.Printf("%#v\n", video)
	}
}

func TestQueryVideoBefore(t *testing.T) {
	dao.Init()
	videos, err := dao.NewVideoDaoInstance().QueryVideoBefore(1676108955)
	if err != nil {
		t.Error(err)
	}
	i := 1
	for _, video := range videos {
		fmt.Printf("%d.%#v\n", i, video)
		i++
	}
}

func TestUpdateVideoFavoriteCountAddOne(t *testing.T) {
	dao.Init()
	err := dao.NewVideoDaoInstance().UpdateVideoFavoriteCountAddOne(1)
	if err != nil {
		t.Error(err)
	}
}

func TestUpdateVideoFavoriteCountMinusOne(t *testing.T) {
	dao.Init()
	err := dao.NewVideoDaoInstance().UpdateVideoFavoriteCountMinusOne(1)
	if err != nil {
		t.Error(err)
	}
}

func TestUpdateVideoCommentCountAddOne(t *testing.T) {
	dao.Init()
	err := dao.NewVideoDaoInstance().UpdateVideoCommentCountAddOne(1)
	if err != nil {
		t.Error(err)
	}
}

func TestUpdateVideoCommentCountMinusOne(t *testing.T) {
	dao.Init()
	err := dao.NewVideoDaoInstance().UpdateVideoCommentCountMinusOne(1)
	if err != nil {
		t.Error(err)
	}
}
