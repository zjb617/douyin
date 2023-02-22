package controller

import (
	"douyin/service"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

type MessageActionResponse struct {
	service.BasicResponse
}

type QueryMessageResponse struct {
	service.BasicResponse
	MessageList []service.MessageInfo `json:"message_list"`
}

func MessageAction(c *gin.Context) {
	toUserId, err := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, MessageActionResponse{
			BasicResponse: service.BasicResponse{
				StatusCode: -1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}

	actionType := c.Query("action_type") // 1: send message
	if actionType == "1" {
		err = service.SendMessage(c.Query("token"), toUserId, c.Query("content"))
		if err != nil {
			c.JSON(http.StatusOK, MessageActionResponse{
				BasicResponse: service.BasicResponse{
					StatusCode: -1,
					StatusMsg:  err.Error(),
				},
			})
			return
		} else {
			c.JSON(http.StatusOK, MessageActionResponse{
				BasicResponse: service.BasicResponse{
					StatusCode: 0,
					StatusMsg:  "ok",
				},
			})
		}
	} else {
		c.JSON(http.StatusOK, MessageActionResponse{
			BasicResponse: service.BasicResponse{
				StatusCode: -1,
				StatusMsg:  "action_type error",
			},
		})
	}
}

func MessageChat(c *gin.Context) {
	toUserId, err := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, QueryMessageResponse{
			BasicResponse: service.BasicResponse{
				StatusCode: -1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}

	messageList, err := service.QueryMessage(c.Query("token"), toUserId)
	if err != nil {
		c.JSON(http.StatusOK, QueryMessageResponse{
			BasicResponse: service.BasicResponse{
				StatusCode: -1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}
	if messageList == nil {
		c.JSON(http.StatusOK, QueryMessageResponse{
			BasicResponse: service.BasicResponse{
				StatusCode: 0,
				StatusMsg:  "message list is nil",
			},
		})
		return
	}

	c.JSON(http.StatusOK, QueryMessageResponse{
		BasicResponse: service.BasicResponse{
			StatusCode: 0,
			StatusMsg:  "ok",
		},
		MessageList: *messageList,
	})
}
