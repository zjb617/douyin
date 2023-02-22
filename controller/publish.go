package controller

import (
	"context"
	"douyin/service"
	"net/http"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
)

type PublishResponse struct {
	service.BasicResponse
}

type PublishListResponse struct {
	service.BasicResponse
	VideoList []service.VideoInfo `json:"video_list"`
}

func Publish(ctx context.Context, c *app.RequestContext) {
	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, PublishResponse{
			BasicResponse: service.BasicResponse{
				StatusCode: -1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}
	if err := service.Publish(data, c.PostForm("token"), c.PostForm("title")); err != nil {
		c.JSON(http.StatusOK, PublishResponse{
			BasicResponse: service.BasicResponse{
				StatusCode: -1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}
	c.JSON(http.StatusOK, PublishResponse{
		BasicResponse: service.BasicResponse{
			StatusCode: 0,
			StatusMsg:  "ok",
		},
	})
}

func PublishList(ctx context.Context, c *app.RequestContext) {
	userId, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, PublishListResponse{
			BasicResponse: service.BasicResponse{
				StatusCode: -1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}

	videoList, err := service.QueryPublishList(c.Query("token"), userId)
	if err != nil {
		c.JSON(http.StatusOK, PublishListResponse{
			BasicResponse: service.BasicResponse{
				StatusCode: -1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, PublishListResponse{
		BasicResponse: service.BasicResponse{
			StatusCode: 0,
			StatusMsg:  "ok",
		},
		VideoList: *videoList,
	})
}
