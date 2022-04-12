package api

import (
	"mall/global"
	"mall/response"

	"github.com/gin-gonic/gin"
)

func WebFileUpload(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		response.Failed("图片上传出错", c)
	}
	image := global.Config.Upload
	err = c.SaveUploadedFile(file, image.SavePath+file.Filename)
	if err != nil {
		return
	}
	imageURL := image.AccessUrl + file.Filename
	response.Success("图片上传成功", imageURL, c)
}
