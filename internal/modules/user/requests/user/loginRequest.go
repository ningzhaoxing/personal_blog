package user

type LoginRequest struct {
	UID      uint
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required,min=3,max=11"`
}
