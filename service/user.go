package service

import (
	"douyin/dao"
	"errors"
	"unicode/utf8"
)

/*
 register flow start
*/

type RegisterFlow struct {
	username                string
	password                string
	userRegisterOrLoginInfo *UserRegisterOrLoginInfo

	userId int64  //temp for prepareInfo()
	token  string //temp for prepareInfo()
}

func NewRegisterFlow(usernameArg, passwordArg string) *RegisterFlow {
	return &RegisterFlow{
		username: usernameArg,
		password: passwordArg,
	}
}

func (rf *RegisterFlow) checkParam() error {
	if len1 := utf8.RuneCountInString(rf.username); rf.username == "" || len1 > 32 {
		return errors.New("username invalid")
	}
	if len2 := utf8.RuneCountInString(rf.password); rf.password == "" || len2 > 32 {
		return errors.New("password invalid")
	}
	return nil
}

func (rf *RegisterFlow) prepareInfo() error {
	// check if user exists
	user, err := dao.NewUserDaoInstance().QueryUserByUsername(rf.username)
	if err != nil {
		return err
	}
	if user.Id != 0 {
		return errors.New("user exists")
	}

	// insert user
	user, err = dao.NewUserDaoInstance().InsertUser(rf.username, rf.password)
	if err != nil {
		return err
	}
	rf.userId = user.Id
	rf.token = rf.username + rf.password
	dao.RDB.HSet(dao.Ctx, "token", rf.token, rf.userId)	
	return nil
}

func (rf *RegisterFlow) packInfo() error {
	rf.userRegisterOrLoginInfo = &UserRegisterOrLoginInfo{
		UserId: rf.userId,
		Token:  rf.token,
	}
	return nil
}

func (rf *RegisterFlow) Do() (*UserRegisterOrLoginInfo, error) {
	if err := rf.checkParam(); err != nil {
		return nil, err
	}
	if err := rf.prepareInfo(); err != nil {
		return nil, err
	}
	if err := rf.packInfo(); err != nil {
		return nil, err
	}
	return rf.userRegisterOrLoginInfo, nil
}

func Register(username, password string) (*UserRegisterOrLoginInfo, error) {
	return NewRegisterFlow(username, password).Do()
}

/*
 register flow end
*/

/*
 login flow start
*/

type LoginFlow struct {
	username string
	password string
	userInfo *UserRegisterOrLoginInfo

	userId int64  //temp for prepareInfo()
	token  string //temp for prepareInfo()
}

func NewLoginFlow(usernameArg, passwordArg string) *LoginFlow {
	return &LoginFlow{
		username: usernameArg,
		password: passwordArg,
	}
}

func (lf *LoginFlow) checkParam() error {
	if len1 := utf8.RuneCountInString(lf.username); lf.username == "" || len1 > 32 {
		return errors.New("username invalid")
	}
	if len2 := utf8.RuneCountInString(lf.password); lf.password == "" || len2 > 32 {
		return errors.New("password invalid")
	}
	return nil
}

func (lf *LoginFlow) prepareInfo() error {
	user, err := dao.NewUserDaoInstance().QueryUserByUsername(lf.username)
	if err != nil {
		return err
	}
	if user.Id == 0 {
		return errors.New("user not exists")
	}
	if user.Password != lf.password {
		return errors.New("password incorrect")
	}
	lf.userId = user.Id
	lf.token = lf.username + lf.password
	dao.RDB.HSet(dao.Ctx, "token", lf.token, lf.userId)
	return nil
}

func (lf *LoginFlow) packInfo() error {
	lf.userInfo = &UserRegisterOrLoginInfo{
		UserId: lf.userId,
		Token:  lf.token,
	}
	return nil
}

func (lf *LoginFlow) Do() (*UserRegisterOrLoginInfo, error) {
	if err := lf.checkParam(); err != nil {
		return nil, err
	}
	if err := lf.prepareInfo(); err != nil {
		return nil, err
	}
	if err := lf.packInfo(); err != nil {
		return nil, err
	}
	return lf.userInfo, nil
}

func Login(username, password string) (*UserRegisterOrLoginInfo, error) {
	return NewLoginFlow(username, password).Do()
}

/*
 login flow end
*/

/*
 user detail info flow start
*/

type QueryUserDetailInfoFlow struct {
	userId         int64
	token          string
	userDetailInfo *UserDetailInfo

	user     *dao.User //temp for prepareInfo()
	isFollow bool      //temp for prepareInfo()
}

func NewQueryUserDetailInfoFlow(userIdArg int64, tokenArg string) *QueryUserDetailInfoFlow {
	return &QueryUserDetailInfoFlow{
		userId: userIdArg,
		token:  tokenArg,
	}
}

func (qudif *QueryUserDetailInfoFlow) checkParam() error {
	if qudif.userId <= 0 {
		return errors.New("userId invalid")
	}
	if !dao.RDB.HExists(dao.Ctx, "token", qudif.token).Val() {
		return errors.New("token invalid")
	}
	return nil
}

func (qudif *QueryUserDetailInfoFlow) prepareInfo() error {
	user, err := dao.NewUserDaoInstance().QueryUserById(qudif.userId)
	if err != nil {
		return err
	}
	if user.Id == 0 {
		return errors.New("user not exists")
	}
	qudif.user = user
	userFollow, err := dao.NewUserFollowDaoInstance().QueryByUidAndFollowUid(qudif.userId, qudif.userId)
	if err != nil {
		return err
	}
	if userFollow.Id != 0 {
		qudif.isFollow = true
	}
	return nil
}

func (qudif *QueryUserDetailInfoFlow) packInfo() error {
	qudif.userDetailInfo = &UserDetailInfo{
		Id:            qudif.user.Id,
		Name:          qudif.user.Username,
		FollowCount:   qudif.user.FollowCount,
		FollowerCount: qudif.user.FollowerCount,
		IsFollow:      qudif.isFollow,
	}
	return nil
}

func (qudif *QueryUserDetailInfoFlow) Do() (*UserDetailInfo, error) {
	if err := qudif.checkParam(); err != nil {
		return nil, err
	}
	if err := qudif.prepareInfo(); err != nil {
		return nil, err
	}
	if err := qudif.packInfo(); err != nil {
		return nil, err
	}
	return qudif.userDetailInfo, nil
}

func QueryUserDetailInfo(userId int64, token string) (*UserDetailInfo, error) {
	return NewQueryUserDetailInfoFlow(userId, token).Do()
}
