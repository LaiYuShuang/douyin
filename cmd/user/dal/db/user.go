package db

import (
	"context"

	"douyin/pkg/constants"
)

type User struct {
	ID       int64  `gorm:"column:id;primaryKey;not null"`
	Username string `gorm:"column:u_name;unique;type:varchar(30);not null"`
	Password string `gorm:"column:passwd;type:varchar(60);not null"`
}

func (u *User) TableName() string {
	return constants.UserTableName
}

// QueryUserByID get user information by id
func QueryUserByID(ctx context.Context, targetID int64) ([]*User, error) {
	res := make([]*User, 0)
	if err := DB.WithContext(ctx).Where("id = ?", targetID).Limit(1).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// CreateUser create user info
func CreateUser(ctx context.Context, uname, password string) ([]*User, error) {
	users := []*User{{
		Username: uname,
		Password: password,
	}}
	res := make([]*User, 0)

	err := DB.WithContext(ctx).Create(users).Error
	if err != nil {
		return nil, err
	}

	err = DB.WithContext(ctx).Where("u_name = ?", uname).Limit(1).Find(&res).Error
	if err != nil {
		return nil, err
	}

	return res, nil
}

// QueryUser query list of user info
func QueryUser(ctx context.Context, userName string) ([]*User, error) {
	res := make([]*User, 0)
	if err := DB.WithContext(ctx).Where("u_name = ?", userName).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
