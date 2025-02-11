package user

type RegisterRequest struct {
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required,min=3,max=11"`
	Name     string `form:"name" binding:"required,min=1,max=10"`
}
