package file

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"os"
	"path/filepath"
)

type Upload struct {
	err  error
	c    *gin.Context
	file *multipart.FileHeader
	path string
}

func NewUpload(c *gin.Context, path string) *Upload {
	return &Upload{c: c, path: path}
}

func (u *Upload) Upload() (string, error) {
	u.GetFile().CreateDir().SaveFile()
	if u.err != nil {
		return "", u.err
	}

	return fmt.Sprintf("%s%s", "http://localhost:8899/", u.path), u.err
}

// GetFile 获取文件
func (u *Upload) GetFile() *Upload {
	if u.err != nil {
		return u
	}

	u.file, u.err = u.c.FormFile("file")
	return u
}

// CreateDir 创建不存在的目录
func (u *Upload) CreateDir() *Upload {
	if u.err != nil {
		return u
	}
	u.path = fmt.Sprintf("%s/%s", u.path, u.file.Filename)

	u.err = os.MkdirAll(filepath.Dir(u.path), os.ModePerm)
	return u
}

// SaveFile 保存文件
func (u *Upload) SaveFile() *Upload {
	if u.err != nil {
		return u
	}

	u.err = u.c.SaveUploadedFile(u.file, u.path)
	return u
}
