package handlers

import (
	"fmt"
	"gin-blog/internal/modules/user/logic/file"
	"gin-blog/pkg/response"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func UploadFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		filePath := fmt.Sprintf("uploads/%d/%s/%d", time.Now().Year(), time.Now().Month(), time.Now().Day())
		url, err := file.NewUpload(c, filePath).Upload()
		if err != nil {
			log.Println("user.handlers.UploadFile().Upload err=", err)
			response.HTTPFail(err, c)
			return
		}

		response.HTTPSuccess("上传成功", response.Data{
			Content: url,
		}, c)
	}
}
