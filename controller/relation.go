package controller

import (
	"douyin/service"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

type FollowActionResponse struct {
	service.BasicResponse
}

type FollowListResponse struct {
	service.BasicResponse
	UserList []service.UserDetailInfo `json:"user_list"`
}

type FollowerListResponse struct {
	service.BasicResponse
	UserList []service.UserDetailInfo `json:"user_list"`
}

type FriendListResponse struct {
	service.BasicResponse
	UserList []service.UserDetailInfo `json:"user_list"`
}

func FollowAction(c *gin.Context) {
	toUserId, err := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
	actionType := c.Query("action_type") // 1: follow, 2: cancel follow
	if err != nil {
		c.JSON(http.StatusOK, FollowActionResponse{
			BasicResponse: service.BasicResponse{
				StatusCode: -1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}

	if actionType == "1" {
		err = service.FollowAction(c.Query("token"), toUserId)
		if err != nil {
			c.JSON(http.StatusOK, FollowActionResponse{
				BasicResponse: service.BasicResponse{
					StatusCode: -1,
					StatusMsg:  err.Error(),
				},
			})
			return
		} else {
			c.JSON(http.StatusOK, FollowActionResponse{
				BasicResponse: service.BasicResponse{
					StatusCode: 0,
					StatusMsg:  "ok",
				},
			})
			return
		}
	} else if actionType == "2" {
		err = service.CancelFollow(c.Query("token"), toUserId)
		if err != nil {
			c.JSON(http.StatusOK, FollowActionResponse{
				BasicResponse: service.BasicResponse{
					StatusCode: -1,
					StatusMsg:  err.Error(),
				},
			})
			return
		} else {
			c.JSON(http.StatusOK, FollowActionResponse{
				BasicResponse: service.BasicResponse{
					StatusCode: 0,
					StatusMsg:  "ok",
				},
			})
			return
		}
	} else {
		c.JSON(http.StatusOK, FollowActionResponse{
			BasicResponse: service.BasicResponse{
				StatusCode: -1,
				StatusMsg:  "action_type error",
			},
		})
		return
	}
}

func FollowList(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, FollowListResponse{
			BasicResponse: service.BasicResponse{
				StatusCode: -1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}

	userList, err := service.QueryFollowList(userId, c.Query("token"))
	if err != nil {
		c.JSON(http.StatusOK, FollowListResponse{
			BasicResponse: service.BasicResponse{
				StatusCode: -1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, FollowListResponse{
		BasicResponse: service.BasicResponse{
			StatusCode: 0,
			StatusMsg:  "ok",
		},
		UserList: *userList,
	})
}

func FollowerList(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, FollowerListResponse{
			BasicResponse: service.BasicResponse{
				StatusCode: -1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}

	userList, err := service.QueryFollowerList(userId, c.Query("token"))
	if err != nil {
		c.JSON(http.StatusOK, FollowerListResponse{
			BasicResponse: service.BasicResponse{
				StatusCode: -1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, FollowerListResponse{
		BasicResponse: service.BasicResponse{
			StatusCode: 0,
			StatusMsg:  "ok",
		},
		UserList: *userList,
	})
}

func FriendList(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, FriendListResponse{
			BasicResponse: service.BasicResponse{
				StatusCode: -1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}

	userList, err := service.QueryFriendList(userId, c.Query("token"))
	if err != nil {
		c.JSON(http.StatusOK, FriendListResponse{
			BasicResponse: service.BasicResponse{
				StatusCode: -1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, FriendListResponse{
		BasicResponse: service.BasicResponse{
			StatusCode: 0,
			StatusMsg:  "ok",
		},
		UserList: *userList,
	})
}
