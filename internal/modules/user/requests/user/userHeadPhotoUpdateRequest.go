package user

type UserHeadPhotoUpdateRequest struct {
	UserID uint
	Url    string `form:"url" json:"url" binding:"required,url"`
}
