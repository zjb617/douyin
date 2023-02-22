package service

import (
	"douyin/dao"
	"errors"
	"fmt"
	"sort"
	"strconv"
	"time"
)

/*
Send message
*/

type SendMessageFlow struct {
	token   string
	toUid   int64
	content string
}

func NewSendMessageFlow(tokenArg string, toUidArg int64, contentArg string) *SendMessageFlow {
	return &SendMessageFlow{
		token:   tokenArg,
		toUid:   toUidArg,
		content: contentArg,
	}
}

func (smf *SendMessageFlow) checkParam() error {
	if !dao.RDB.HExists(dao.Ctx, "token", smf.token).Val() {
		return errors.New("token invalid")
	}
	if smf.toUid <= 0 {
		return errors.New("toUid invalid")
	}
	if len(smf.content) == 0 {
		return errors.New("content invalid")
	}
	return nil
}

func (smf *SendMessageFlow) packInfo() error {
	myUid, err := strconv.ParseInt(dao.RDB.HGet(dao.Ctx, "token", smf.token).Val(), 10, 64)
	if err != nil {
		return err
	}
	// update latest time while sending message
	dao.RDB.HSet(dao.Ctx, "latest_msg_time", fmt.Sprintf("%d_%d", myUid, smf.toUid), time.Now().Unix())
	err = dao.NewMessageDaoInstance().InsertMessage(myUid, smf.toUid, smf.content)
	if err != nil {
		return err
	}

	dao.RDB.HSet(dao.Ctx, "message_state", fmt.Sprintf("%d_%d", myUid, smf.toUid), "1") // 1: unread
	dao.RDB.HSet(dao.Ctx, "message_state", fmt.Sprintf("%d_%d", smf.toUid, myUid), "1") // 1: unread
	
	return nil
}

func (smf *SendMessageFlow) Do() error {
	if err := smf.checkParam(); err != nil {
		return err
	}
	if err := smf.packInfo(); err != nil {
		return err
	}
	return nil
}

func SendMessage(token string, toUid int64, content string) error {
	return NewSendMessageFlow(token, toUid, content).Do()
}

/*
Query message record
*/

type QueryMessageFlow struct {
	token       string
	toUid       int64
	messageList *MessageList

	mList []MessageInfo
}

type MessageList = []MessageInfo

func NewQueryMessageFlow(tokenArg string, toUidArg int64) *QueryMessageFlow {
	return &QueryMessageFlow{
		token: tokenArg,
		toUid: toUidArg,
	}
}

func (qmf *QueryMessageFlow) checkParam() error {
	if !dao.RDB.HExists(dao.Ctx, "token", qmf.token).Val() {
		return errors.New("token invalid")
	}
	if qmf.toUid <= 0 {
		return errors.New("toUid invalid")
	}
	return nil
}

func (qmf *QueryMessageFlow) prepareInfo() error {
	myUid, err := strconv.ParseInt(dao.RDB.HGet(dao.Ctx, "token", qmf.token).Val(), 10, 64)
	if err != nil {
		return err
	}

	if dao.RDB.HGet(dao.Ctx, "magic_key", fmt.Sprintf("%d_%d", myUid, qmf.toUid)).Val() == "1" { // skip this query
		dao.RDB.HSet(dao.Ctx, "magic_key", fmt.Sprintf("%d_%d", myUid, qmf.toUid), "0") // 0: not skip
		return nil
	}

	if dao.RDB.HGet(dao.Ctx, "message_state", fmt.Sprintf("%d_%d", myUid, qmf.toUid)).Val() == "1" { // 1: unread
		dao.RDB.HSet(dao.Ctx, "message_state", fmt.Sprintf("%d_%d", myUid, qmf.toUid), "0") // 0: read
	} else {
		return nil
	}

	latestMsgTime, err := strconv.ParseInt(dao.RDB.HGet(dao.Ctx, "latest_msg_time", fmt.Sprintf("%d_%d", myUid, qmf.toUid)).Val(), 10, 64)
	if err != nil {
		return err
	}

	messages, err := dao.NewMessageDaoInstance().QueryMessageAfterByUidAndToUid(myUid, qmf.toUid, latestMsgTime)
	if err != nil {
		return err
	}

	latestMsgTime, err = strconv.ParseInt(dao.RDB.HGet(dao.Ctx, "latest_msg_time", fmt.Sprintf("%d_%d", myUid, qmf.toUid)).Val(), 10, 64)
	if err != nil {
		return err
	}

	// delete repeat message while sending message(client bug, not me :). )
	if latestMsgTime != 0 {
		for i := 0; i < len(messages); i++ {
			if messages[i].Uid == myUid {
				messages = append(messages[:i], messages[i+1:]...)
				i--
			}
		}
	}

	// func means if true then i should be before j
	sort.Slice(messages, func(i, j int) bool {
		return messages[i].Create_time < messages[j].Create_time
	})

	if len1 := len(messages); len1 > 0 {
		dao.RDB.HSet(dao.Ctx, "latest_msg_time", fmt.Sprintf("%d_%d", myUid, qmf.toUid), messages[len1-1].Create_time)
	}

	mList := make([]MessageInfo, 0, len(messages))
	for _, message := range messages {
		mList = append(mList, MessageInfo{
			Id:         message.Id,
			ToUserId:   message.ToUid,
			FromUserId: message.Uid,
			Content:    message.Content,
			CreateTime: message.Create_time,
		})
	}
	qmf.mList = mList
	return nil
}

func (qmf *QueryMessageFlow) packInfo() error {
	if qmf.mList == nil {
		return nil
	}
	qmf.messageList = &qmf.mList
	return nil
}

func (qmf *QueryMessageFlow) Do() (*MessageList, error) {
	if err := qmf.checkParam(); err != nil {
		return nil, err
	}
	if err := qmf.prepareInfo(); err != nil {
		return nil, err
	}
	if err := qmf.packInfo(); err != nil {
		return nil, err
	}
	return qmf.messageList, nil
}

func QueryMessage(token string, toUid int64) (*MessageList, error) {
	return NewQueryMessageFlow(token, toUid).Do()
}
