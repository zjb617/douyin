package test

import (
	"douyin/dao"
	"fmt"
	"testing"
)

func TestInsertVideoFavorite(t *testing.T) {
	dao.Init()
	for i := 1; i < 100; i++ {
		err := dao.NewVideoFavoriteDaoInstance().InsertVideoFavorite(1, int64(260-i))
		if err != nil {
			t.Error(err)
		}
	}
}

func TestQueryVideoFavoriteByUid(t *testing.T) {
	dao.Init()
	videoIds, err := dao.NewVideoFavoriteDaoInstance().QueryVideoFavoriteByUid(1)
	if err != nil {
		t.Error(err)
	}
	for _, videoId := range videoIds {
		fmt.Printf("%#v\n", videoId)
	}
}

func TestDeleteVideoFavorite(t *testing.T) {
	dao.Init()
	for i := 0; i < 50; i++ {
		err := dao.NewVideoFavoriteDaoInstance().DeleteVideoFavorite(1, int64(i))
		if err != nil {
			t.Error(err)
		}
	}
}

func TestQueryByUidAndVideoId(t *testing.T) {
	dao.Init()
	videoFavorite, err := dao.NewVideoFavoriteDaoInstance().QueryByUidAndVideoId(1, 3)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%#v\n", videoFavorite)
}
