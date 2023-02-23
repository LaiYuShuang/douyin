package minicache

import (
	"context"
	"douyin/cmd/video/dal/db"
	"douyin/cmd/video/dal/minicache/cache"
	"errors"
	"math/rand"
	"time"

	"gorm.io/gorm"
)

func GetVideo(ctx context.Context, videoId int64) (db.Video, error) {
	key := cache.StringInt64(videoId)
	// 随机选择1-60的过期时间
	rand.Seed(time.Now().UnixNano())
	sum32 := rand.Intn(60)
	sum64 := int64(sum32)
	sum64 += 1
	onMissed := func(ctx context.Context, key interface{}) (data interface{}, err error) {
		keys := int64(key.(cache.StringInt64))
		return db.GetVideoById(ctx, keys)
	}
	tmpStr := "video"
	//fmt.Println("-----------------------------------------------------------------------")
	v, err, _ := Caches.Get(key, tmpStr, cache.WithOpOnMissed(onMissed), cache.WithOpTTL(time.Duration(sum64)*time.Second), cache.WithOpContext(ctx))

	video := (db.Video)(v.(db.Video))

	return video, err
}

func SaveVideo(ctx context.Context, vid db.Video) error {
	//create table and delete cache becase change the videolist
	if err := db.CreateVideo(ctx, vid); err != nil {
		return err
	} else {
		vid := cache.StringInt64(vid.UserId)
		Caches.DeleteVideoCache(ctx, vid, "user")
	}

	return nil
}

func GetVideoList(ctx context.Context, userId int64) ([]*db.Video, error) {
	//
	vid := make([]*db.Video, 0)
	key := cache.StringInt64(userId)
	// 随机选择1-60的过期时间
	rand.Seed(time.Now().UnixNano())
	sum32 := rand.Intn(60)
	sum64 := int64(sum32)
	sum64 += 1
	onMissed := func(ctx context.Context, key interface{}) (data interface{}, err error) {
		keys := int64(key.(cache.StringInt64))
		return db.GetVideoIdList(ctx, keys)
	}
	tmpStr := "user"
	v, err, _ := Caches.Get(key, tmpStr, cache.WithOpOnMissed(onMissed), cache.WithOpTTL(time.Duration(sum64)*time.Second), cache.WithOpContext(ctx))

	vidList := ([]int64)(v.([]int64))

	for _, id := range vidList {
		video, err := GetVideo(ctx, id)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		vid = append(vid, &video)
	}

	return vid, err
}

func GetFeedList(ctx context.Context, lastTime int64, limit int) ([]*db.Video, error) {
	var vid []*db.Video
	vidList, err := db.GetFeedIdList(ctx, lastTime, limit)
	if err != nil {
		return nil, err
	}
	for _, id := range vidList {
		video, err := GetVideo(ctx, id)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		vid = append(vid, &video)
	}
	return vid, nil

}
