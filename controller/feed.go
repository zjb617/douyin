package controller

import (
	// "context"
	"context"
	"douyin/service"
	"net/http"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	// "github.com/cloudwego/hertz/pkg/app"
)

type FeedResponse struct {
	service.BasicResponse
	NextTime  int64               `json:"next_time"`
	VideoList []service.VideoInfo `json:"video_list"`
}

func Feed(ctx context.Context, c *app.RequestContext) {
	var latestTime int64
	var err error
	if temp := c.Query("latest_time"); temp == "0" {
		latestTime = -1
	} else {
		latestTime, err = strconv.ParseInt(c.Query("latest_time"), 10, 64)
		if err != nil {
			c.JSON(http.StatusOK, FeedResponse{
				BasicResponse: service.BasicResponse{
					StatusCode: -1,
					StatusMsg:  err.Error(),
				},
			})
			return
		}
	}

	feed, err := service.GetFeedInfo(latestTime, c.Query("token"))
	if err != nil {
		c.JSON(http.StatusOK, FeedResponse{
			BasicResponse: service.BasicResponse{
				StatusCode: -1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, FeedResponse{
		BasicResponse: service.BasicResponse{
			StatusCode: 0,
			StatusMsg:  "ok",
		},
		NextTime:  feed.NextTime,
		VideoList: feed.VideoList,
	})
}
