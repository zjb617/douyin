package controller

import (
	"context"
	"douyin/service"
	"net/http"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
)

type RegisterResponse struct {
	service.BasicResponse
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}

type LoginResponse struct {
	service.BasicResponse
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}

type UserInfoResponse struct {
	service.BasicResponse
	User service.UserDetailInfo `json:"user"`
}

func Register(ctx context.Context, c *app.RequestContext) {
	username := c.Query("username")
	password := c.Query("password")
	info, err := service.Register(username, password)
	if err != nil {
		c.JSON(http.StatusOK, RegisterResponse{
			BasicResponse: service.BasicResponse{
				StatusCode: -1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, RegisterResponse{
		BasicResponse: service.BasicResponse{
			StatusCode: 0,
			StatusMsg:  "ok",
		},
		UserId: info.UserId,
		Token:  info.Token,
	})
}

func Login(ctx context.Context, c *app.RequestContext) {
	username := c.Query("username")
	password := c.Query("password")
	info, err := service.Login(username, password)
	if err != nil {
		c.JSON(http.StatusOK, LoginResponse{
			BasicResponse: service.BasicResponse{
				StatusCode: -1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, LoginResponse{
		BasicResponse: service.BasicResponse{
			StatusCode: 0,
			StatusMsg:  "ok",
		},
		UserId: info.UserId,
		Token:  info.Token,
	})
}

func UserInfo(ctx context.Context, c *app.RequestContext) {
	userId, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, UserInfoResponse{
			BasicResponse: service.BasicResponse{
				StatusCode: -1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}
	userDetailInfo, err := service.QueryUserDetailInfo(userId, c.Query("token"))
	if err != nil {
		c.JSON(http.StatusOK, UserInfoResponse{
			BasicResponse: service.BasicResponse{
				StatusCode: -1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, UserInfoResponse{
		BasicResponse: service.BasicResponse{
			StatusCode: 0,
			StatusMsg:  "ok",
		},
		User: *userDetailInfo,
	})
}
