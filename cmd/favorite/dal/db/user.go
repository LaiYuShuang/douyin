package db

import (
	"douyin/pkg/constants"

	"context"
)

type User struct {
	Id   int64  `gorm:"column:id"`
	Name string `gorm:"column:u_name"`
}

func (v User) TableName() string {
	return constants.UserTableName
}

func QueryUser(ctx context.Context, userId int64) (User, error) {
	var usr User
	if err := DB.WithContext(ctx).Where("id=?", userId).Take(&usr).Error; err != nil {
		return usr, err
	}
	return usr, nil
}
