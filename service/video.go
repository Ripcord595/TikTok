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
	ffmpeg "github.com/u2takey/ffmpeg-go"
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


// CosUpload 上传至云端，返回URL
func CosUpload(fileName string, reader io.Reader) (string, error) {
    // 解析 COS 的基本 URL
    u, _ := url.Parse(fmt.Sprintf(data.COS_URL_FORMAT, data.COS_BUCKET_NAME, data.COS_APP_ID, data.COS_REGION))
    b := &cos.BaseURL{BucketURL: u}
    
    // 创建 COS 客户端
    client := cos.NewClient(b, &http.Client{
        Transport: &cos.AuthorizationTransport{
            SecretID:  data.COS_SECRET_ID,
            SecretKey: data.COS_SECRET_KEY,
        },
    })
    
    // 使用 COS 客户端将文件上传到云存储，path 为文件在云存储中的路径
    _, err := client.Object.Put(context.Background(), fileName, reader, nil)
    if err != nil {
        panic(err) // 如果上传过程中发生错误，抛出一个 panic
    }
    
    // 返回上传后的文件的完整 URL
    return "https://dong-136240066.cos.ap-nanjing.myqcloud.com/" + fileName, nil
}

// ExampleReadFrameAsJpeg 获取封面
func ExampleReadFrameAsJpeg(inFileName string, frameNum int) io.Reader {
    buf := bytes.NewBuffer(nil) // 创建一个新的字节缓冲区
    
    // 使用 ffmpeg-go 库处理视频文件
    err := ffmpeg.Input(inFileName).
        Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}). // 过滤器，选择指定帧数之后的帧
        Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}). // 输出配置，提取一帧并以JPEG格式输出
        WithOutput(buf, os.Stdout). // 将输出指定到字节缓冲区 buf 中
        Run() // 执行命令
    
    if err != nil {
        panic(err) // 如果执行过程中出现错误，抛出一个 panic
    }
    
    return buf // 返回字节缓冲区，其中包含提取的封面图像数据
}
