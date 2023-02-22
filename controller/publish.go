package controller

import (
	"douyin/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type PublishResponse struct {
	service.BasicResponse
}

type PublishListResponse struct {
	service.BasicResponse
	VideoList []service.VideoInfo `json:"video_list"`
}

func Publish(c *gin.Context) {
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

func PublishList(c *gin.Context) {
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
