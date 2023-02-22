package service

import (
	"douyin/dao"
	"errors"
	"fmt"
	"strconv"
)

/*
Follow action
*/

type FollowActionFlow struct {
	token string
	toUid int64
}

func NewFollowActionFlow(tokenArg string, toUidArg int64) *FollowActionFlow {
	return &FollowActionFlow{
		token: tokenArg,
		toUid: toUidArg,
	}
}

func (faf *FollowActionFlow) checkParam() error {
	if !dao.RDB.HExists(dao.Ctx, "token", faf.token).Val() {
		return errors.New("token invalid")
	}
	if faf.toUid <= 0 {
		return errors.New("toUid invalid")
	}
	return nil
}

func (faf *FollowActionFlow) packInfo() error {
	myUid, err := strconv.ParseInt(dao.RDB.HGet(dao.Ctx, "token", faf.token).Val(), 10, 64)
	if err != nil {
		return err
	}
	err = dao.NewUserFollowDaoInstance().InsertUserFollow(myUid, faf.toUid)
	if err != nil {
		return err
	}
	// user table follow_count + 1 and follower_count + 1
	err = dao.NewUserDaoInstance().UpdateUserFollowCountAddOne(myUid)
	if err != nil {
		return err
	}
	err = dao.NewUserDaoInstance().UpdateUserFollowerCountAddOne(faf.toUid)
	if err != nil {
		return err
	}

	return nil
}

func (faf *FollowActionFlow) Do() error {
	if err := faf.checkParam(); err != nil {
		return err
	}
	if err := faf.packInfo(); err != nil {
		return err
	}
	return nil
}

func FollowAction(token string, toUid int64) error {
	return NewFollowActionFlow(token, toUid).Do()
}

/*
Cancel follow
*/

type CancelFollowFlow struct {
	token string
	toUid int64
}

func NewCancelFollowFlow(tokenArg string, toUidArg int64) *CancelFollowFlow {
	return &CancelFollowFlow{
		token: tokenArg,
		toUid: toUidArg,
	}
}

func (cff *CancelFollowFlow) checkParam() error {
	if !dao.RDB.HExists(dao.Ctx, "token", cff.token).Val() {
		return errors.New("token invalid")
	}
	if cff.toUid <= 0 {
		return errors.New("toUid invalid")
	}
	return nil
}

func (cff *CancelFollowFlow) packInfo() error {
	myUid, err := strconv.ParseInt(dao.RDB.HGet(dao.Ctx, "token", cff.token).Val(), 10, 64)
	if err != nil {
		return err
	}
	err = dao.NewUserFollowDaoInstance().DeleteUserFollow(myUid, cff.toUid)
	if err != nil {
		return err
	}
	// user table follow_count - 1 and follower_count - 1
	err = dao.NewUserDaoInstance().UpdateUserFollowCountMinusOne(myUid)
	if err != nil {
		return err
	}
	err = dao.NewUserDaoInstance().UpdateUserFollowerCountMinusOne(cff.toUid)
	if err != nil {
		return err
	}
	return nil
}

func (cff *CancelFollowFlow) Do() error {
	if err := cff.checkParam(); err != nil {
		return err
	}
	if err := cff.packInfo(); err != nil {
		return err
	}
	return nil
}

func CancelFollow(token string, toUid int64) error {
	return NewCancelFollowFlow(token, toUid).Do()
}

/*
Query follow list
*/

type QueryFollowListFlow struct {
	userId     int64
	token      string
	followList *FollowList

	fList []UserDetailInfo
}

type FollowList = []UserDetailInfo

func NewQueryFollowListFlow(userIdArg int64, tokenArg string) *QueryFollowListFlow {
	return &QueryFollowListFlow{
		userId: userIdArg,
		token:  tokenArg,
	}
}

func (qflf *QueryFollowListFlow) checkParam() error {
	if qflf.userId <= 0 {
		return errors.New("userId invalid")
	}
	if qflf.token != "" && !dao.RDB.HExists(dao.Ctx, "token", qflf.token).Val() {
		return errors.New("token invalid")
	}
	return nil
}

func (qflf *QueryFollowListFlow) prepareInfo() error {
	userList, err := dao.NewUserFollowDaoInstance().QueryFollowByUid(qflf.userId)
	if err != nil {
		return err
	}

	fList := make([]UserDetailInfo, 0, len(userList))
	for _, user := range userList {
		userDetailInfo := UserDetailInfo{
			Id:            user.Id,
			Name:          user.Username,
			FollowCount:   user.FollowCount,
			FollowerCount: user.FollowerCount,
			IsFollow:      true,
		}
		fList = append(fList, userDetailInfo)
	}
	qflf.fList = fList
	return nil
}

func (qflf *QueryFollowListFlow) packInfo() error {
	qflf.followList = &qflf.fList
	return nil
}

func (qflf *QueryFollowListFlow) Do() (*FollowList, error) {
	if err := qflf.checkParam(); err != nil {
		return nil, err
	}
	if err := qflf.prepareInfo(); err != nil {
		return nil, err
	}
	if err := qflf.packInfo(); err != nil {
		return nil, err
	}
	return qflf.followList, nil
}

