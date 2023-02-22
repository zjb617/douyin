package service

import (
	"douyin/dao"
	"errors"
	"strconv"
	"time"
)

type GetFeedInfoFlow struct {
	latestTime int64  // -1 means lastest_time was not passed
	token      string // "" means user was not logged in
	feedInfo   *FeedInfo

	nextTime  int64       // temp for prepareInfo()
	videoList []VideoInfo // temp for prepareInfo()
}

func NewGetFeedInfoFlow(latestTimeArg int64, tokenArg string) *GetFeedInfoFlow {
	return &GetFeedInfoFlow{
		latestTime: latestTimeArg,
		token:      tokenArg,
	}
}

func (gfif *GetFeedInfoFlow) checkParam() error {
	if gfif.latestTime != -1 && gfif.latestTime < 0 {
		return errors.New("latest_time invalid")
	}
	// check token
	if gfif.token != "" && !dao.RDB.HExists(dao.Ctx, "token", gfif.token).Val() {
		return errors.New("token invalid")
	}
	return nil
}

func (gfif *GetFeedInfoFlow) prepareInfo() error {
	if gfif.latestTime == -1 {
		gfif.latestTime = time.Now().Unix()
	}
	videoList, err := dao.NewVideoDaoInstance().QueryVideoBefore(gfif.latestTime)
	if err != nil {
		return err
	}
	lenVideoList := len(videoList)
	if lenVideoList == 0 {
		videoList, err = dao.NewVideoDaoInstance().QueryVideoBefore(time.Now().Unix())
		if err != nil {
			return err
		}
		lenVideoList = len(videoList)
	}
	gfif.nextTime = videoList[lenVideoList-1].CreateTime
	videoInfoList := make([]VideoInfo, 0, lenVideoList)

	var myUid int64 = -1 // assume not logged in

	if gfif.token != "" {
		myUid, err = strconv.ParseInt(dao.RDB.HGet(dao.Ctx, "token", gfif.token).Val(), 10, 64)
		if err != nil {
			return err
		}
	}
	for _, video := range videoList {
		author, err := dao.NewUserDaoInstance().QueryUserById(video.Uid)
		if err != nil {
			return err
		}
		if author.Id == 0 {
			return errors.New("author not found")
		}

		var isFollow bool = false // assume not logged in
		if myUid != -1 {
			// check if I follow the author
			userFollow, err := dao.NewUserFollowDaoInstance().QueryByUidAndFollowUid(myUid, author.Id)
			if err != nil {
				return err
			}
			if userFollow.Id != 0 {
				isFollow = true
			}
		}

		authorDetailInfo := UserDetailInfo{
			Id:            author.Id,
			Name:          author.Username,
			FollowCount:   author.FollowCount,
			FollowerCount: author.FollowerCount,
			IsFollow:      isFollow,
		}

		var isFavorite bool = false // assume not logged in
		if myUid != -1 {
			// check if I favorite the video
			videoFavorite, err := dao.NewVideoFavoriteDaoInstance().QueryByUidAndVideoId(myUid, video.Id)
			if err != nil {
				return err
			}
			if videoFavorite.Id != 0 {
				isFavorite = true
			}
		}

		videoInfo := VideoInfo{
			Id:            video.Id,
			Author:        authorDetailInfo,
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    isFavorite,
			Title:         video.Title,
		}
		videoInfoList = append(videoInfoList, videoInfo)
	}
	gfif.videoList = videoInfoList
	return nil
}

func (gfif *GetFeedInfoFlow) packInfo() error {
	gfif.feedInfo = &FeedInfo{
		NextTime:  gfif.nextTime,
		VideoList: gfif.videoList,
	}
	return nil
}

func (gfif *GetFeedInfoFlow) Do() (*FeedInfo, error) {
	if err := gfif.checkParam(); err != nil {
		return nil, err
	}
	if err := gfif.prepareInfo(); err != nil {
		return nil, err
	}
	if err := gfif.packInfo(); err != nil {
		return nil, err
	}
	return gfif.feedInfo, nil
}

func GetFeedInfo(latestTime int64, token string) (*FeedInfo, error) {
	return NewGetFeedInfoFlow(latestTime, token).Do()
}
