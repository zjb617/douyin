package service

import (
	"douyin/dao"
	"errors"
	"strconv"
	"time"
	"unicode/utf8"
)

type CommentFlow struct {
	token       string
	videoId     int64
	commentText string
	commentInfo *CommentInfo

	cInfo CommentInfo
}

func NewCommentFlow(tokenArg string, videoIdArg int64, commentTextArg string) *CommentFlow {
	return &CommentFlow{
		token:       tokenArg,
		videoId:     videoIdArg,
		commentText: commentTextArg,
	}
}

func (cf *CommentFlow) checkParam() error {
	if !dao.RDB.HExists(dao.Ctx, "token", cf.token).Val() {
		return errors.New("token invalid")
	}
	if cf.videoId <= 0 {
		return errors.New("video_id invalid")
	}
	if cf.commentText == "" || utf8.RuneCountInString(cf.commentText) > 1000 {
		return errors.New("comment text too long")
	}
	return nil
}

func (cf *CommentFlow) prepareInfo() error {
	myUid, err := strconv.ParseInt(dao.RDB.HGet(dao.Ctx, "token", cf.token).Val(), 10, 64)
	if err != nil {
		return err
	}

	comment, err := dao.NewCommentDaoInstance().InsertComment(myUid, cf.videoId, cf.commentText)
	if err != nil {
		return err
	}
	// update video comment count
	err = dao.NewVideoDaoInstance().UpdateVideoCommentCountAddOne(cf.videoId)
	if err != nil {
		return err
	}

	user, err := dao.NewUserDaoInstance().QueryUserById(myUid)
	if err != nil {
		return err
	}

	isFollow := false
	userFollow, err := dao.NewUserFollowDaoInstance().QueryByUidAndFollowUid(myUid, myUid)
	if err != nil {
		return err
	}

	if userFollow.Id != 0 {
		isFollow = true
	}

	userDetailInfo := UserDetailInfo{
		Id:            user.Id,
		Name:          user.Username,
		FollowCount:   user.FollowCount,
		FollowerCount: user.FollowerCount,
		IsFollow:      isFollow,
	}

	commentInfo := CommentInfo{
		Id:         comment.Id,
		User:       userDetailInfo,
		Content:    comment.Content,
		CreateDate: time.Unix(comment.CreateDate, 0).Format("01-02"),
	}

	cf.cInfo = commentInfo
	return nil
}

func (cf *CommentFlow) packInfo() error {
	cf.commentInfo = &cf.cInfo
	return nil
}

func (cf *CommentFlow) Do() (*CommentInfo, error) {
	if err := cf.checkParam(); err != nil {
		return nil, err
	}
	if err := cf.prepareInfo(); err != nil {
		return nil, err
	}
	if err := cf.packInfo(); err != nil {
		return nil, err
	}
	return cf.commentInfo, nil
}

func Comment(token string, videoId int64, commentText string) (*CommentInfo, error) {
	return NewCommentFlow(token, videoId, commentText).Do()
}

type DeleteCommentFlow struct {
	token     string
	videoId   int64
	commentId int64
}

func NewDeleteCommentFlow(tokenArg string, videoIdArg int64, commentIdArg int64) *DeleteCommentFlow {
	return &DeleteCommentFlow{
		token:     tokenArg,
		videoId:   videoIdArg,
		commentId: commentIdArg,
	}
}

func (dcf *DeleteCommentFlow) checkParam() error {
	if !dao.RDB.HExists(dao.Ctx, "token", dcf.token).Val() {
		return errors.New("token invalid")
	}
	if dcf.videoId <= 0 {
		return errors.New("video_id invalid")
	}
	if dcf.commentId <= 0 {
		return errors.New("comment_id invalid")
	}
	return nil
}

func (dcf *DeleteCommentFlow) packInfo() error {
	err := dao.NewCommentDaoInstance().DeleteComment(dcf.commentId)
	if err != nil {
		return err
	}
	return nil
}

func (dcf *DeleteCommentFlow) Do() error {
	if err := dcf.checkParam(); err != nil {
		return err
	}
	if err := dcf.packInfo(); err != nil {
		return err
	}
	return nil
}

func DeleteComment(token string, videoId int64, commentId int64) error {
	return NewDeleteCommentFlow(token, videoId, commentId).Do()
}

/*
Query comment list of a video
*/

type QueryCommentListFlow struct {
	token       string
	videoId     int64
	commentList *CommentList

	cList []CommentInfo
}

type CommentList = []CommentInfo

func NewQueryCommentListFlow(tokenArg string, videoIdArg int64) *QueryCommentListFlow {
	return &QueryCommentListFlow{
		token:   tokenArg,
		videoId: videoIdArg,
	}
}

func (qclf *QueryCommentListFlow) checkParam() error {
	if qclf.token != "" && !dao.RDB.HExists(dao.Ctx, "token", qclf.token).Val() {
		return errors.New("token invalid")
	}
	if qclf.videoId <= 0 {
		return errors.New("video_id invalid")
	}
	return nil
}

func (qclf *QueryCommentListFlow) prepareInfo() error {
	var myUid int64 = -1 // not login
	var err error
	if qclf.token != "" {
		myUid, err = strconv.ParseInt(dao.RDB.HGet(dao.Ctx, "token", qclf.token).Val(), 10, 64)
		if err != nil {
			return err
		}
	}

	comments, err := dao.NewCommentDaoInstance().QueryCommentByVideoId(qclf.videoId)
	if err != nil {
		return err
	}

	cList := make([]CommentInfo, 0, len(comments))

	for _, comment := range comments {
		author, err := dao.NewUserDaoInstance().QueryUserById(comment.Uid)
		if err != nil {
			return err
		}
		if author.Id == 0 {
			return errors.New("author not found")
		}

		isFollow := false
		if myUid != -1 {
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

		commentInfo := CommentInfo{
			Id:         comment.Id,
			User:       authorDetailInfo,
			Content:    comment.Content,
			CreateDate: time.Unix(comment.CreateDate, 0).Format("01-02"),
		}
		cList = append(cList, commentInfo)
	}
	qclf.cList = cList
	return nil
}

func (qclf *QueryCommentListFlow) packInfo() error {
	qclf.commentList = &qclf.cList
	return nil
}

func (qclf *QueryCommentListFlow) Do() (*CommentList, error) {
	if err := qclf.checkParam(); err != nil {
		return nil, err
	}
	if err := qclf.prepareInfo(); err != nil {
		return nil, err
	}
	if err := qclf.packInfo(); err != nil {
		return nil, err
	}
	return qclf.commentList, nil
}

func QueryCommentList(token string, videoId int64) (*CommentList, error) {
	return NewQueryCommentListFlow(token, videoId).Do()
}
