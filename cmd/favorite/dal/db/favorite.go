package db

import (
	"context"
	"douyin/pkg/constants"
)

type Favorite struct {
	UserId  int64 `gorm:"column:u_id;not null"`
	VideoId int64 `gorm:"column:v_id;not null"`
}

func (f Favorite) TableName() string {
	return constants.FavoriteTableName
}

func UpdateFavorite(ctx context.Context, uid int64, vid int64, actionType int32) error {
	var favorite Favorite
	favorite.UserId = uid
	favorite.VideoId = vid
	if actionType == 1 {
		if err := DB.WithContext(ctx).Create(&favorite).Error; err != nil {
			return err
		}
	} else {
		if err := DB.WithContext(ctx).Where("u_id = ? and v_id = ?", uid, vid).Delete(&favorite).Error; err != nil {
			return err
		}
	}
	return nil
}

func GetVideoIdList(ctx context.Context, uid int64) ([]int64, error) {
	res := make([]int64, 0)
	favorites := make([]*Favorite, 0)
	err := DB.WithContext(ctx).Select("v_id").Where("u_id = ?", uid).Find(&favorites).Error
	if err != nil {
		return nil, err
	}
	for _, ff := range favorites {
		res = append(res, ff.VideoId)
	}
	return res, nil
}
