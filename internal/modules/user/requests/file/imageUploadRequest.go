package file

import "mime/multipart"

type ImageUploadRequest struct {
	Image *multipart.FileHeader `form:"image" binding:"required"`
}
