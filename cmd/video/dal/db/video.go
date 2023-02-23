package db

import (
	"context"
	"douyin/pkg/constants"
)

type Video struct {
	Id            int64  `gorm:"column:id;primaryKey;not null"`
	UserId        int64  `gorm:"column:u_id;not null"`
	PlayUrl       string `gorm:"column:play_url;not null"`
	CoverUrl      string `gorm:"column:cover_url;not null"`
	Title         string `gorm:"column:title;not null"`
	FavoriteCount int64  `gorm:"column:favorite_count"`
	CommentCount  int64  `gorm:"column:comment_count"`
	CreateTime    int64  `gorm:"column:create_time"`
}
type VID struct {
	Id int64 `gorm:"column:id"`
}

func (v VID) TableName() string {
	return constants.VideoTableName
}

func (v Video) TableName() string {
	return constants.VideoTableName
}

func CreateVideo(ctx context.Context, vio Video) error {
	if err := DB.WithContext(ctx).Create(&vio).Error; err != nil {
		return err
	}
	return nil
}

func GetVideoIdList(ctx context.Context, uid int64) ([]int64, error) {
	vid := make([]VID, 0)
	res := make([]int64, 0)
	if err := DB.WithContext(ctx).Select("id").Where("u_id=?", uid).Find(&vid).Error; err != nil {
		return nil, err
	}
	for _, i := range vid {
		res = append(res, i.Id)
	}
	return res, nil
}

func GetFeedIdList(ctx context.Context, time int64, limit int) ([]int64, error) {
	res := make([]int64, 0)
	vid := make([]VID, 0)
	if time == 0 {
		time = constants.MaxTime
	}
	err := DB.WithContext(ctx).Select("id").Where("create_time <= ?", time).Order("create_time desc").Limit(limit).Find(&vid).Error
	if err != nil {
		return nil, err
	}
	for _, i := range vid {
		res = append(res, i.Id)
	}
	return res, nil
}

func GetVideoById(ctx context.Context, videoId int64) (Video, error) {
	var vid Video
	err := DB.WithContext(ctx).Where("id=?", videoId).Find(&vid).Error
	if err != nil {
		return vid, err
	}
	return vid, nil
}
