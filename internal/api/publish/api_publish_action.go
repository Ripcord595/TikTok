package publish

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"tiktok/conf"
	"tiktok/internal/repository/models"
	"tiktok/internal/services"
	"tiktok/pkg/util"
)

var (
	videoIndexMap = map[string]struct{}{
		".mp4":  {},
		".avi":  {},
		".wmv":  {},
		".flv":  {},
		".mpeg": {},
		".mov":  {},
	}
	pictureIndexMap = map[string]struct{}{
		".jpg": {},
		".bmp": {},
		".png": {},
		".svg": {},
	}
)

func PublishVideoHandler(c *gin.Context) {
	rawId, _ := c.Get("user_id")

	userId, ok := rawId.(int64)
	if !ok {
		PublishVideoError(c, "解析UserId出错")
		return
	}

	title := c.PostForm("title")

	form, err := c.MultipartForm()
	if err != nil {
		PublishVideoError(c, err.Error())
		return
	}

	files := form.File["data"]
	for _, file := range files {
		suffix := filepath.Ext(file.Filename)
		if _, ok := videoIndexMap[suffix]; !ok {
			PublishVideoError(c, "不支持的视频格式")
			continue
		}
		name := util.NewFileName(userId)
		filename := name + suffix
		config := conf.NewConfig()
		savePath := filepath.Join(config.Path.StaticSourcePath, filename)
		err = c.SaveUploadedFile(file, savePath)
		if err != nil {
			PublishVideoError(c, err.Error())
			continue
		}
		err = util.SaveImageFromVideo(name, false)
		if err != nil {
			PublishVideoError(c, err.Error())
			continue
		}
		err := services.PostVideo(userId, filename, name+util.GetDefaultImageSuffix(), title)
		if err != nil {
			PublishVideoError(c, err.Error())
			continue
		}
		PublishVideoOk(c, file.Filename+"上传成功")
	}
}

func PublishVideoError(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, models.CommonResponse{StatusCode: 1,
		StatusMsg: msg})
}

func PublishVideoOk(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, models.CommonResponse{StatusCode: 0, StatusMsg: msg})
}
