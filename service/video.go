package service

import (
	"bytes"
	"context"
	"TikTok/data"
	"TikTok/model"
	"fmt"
	"TikTok/handler"
	_ "config/github.com/go-sql-driver/mysql"
	"github.com/tencentyun/cos-go-sdk-v5"
        "io"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/jinzhu/gorm"
)

const videoNum = 2 // 每次返回的视频数量

// FeedGet 获取视频列表
func FeedGet(lastTime int64) ([]model.Video, error) {
    if lastTime == 0 { // 如果没有传入参数或者视频已经刷完
        lastTime = time.Now().Unix() // 设置 lastTime 为当前时间的 Unix 时间戳
    }
    strTime := fmt.Sprint(time.Unix(lastTime, 0).Format("2006-01-02 15:04:05")) // 将 Unix 时间戳格式化为字符串
    fmt.Println("查询的时间", strTime) // 打印查询的时间
    var VideoList []model.Video
    VideoList = make([]model.Video, 0) // 创建一个空的 VideoList 切片
    // 使用数据库操作对象查询数据库中创建时间小于 strTime 的 videoNum 个视频，
    // 按创建时间降序排序，并限制结果数量为 videoNum，将结果存储到 VideoList 中
    err := data.SqlSession.Table("videos").Where("created_at < ?", strTime).Order("created_at desc").Limit(videoNum).Find(&VideoList).Error
    return VideoList, err // 返回 VideoList 切片和可能发生的错误
}

// AddCommentCount 增加评论计数
func AddCommentCount(videoId uint) error {
    // 使用数据库操作对象 data.SqlSession，更新 "videos" 表中 id 为 videoId 的记录的 comment_count 字段，
    // 使其自增 1，如果更新过程中发生错误，则将错误返回
    if err := data.SqlSession.Table("videos").Where("id = ?", videoId).Update("comment_count", gorm.Expr("comment_count + 1")).Error; err != nil {
        return err
    }
    return nil // 返回 nil 表示操作成功，无错误
}

// ReduceCommentCount 减少评论计数
func ReduceCommentCount(videoId uint) error {
    // 使用数据库操作对象 data.SqlSession，更新 "videos" 表中 id 为 videoId 的记录的 comment_count 字段，
    // 使其自减 1，如果更新过程中发生错误，则将错误返回
    if err := data.SqlSession.Table("videos").Where("id = ?", videoId).Update("comment_count", gorm.Expr("comment_count - 1")).Error; err != nil {
        return err
    }
    return nil // 返回 nil 表示操作成功，无错误
}


// GetVideoAuthor 获取视频的作者
func GetVideoAuthor(videoId uint) (uint, error) {
    var video model.Video
    // 使用数据库操作对象 data.SqlSession，查询 "videos" 表中 id 为 videoId 的记录，
    // 将查询结果存储到 video 变量中。如果查询过程中发生错误，则将错误返回
    if err := data.SqlSession.Table("videos").Where("id = ?", videoId).Find(&video).Error; err != nil {
        return video.ID, err // 返回 video 的 ID 和可能发生的错误
    }
    return video.AuthorId, nil // 返回 video 的 AuthorId 以及 nil 表示操作成功，无错误
}


// CreateVideo 添加一条视频信息
func CreateVideo(video *model.Video) {
    // 使用数据库操作对象 data.SqlSession，向 "videos" 表中插入一条 video 记录
    data.SqlSession.Table("videos").Create(&video)
}

// GetVideoList 根据用户ID查找所有与该用户相关视频信息
func GetVideoList(userId uint) []model.Video {
    var videoList []model.Video
    // 使用数据库操作对象 data.SqlSession，查询 "videos" 表中 author_id 为 userId 的所有记录，
    // 将查询结果存储到 videoList 切片中
    data.SqlSession.Table("videos").Where("author_id=?", userId).Find(&videoList)
    return videoList // 返回 videoList 切片，其中包含与该用户相关的所有视频信息
}


// CosUpload 上传至云端，返回url
func CosUpload(fileName string, reader io.Reader) (string, error) {
	u, _ := url.Parse(fmt.Sprintf(dao.COS_URL_FORMAT, dao.COS_BUCKET_NAME, dao.COS_APP_ID, dao.COS_REGION))
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  dao.COS_SECRET_ID,
			SecretKey: dao.COS_SECRET_KEY,
		},
	})
	//path为本地的保存路径
	_, err := client.Object.Put(context.Background(), fileName, reader, nil)
	if err != nil {
		panic(err)
	}
	return "https://dong-1305843950.cos.ap-nanjing.myqcloud.com/" + fileName, nil
}

// ExampleReadFrameAsJpeg 获取封面
func ExampleReadFrameAsJpeg(inFileName string, frameNum int) io.Reader {
	buf := bytes.NewBuffer(nil)
	err := ffmpeg.Input(inFileName).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		panic(err)
	}
	return buf
}
