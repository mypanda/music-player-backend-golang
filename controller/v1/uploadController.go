package v1

import (
	"music-player-backend/response"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type UploadController struct{}

func NewUploadController() UploadController {
	return UploadController{}
}

func (g *UploadController) Create(DB *mongo.Client) func(c *gin.Context) {
	return func(c *gin.Context) {

		// 获取上传文件
		file, err := c.FormFile("file")
		if err != nil {
			response.Error(100001, "文件不存在")
			return
		}

		// 判断后缀
		fileExt := path.Ext(file.Filename)
		extMap := map[string]bool{
			".mp3": true,
			".mp4": true,
			".m4a": true,
		}
		if _, ok := extMap[fileExt]; !ok {
			response.Error(100002, "文件类型不支持")
			return
		}

		// 创建文件夹
		day := time.Now().Format("20060102")
		dir := "./static/" + day
		if err := os.MkdirAll(dir, 0666); err != nil {
			response.Error(100003, "创建文件夹失败")
			return
		}

		// 创建文件名
		fileUnixName := strconv.FormatInt(time.Now().UnixNano(), 10)
		dst := path.Join(dir, fileUnixName+fileExt)

		// 保存
		err = c.SaveUploadedFile(file, dst)
		if err != nil {
			response.Error(100001, "文件保存错误")
			return
		}

		response.Success("/" + dst)
	}
}
