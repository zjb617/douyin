package controller

import (
	"douyin/service"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

type FavoriteActionResponse struct {
	service.BasicResponse
}

type FavoriteListResponse struct {
	service.BasicResponse
	VideoList []service.VideoInfo `json:"video_list"`
}

func FavoriteAction(c *gin.Context) {
	videoId, err := strconv.ParseInt(c.Query("video_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, FavoriteActionResponse{
			BasicResponse: service.BasicResponse{
				StatusCode: -1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}

	actionType := c.Query("action_type") // 1: favorite, 2: cancel favorite
	if actionType == "1" {
		err = service.FavoriteAction(c.Query("token"), videoId)
		if err != nil {
			c.JSON(http.StatusOK, FavoriteActionResponse{
				BasicResponse: service.BasicResponse{
					StatusCode: -1,
					StatusMsg:  err.Error(),
				},
			})
			return
		} else {
			c.JSON(http.StatusOK, FavoriteActionResponse{
				BasicResponse: service.BasicResponse{
					StatusCode: 0,
					StatusMsg:  "ok",
				},
			})
			return
		}
	} else if actionType == "2" {
		err = service.CancelFavorite(c.Query("token"), videoId)
		if err != nil {
			c.JSON(http.StatusOK, FavoriteActionResponse{
				BasicResponse: service.BasicResponse{
					StatusCode: -1,
					StatusMsg:  err.Error(),
				},
			})
			return
		} else {
			c.JSON(http.StatusOK, FavoriteActionResponse{
				BasicResponse: service.BasicResponse{
					StatusCode: 0,
					StatusMsg:  "ok",
				},
			})
			return
		}
	} else {
		c.JSON(http.StatusOK, FavoriteActionResponse{
			BasicResponse: service.BasicResponse{
				StatusCode: -1,
				StatusMsg:  "action_type is invalid",
			},
		})
		return
	}
}

func FavoriteList(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, FavoriteListResponse{
			BasicResponse: service.BasicResponse{
				StatusCode: -1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}

	videoList, err := service.QueryFavoriteList(userId, c.Query("token"))
	if err != nil {
		c.JSON(http.StatusOK, FavoriteListResponse{
			BasicResponse: service.BasicResponse{
				StatusCode: -1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, FavoriteListResponse{
		BasicResponse: service.BasicResponse{
			StatusCode: 0,
			StatusMsg:  "ok",
		},
		VideoList: *videoList,
	})
}
