package db

import (
	"context"
	"douyin/pkg/constants"
	"errors"
	"gorm.io/gorm"
)

type Favorite struct {
	UserId  int64 `gorm:"column:u_id;not null"`
	VideoId int64 `gorm:"column:v_id;not null"`
}

func (v Favorite) TableName() string {
	return constants.FavoriteTableName
}

func IsFavorite(ctx context.Context, uid int64, vid int64) (bool, error) {
	res := make([]*Favorite, 0)
	err := DB.WithContext(ctx).Where("u_id = ? and v_id = ?", uid, vid).Find(&res).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	} else {
		return false, err
	}
	return true, nil
}
