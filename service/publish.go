package service

import (
	"douyin/dao"
	"errors"
	"fmt"
	"mime/multipart"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"
)

type PublishFlow struct {
	data  *multipart.FileHeader
	token string
	title string
}

func NewPublishFlow(dataArg *multipart.FileHeader, tokenArg, titleArg string) *PublishFlow {
	return &PublishFlow{
		data:  dataArg,
		token: tokenArg,
		title: titleArg,
	}
}

func videoCompress(src, dst string) {
	// using ffmpeg to compress video
	args := []string{"-i", "./public/" + src,"-b:v", "1M", "-s", "640x360", "-r", "24", "./public/" + dst}
	cmd := exec.Command("ffmpeg", args...)
	if err := cmd.Run(); err != nil {
		return
	}
	if err := os.Remove("./public/" + src); err != nil {
		return
	}
}

func (pf *PublishFlow) checkParam() error {
	if pf.token == "" || !dao.RDB.HExists(dao.Ctx, "token", pf.token).Val() {
		return errors.New("token invalid")
	}
	if pf.title == "" {
		return errors.New("title invalid")
	}
	return nil
}

func (pf *PublishFlow) packVideo() error {
	filename := filepath.Base(pf.data.Filename)
	userId, err := strconv.ParseInt(dao.RDB.HGet(dao.Ctx, "token", pf.token).Val(), 10, 64)
	if err != nil {
		return err
	}
	finalName := fmt.Sprintf("%d_%s", userId, filename)
	var c *app.RequestContext
	finalName = strings.Replace(finalName, " ", "_", -1)
	if err := c.SaveUploadedFile(pf.data, "./public/"+finalName); err != nil {
		return err
	}
	go videoCompress(finalName, "compress_"+finalName)
	err = dao.NewVideoDaoInstance().InsertVideo(userId, "http://101.34.36.126:8081/video/"+"compress_"+finalName, "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg", pf.title)
	if err != nil {
		return err
	}
	return nil
}

func (pf *PublishFlow) Do() error {
	if err := pf.checkParam(); err != nil {
		return err
	}
	if err := pf.packVideo(); err != nil {
		return err
	}
	return nil
}

func Publish(data *multipart.FileHeader, token, title string) error {
	return NewPublishFlow(data, token, title).Do()
}

/*
QueryPublishListFlow start
*/

type QueryPublishListFlow struct {
	token       string
	userId      int64
	publishList *PublishList

	videoAuthor UserDetailInfo
	pubList     []VideoInfo
}

type PublishList = []VideoInfo

func NewQueryPublishListFlow(token string, userId int64) *QueryPublishListFlow {
	return &QueryPublishListFlow{
		token:  token,
		userId: userId,
	}
}

func (qplf *QueryPublishListFlow) checkParam() error {
	if !dao.RDB.HExists(dao.Ctx, "token", qplf.token).Val() {
		return errors.New("token invalid")
	}
	if qplf.userId <= 0 {
		return errors.New("userId invalid")
	}
	return nil
}

func (qplf *QueryPublishListFlow) prepareList() error {
	myUid, err := strconv.ParseInt(dao.RDB.HGet(dao.Ctx, "token", qplf.token).Val(), 10, 64)
	if err != nil {
		return err
	}

	var isFollow bool

	// check if I follow this user
	userFollow, err := dao.NewUserFollowDaoInstance().QueryByUidAndFollowUid(myUid, qplf.userId)
	if err != nil {
		return err
	}
	if userFollow.Id != 0 {
		isFollow = true
	}

	// fill video author info
	videoAuthor, err := dao.NewUserDaoInstance().QueryUserById(qplf.userId)
	if err != nil {
		return err
	}
	if videoAuthor.Id == 0 {
		return errors.New("user not found")
	}
	qplf.videoAuthor = UserDetailInfo{
		Id:            videoAuthor.Id,
		Name:          videoAuthor.Username,
		FollowCount:   videoAuthor.FollowCount,
		FollowerCount: videoAuthor.FollowerCount,
		IsFollow:      isFollow,
	}

	// fill pubList
	videoList, err := dao.NewVideoDaoInstance().QueryVideoByUid(qplf.userId)
	if err != nil {
		return err
	}
	lenVideoList := len(videoList)
	pubList := make(PublishList, 0, lenVideoList)
	for _, video := range videoList {
		var isFavorite bool = false
		// check if I favorite this video
		videoFavorite, err := dao.NewVideoFavoriteDaoInstance().QueryByUidAndVideoId(myUid, video.Id)
		if err != nil {
			return err
		}
		if videoFavorite.Id != 0 {
			isFavorite = true
		}
		videoInfo := VideoInfo{
			Id:            video.Id,
			Author:        qplf.videoAuthor,
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    isFavorite,
			Title:         video.Title,
		}
		pubList = append(pubList, videoInfo)
	}
	qplf.pubList = pubList
	return nil
}

func (qplf *QueryPublishListFlow) packList() error {
	qplf.publishList = &qplf.pubList
	return nil
}

func (qplf *QueryPublishListFlow) Do() (*PublishList, error) {
	if err := qplf.checkParam(); err != nil {
		return nil, err
	}
	if err := qplf.prepareList(); err != nil {
		return nil, err
	}
	if err := qplf.packList(); err != nil {
		return nil, err
	}
	return qplf.publishList, nil
}

func QueryPublishList(token string, userId int64) (*PublishList, error) {
	return NewQueryPublishListFlow(token, userId).Do()
}

/*
QueryPublishListFlow end
*/
