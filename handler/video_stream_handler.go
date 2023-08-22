package handler

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

func StreamVideo(c *gin.Context) {
	// 获取视频ID或文件名，你可以从路由参数中获取
	videoID := c.Param("video_id")

	// 从数据库或文件系统中获取视频文件路径
	videoFilePath := GetVideoFilePath(videoID)

	// 打开视频文件
	videoFile, err := os.Open(videoFilePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open video file"})
		return
	}
	defer videoFile.Close()

	// 设置响应头以指示返回的内容是视频流
	c.Header("Content-Type", "video/mp4")

	// 从视频文件中将数据复制到响应主体，以进行视频流
	_, err = io.Copy(c.Writer, videoFile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to stream video"})
		return
	}
}