func QueryFollowList(userId int64, token string) (*FollowList, error) {
	return NewQueryFollowListFlow(userId, token).Do()
}

/*
Query follower list
*/

type QueryFollowerListFlow struct {
	userId       int64
	token        string
	followerList *FollowerList

	fList []UserDetailInfo
}

type FollowerList = []UserDetailInfo

func NewQueryFollowerListFlow(userIdArg int64, tokenArg string) *QueryFollowerListFlow {
	return &QueryFollowerListFlow{
		userId: userIdArg,
		token:  tokenArg,
	}
}

func (qflf *QueryFollowerListFlow) checkParam() error {
	if qflf.userId <= 0 {
		return errors.New("userId invalid")
	}
	if !dao.RDB.HExists(dao.Ctx, "token", qflf.token).Val() {
		return errors.New("token invalid")
	}
	return nil
}

func (qflf *QueryFollowerListFlow) prepareInfo() error {
	myUid, err := strconv.ParseInt(dao.RDB.HGet(dao.Ctx, "token", qflf.token).Val(), 10, 64)
	if err != nil {
		return err
	}
	userList, err := dao.NewUserFollowDaoInstance().QueryUserByFollowUid(qflf.userId)
	if err != nil {
		return err
	}

	fList := make([]UserDetailInfo, 0, len(userList))
	for _, user := range userList {
		// check if I follow this user
		isFollow := false
		userFollow, err := dao.NewUserFollowDaoInstance().QueryByUidAndFollowUid(myUid, user.Id)
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
		fList = append(fList, userDetailInfo)
	}
	qflf.fList = fList
	return nil
}

func (qflf *QueryFollowerListFlow) packInfo() error {
	qflf.followerList = &qflf.fList
	return nil
}

func (qflf *QueryFollowerListFlow) Do() (*FollowerList, error) {
	if err := qflf.checkParam(); err != nil {
		return nil, err
	}
	if err := qflf.prepareInfo(); err != nil {
		return nil, err
	}
	if err := qflf.packInfo(); err != nil {
		return nil, err
	}
	return qflf.followerList, nil
}

func QueryFollowerList(userId int64, token string) (*FollowerList, error) {
	return NewQueryFollowerListFlow(userId, token).Do()
}

/*
Query friend list
*/

type QueryFriendListFlow struct {
	userId     int64
	token      string
	friendList *FriendList

	fList []UserDetailInfo
}

type FriendList = []UserDetailInfo

func NewQueryFriendListFlow(userIdArg int64, tokenArg string) *QueryFriendListFlow {
	return &QueryFriendListFlow{
		userId: userIdArg,
		token:  tokenArg,
	}
}

func (qflf *QueryFriendListFlow) checkParam() error {
	if qflf.userId <= 0 {
		return errors.New("userId invalid")
	}
	if !dao.RDB.HExists(dao.Ctx, "token", qflf.token).Val() {
		return errors.New("token invalid")
	}
	return nil
}

func (qflf *QueryFriendListFlow) prepareInfo() error {
	userList, err := dao.NewUserFollowDaoInstance().QueryFriendByUid(qflf.userId)
	if err != nil {
		return err
	}

	fList := make([]UserDetailInfo, 0, len(userList))
	for _, user := range userList {
		userDetailInfo := UserDetailInfo{
			Id:            user.Id,
			Name:          user.Username,
			FollowCount:   user.FollowCount,
			FollowerCount: user.FollowerCount,
			IsFollow:      true,
		}
		fList = append(fList, userDetailInfo)
	}
	qflf.fList = fList

	for _, friend := range fList {
		// reset "latest_msg_time" to "0", get all messages
		dao.RDB.HSet(dao.Ctx, "latest_msg_time", fmt.Sprintf("%d_%d", qflf.userId, friend.Id), 0)
		// reset "message_state" to "1", can get new/all messages
		dao.RDB.HSet(dao.Ctx, "message_state", fmt.Sprintf("%d_%d", qflf.userId, friend.Id), 1)
		// magic key, used to sync message 
		// ("/relation/friend/list"'s next request may be "/message/chat/",
		// so use this to ignore the first "/message/chat/" after "/relation/friend/list")
		dao.RDB.HSet(dao.Ctx, "magic_key", fmt.Sprintf("%d_%d", qflf.userId, friend.Id), 1)
	}

	return nil
}

func (qflf *QueryFriendListFlow) packInfo() error {
	qflf.friendList = &qflf.fList
	return nil
}

func (qflf *QueryFriendListFlow) Do() (*FriendList, error) {
	if err := qflf.checkParam(); err != nil {
		return nil, err
	}
	if err := qflf.prepareInfo(); err != nil {
		return nil, err
	}
	if err := qflf.packInfo(); err != nil {
		return nil, err
	}
	return qflf.friendList, nil
}

func QueryFriendList(userId int64, token string) (*FriendList, error) {
	return NewQueryFriendListFlow(userId, token).Do()
}
