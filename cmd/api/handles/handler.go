package handlers

import (
	"douyin/pkg/errno"
	"github.com/gin-gonic/gin"
	"net/http"
)

// QueryUserParam req format for get_user_info
type QueryUserParam struct {
	UserID int64  `json:"user_id" form:"user_id"`
	Token  string `json:"token" form:"token"`
}

type UserInfo struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// QueryUserResponse resp format of get_user_info
type QueryUserResponse struct {
	StatusCode int32    `json:"status_code"`
	StatusMsg  string   `json:"status_msg"`
	Data       UserInfo `json:"user"`
}

// SendQueryUserResponse send response of get_user_info
func SendQueryUserResponse(c *gin.Context, err error, userInfo *UserInfo) {
	Err := errno.ConvertErr(err)
	c.JSON(http.StatusOK, QueryUserResponse{
		StatusCode: Err.ErrCode,
		StatusMsg:  Err.ErrMsg,
		Data: UserInfo{
			ID:   userInfo.ID,
			Name: userInfo.Name,
		},
	})
}

// UserParam req format for register/login
type UserParam struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

// UserResponse resp format of register/login
type UserResponse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	UserID     int64  `json:"user_id"`
	Token      string `json:"token"`
}

// SendUserResponse send response of register/login
func SendUserResponse(c *gin.Context, err error, userID int64, token string) {
	Err := errno.ConvertErr(err)
	c.JSON(http.StatusOK, UserResponse{
		StatusCode: Err.ErrCode,
		StatusMsg:  Err.ErrMsg,
		UserID:     userID,
		Token:      token,
	})
}

type PublishActionParam struct {
	Token string `json:"token" form:"token"`
	Data  []byte `json:"data" form:"data"`
	Title string `json:"title" form:"title"`
}

type PublishActionResponse struct {
	Code    int32       `json:"status_code"`
	Message string      `json:"status_msg"`
	Data    interface{} `json:"video_list"`
}

type PublishListParam struct {
	Token  string `json:"token" form:"token"`
	UserId int64  `json:"user_id" form:"user_id"`
}

// SendPublishActionResponse pack video response
func SendVideoResponse(c *gin.Context, err error, videolist interface{}) {
	Err := errno.ConvertErr(err)
	c.JSON(http.StatusOK, PublishActionResponse{
		Code:    Err.ErrCode,
		Message: Err.ErrMsg,
		Data:    videolist,
	})
}

type FeedParam struct {
	Token      string `json:"token" form:"token"`
	LatestTime int64  `json:"latest_time,omitempty" form:"latest_time"`
}
type FeedResponse struct {
	Code     int32       `json:"status_code"`
	Message  string      `json:"status_msg"`
	NextTime int64       `json:"next_time"`
	Data     interface{} `json:"video_list"`
}

// SendResponse pack response
func SendFeedResponse(c *gin.Context, err error, videolist interface{}, nexttime int64) {
	Err := errno.ConvertErr(err)
	c.JSON(http.StatusOK, FeedResponse{
		Code:     Err.ErrCode,
		Message:  Err.ErrMsg,
		NextTime: nexttime,
		Data:     videolist,
	})
}

type FavoriteActionParam struct {
	VideoId    int64  `json:"video_id" form:"video_id"`
	Token      string `json:"token" form:"token"`
	ActionType int32  `json:"action_type" form:"action_type"`
}
type FavoriteActionResponse struct {
	Code    int32       `json:"status_code"`
	Message string      `json:"status_msg"`
	Data    interface{} `json:"video_list"`
}

type FavoriteListParam struct {
	UserId int64  `json:"user_id" form:"user_id"`
	Token  string `json:"token" form:"token"`
}

func SendFavoriteResponse(c *gin.Context, err error, videolist interface{}) {
	Err := errno.ConvertErr(err)
	c.JSON(http.StatusOK, PublishActionResponse{
		Code:    Err.ErrCode,
		Message: Err.ErrMsg,
		Data:    videolist,
	})
}
