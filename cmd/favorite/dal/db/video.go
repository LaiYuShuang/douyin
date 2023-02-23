package db

import (
	"context"
	"douyin/pkg/constants"
	"gorm.io/gorm"
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

func (v Video) TableName() string {
	return constants.VideoTableName
}

func SetVideo(ctx context.Context, vid int64, actionType int32) error {
	var expr string
	if actionType == 1 {
		expr = "favorite_count+1"
	} else {
		expr = "favorite_count-1"
	}
	if err := DB.WithContext(ctx).Model(&Video{}).Where("id = ? ", vid).Update("favorite_count", gorm.Expr(expr)).Error; err != nil {
		return err
	}
	return nil
}

func GetVideoById(ctx context.Context, videoId int64) (Video, error) {
	var vid Video
	err := DB.WithContext(ctx).Where("id=?", videoId).Find(&vid).Error
	if err != nil {
		return vid, err
	}
	return vid, nil
}
