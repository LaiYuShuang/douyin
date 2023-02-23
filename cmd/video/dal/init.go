package dal

import (
	"douyin/cmd/video/dal/db"
	"douyin/cmd/video/dal/minicache"
)

// Init init dal
func Init() {
	db.Init()        // mysql init
	minicache.Init() //cache
}
