package controller

import (
	"douyin/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CommentActionResponse struct {
	service.BasicResponse
	Comment service.CommentInfo `json:"comment"`
}

type CommentListResponse struct {
	service.BasicResponse
	CommentList []service.CommentInfo `json:"comment_list"`
}

func CommentAction(c *gin.Context) {
	videoId, err := strconv.ParseInt(c.Query("video_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, CommentActionResponse{
			BasicResponse: service.BasicResponse{
				StatusCode: -1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}

	actionType := c.Query("action_type") // 1: comment, 2: cancel comment
	if actionType == "1" {
		comment, err := service.Comment(c.Query("token"), videoId, c.Query("comment_text"))
		if err != nil {
			c.JSON(http.StatusOK, CommentActionResponse{
				BasicResponse: service.BasicResponse{
					StatusCode: -1,
					StatusMsg:  err.Error(),
				},
			})
			return
		} else {
			c.JSON(http.StatusOK, CommentActionResponse{
				BasicResponse: service.BasicResponse{
					StatusCode: 0,
					StatusMsg:  "ok",
				},
				Comment: *comment,
			})
			return
		}
	} else if actionType == "2" {
		commentId, err := strconv.ParseInt(c.Query("comment_id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusOK, CommentActionResponse{
				BasicResponse: service.BasicResponse{
					StatusCode: -1,
					StatusMsg:  err.Error(),
				},
			})
			return
		}
		err = service.DeleteComment(c.Query("token"), videoId, commentId)
		if err != nil {
			c.JSON(http.StatusOK, CommentActionResponse{
				BasicResponse: service.BasicResponse{
					StatusCode: -1,
					StatusMsg:  err.Error(),
				},
			})
			return
		} else {
			c.JSON(http.StatusOK, CommentActionResponse{
				BasicResponse: service.BasicResponse{
					StatusCode: 0,
					StatusMsg:  "ok",
				},
			})
			return
		}
	} else {
		c.JSON(http.StatusOK, CommentActionResponse{
			BasicResponse: service.BasicResponse{
				StatusCode: -1,
				StatusMsg:  "action_type error",
			},
		})
		return
	}
}

func CommentList(c *gin.Context) {
	videoId, err := strconv.ParseInt(c.Query("video_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, CommentListResponse{
			BasicResponse: service.BasicResponse{
				StatusCode: -1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}

	commentList, err := service.QueryCommentList(c.Query("token"), videoId)
	if err != nil {
		c.JSON(http.StatusOK, CommentListResponse{
			BasicResponse: service.BasicResponse{
				StatusCode: -1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, CommentListResponse{
		BasicResponse: service.BasicResponse{
			StatusCode: 0,
			StatusMsg:  "ok",
		},
		CommentList: *commentList,
	})
}
