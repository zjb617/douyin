package service

import (
	"douyin/dao"
	"errors"
	"strconv"
)

/*
video favorite
*/

type FavoriteActionFlow struct {
	token   string
	videoId int64
}

func NewFavoriteActionFlow(tokenArg string, videoIdArg int64) *FavoriteActionFlow {
	return &FavoriteActionFlow{
		token:   tokenArg,
		videoId: videoIdArg,
	}
}

func (faf *FavoriteActionFlow) checkParam() error {
	if !dao.RDB.HExists(dao.Ctx, "token", faf.token).Val() {
		return errors.New("token invalid")
	}
	if faf.videoId <= 0 {
		return errors.New("video_id invalid")
	}
	return nil
}

func (faf *FavoriteActionFlow) packInfo() error {
	myUid, err := strconv.ParseInt(dao.RDB.HGet(dao.Ctx, "token", faf.token).Val(), 10, 64)
	if err != nil {
		return err
	}
	err = dao.NewVideoFavoriteDaoInstance().InsertVideoFavorite(myUid, faf.videoId)
	if err != nil {
		return err
	}
	// update video favorite count
	err = dao.NewVideoDaoInstance().UpdateVideoFavoriteCountAddOne(faf.videoId)
	if err != nil {
		return err
	}
	return nil
}

func (faf *FavoriteActionFlow) Do() error {
	if err := faf.checkParam(); err != nil {
		return err
	}
	if err := faf.packInfo(); err != nil {
		return err
	}
	return nil
}

func FavoriteAction(token string, videoId int64) error {
	return NewFavoriteActionFlow(token, videoId).Do()
}

/*
Cancel favorite
*/

type CancelFavoriteFlow struct {
	token   string
	videoId int64
}

func NewCancelFavoriteFlow(tokenArg string, videoIdArg int64) *CancelFavoriteFlow {
	return &CancelFavoriteFlow{
		token:   tokenArg,
		videoId: videoIdArg,
	}
}

func (cff *CancelFavoriteFlow) checkParam() error {
	if !dao.RDB.HExists(dao.Ctx, "token", cff.token).Val() {
		return errors.New("token invalid")
	}
	if cff.videoId <= 0 {
		return errors.New("video_id invalid")
	}
	return nil
}

func (cff *CancelFavoriteFlow) packInfo() error {
	myUid, err := strconv.ParseInt(dao.RDB.HGet(dao.Ctx, "token", cff.token).Val(), 10, 64)
	if err != nil {
		return err
	}
	err = dao.NewVideoFavoriteDaoInstance().DeleteVideoFavorite(myUid, cff.videoId)
	if err != nil {
		return err
	}
	// update video favorite count
	err = dao.NewVideoDaoInstance().UpdateVideoFavoriteCountMinusOne(cff.videoId)
	if err != nil {
		return err
	}
	return nil
}

func (cff *CancelFavoriteFlow) Do() error {
	if err := cff.checkParam(); err != nil {
		return err
	}
	if err := cff.packInfo(); err != nil {
		return err
	}
	return nil
}

func CancelFavorite(token string, videoId int64) error {
	return NewCancelFavoriteFlow(token, videoId).Do()
}

/*
Query favorite list
*/

type QueryFavoriteListFlow struct {
	userId       int64
	token        string
	favoriteList *FavoriteList

	fList []VideoInfo
}

type FavoriteList = []VideoInfo

func NewQueryFavoriteListFlow(userIdArg int64, tokenArg string) *QueryFavoriteListFlow {
	return &QueryFavoriteListFlow{
		userId: userIdArg,
		token:  tokenArg,
	}
}

func (qflf *QueryFavoriteListFlow) checkParam() error {
	if qflf.userId <= 0 {
		return errors.New("user_id invalid")
	}
	if !dao.RDB.HExists(dao.Ctx, "token", qflf.token).Val() {
		return errors.New("token invalid")
	}
	return nil
}

func (qflf *QueryFavoriteListFlow) prepareList() error {
	videos, err := dao.NewVideoFavoriteDaoInstance().QueryVideoFavoriteByUid(qflf.userId)
	if err != nil {
		return err
	}
	videoList := make([]VideoInfo, 0, len(videos))
	myUid, err := strconv.ParseInt(dao.RDB.HGet(dao.Ctx, "token", qflf.token).Val(), 10, 64)
	if err != nil {
		return err
	}
	for _, video := range videos {
		author, err := dao.NewUserDaoInstance().QueryUserById(video.Uid)
		if err != nil {
			return err
		}
		if author.Id == 0{
			return errors.New("author not found")
		}

		// check if I follow the author
		isFollow := false
		userFollow, err := dao.NewUserFollowDaoInstance().QueryByUidAndFollowUid(myUid, author.Id)
		if err != nil {
			return err
		}
		if userFollow.Id != 0 {
			isFollow = true
		}

		authorDetailInfo := UserDetailInfo{
			Id:            author.Id,
			Name:          author.Username,
			FollowCount:   author.FollowCount,
			FollowerCount: author.FollowerCount,
			IsFollow:      isFollow,
		}
		videoInfo := VideoInfo{
			Id:            video.Id,
			Author:        authorDetailInfo,
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    true,
			Title:         video.Title,
		}
		videoList = append(videoList, videoInfo)
	}
	qflf.fList = videoList
	return nil
}

func (qflf *QueryFavoriteListFlow) packList() error {
	qflf.favoriteList = &qflf.fList
	return nil
}

func (qflf *QueryFavoriteListFlow) Do() (*FavoriteList, error) {
	if err := qflf.checkParam(); err != nil {
		return nil, err
	}
	if err := qflf.prepareList(); err != nil {
		return nil, err
	}
	if err := qflf.packList(); err != nil {
		return nil, err
	}
	return qflf.favoriteList, nil
}

func QueryFavoriteList(userId int64, token string) (*FavoriteList, error) {
	return NewQueryFavoriteListFlow(userId, token).Do()
}
