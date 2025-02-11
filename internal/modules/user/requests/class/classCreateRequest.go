package class

type ClassCreateRequest struct {
	Name string `form:"name" binding:"required,min=1,max=10"`
}
