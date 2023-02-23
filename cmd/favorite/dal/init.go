package dal

import (
	"douyin/cmd/favorite/dal/db"
)

// Init init dal
func Init() {
	db.Init() // mysql init
}
