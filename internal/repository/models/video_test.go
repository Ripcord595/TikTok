package models

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var (
	videoListByUserIdCache    = make(map[int][]*Video)
	videoListByUserIdCacheMux sync.Mutex

	videoListByLimitCache    = make(map[cacheKey][]*Video)
	videoListByLimitCacheMux sync.Mutex
)

type cacheKey struct {
	Limit     int
	Timestamp int64
}

func setupTest() {
	InitDB()
}

func getVideoListByUserIdFromCache(userID int) ([]*Video, bool) {
	videoListByUserIdCacheMux.Lock()
	defer videoListByUserIdCacheMux.Unlock()

	videos, exists := videoListByUserIdCache[userID]
	return videos, exists
}

func addVideoListByUserIdToCache(userID int, videos []*Video) {
	videoListByUserIdCacheMux.Lock()
	defer videoListByUserIdCacheMux.Unlock()

	videoListByUserIdCache[userID] = videos
}

func getVideoListByLimitFromCache(limit int, timestamp time.Time) ([]*Video, bool) {
	key := cacheKey{
		Limit:     limit,
		Timestamp: timestamp.Unix(),
	}

	videoListByLimitCacheMux.Lock()
	defer videoListByLimitCacheMux.Unlock()

	videos, exists := videoListByLimitCache[key]
	return videos, exists
}

func addVideoListByLimitToCache(limit int, timestamp time.Time, videos []*Video) {
	key := cacheKey{
		Limit:     limit,
		Timestamp: timestamp.Unix(),
	}

	videoListByLimitCacheMux.Lock()
	defer videoListByLimitCacheMux.Unlock()

	videoListByLimitCache[key] = videos
}

func printVideos(videos []*Video) {
	for _, v := range videos {
		fmt.Printf("%#v\n", *v)
	}
}

func Test_QueryVideoListByUserId(t *testing.T) {
	setupTest()

	userID := 1
	videos, exists := getVideoListByUserIdFromCache(userID)

	if !exists {
		videos = []*Video{}
		err := NewVideoDAO().QueryVideoListByUserId(int64(userID), &videos)
		if err != nil {
			panic(err)
		}

		addVideoListByUserIdToCache(userID, videos)
	}

	printVideos(videos)
}

func Test_QueryVideoListByLimit(t *testing.T) {
	setupTest()

	limit := 2
	timestamp := time.Unix(1630794872, 0)
	videos, exists := getVideoListByLimitFromCache(limit, timestamp)

	if !exists {
		videos = []*Video{}
		err := NewVideoDAO().QueryVideoListByLimitAndTime(limit, timestamp, &videos)
		if err != nil {
			panic(err)
		}

		addVideoListByLimitToCache(limit, timestamp, videos)
	}

	printVideos(videos)
}
