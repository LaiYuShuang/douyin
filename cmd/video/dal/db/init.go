package db

import (
	"douyin/pkg/constants"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormopentracing "gorm.io/plugin/opentracing"
)

var DB *gorm.DB

func Init() {
	var err error
	DB, err = gorm.Open(mysql.Open(constants.MySQLDefaultDSN))
	if err != nil {
		panic(err)
	}

	if err = DB.Use(gormopentracing.New()); err != nil {
		panic(err)
	}

	//create db table
	//m := DB.Migrator()
	//if m.HasTable(constants.VideoTableName) {
	//	return
	//}
	//if err = m.CreateTable(&Video{}); err != nil {
	//	panic(err)
	//}
}
