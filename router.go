package main

import (
	"douyin/controller"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func registerRoute(h *server.Hertz) {
	h.Static("/static", "./public")

	apiRouter := h.Group("/douyin")

	// basic apis
	apiRouter.GET("/feed/", controller.Feed)
	apiRouter.POST("/user/register/", controller.Register)
	apiRouter.POST("/user/login/", controller.Login)
	apiRouter.GET("/user/", controller.UserInfo)
	apiRouter.POST("/publish/action/", controller.Publish)
	apiRouter.GET("/publish/list/", controller.PublishList)

	// interact apis
	apiRouter.POST("/favorite/action/", controller.FavoriteAction)
	apiRouter.GET("/favorite/list/", controller.FavoriteList)
	apiRouter.POST("/comment/action/", controller.CommentAction)
	apiRouter.GET("/comment/list/", controller.CommentList)

	// social apis
	apiRouter.POST("/relation/action/", controller.FollowAction)
	apiRouter.GET("/relation/follow/list/", controller.FollowList)
	apiRouter.GET("/relation/follower/list/", controller.FollowerList)
	apiRouter.GET("/relation/friend/list/", controller.FriendList)
	apiRouter.POST("/message/action/", controller.MessageAction)
	apiRouter.GET("/message/chat/", controller.MessageChat)
}